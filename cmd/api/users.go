package main

import (
	"context"
	"errors"
	"net/http"
	"social/internal/store"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type userKey string

const userCtx userKey = "user"

type FollowUserPayload struct {
	UserID int64 `json:"user_id"`
}

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	user, err := getUserFromContext(r)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) followUserHandler(w http.ResponseWriter, r *http.Request) {
	followerUser, err := getUserFromContext(r)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	// TODO: revert back to auth userID from ctx
	var payload FollowUserPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestError(w, r, err)
		return
	}

	if err := app.store.Followers.Follow(r.Context(), followerUser.ID, payload.UserID); err != nil {
		switch err {
		case store.ErrConflict:
			app.conflictResponse(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (app *application) unfollowUserHandler(w http.ResponseWriter, r *http.Request) {
	followerUser, err := getUserFromContext(r)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	// TODO: revert back to auth userID from ctx
	var payload FollowUserPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestError(w, r, err)
		return
	}

	if err := app.store.Followers.Unfollow(r.Context(), followerUser.ID, payload.UserID); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (app *application) userContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userID, err := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 64)
		if err != nil {
			app.badRequestError(w, r, err)
			return
		}

		user, err := app.store.Users.GetByID(ctx, userID)
		if err != nil {
			switch {
			case errors.Is(err, store.ErrNotFound):
				app.notFoundError(w, r, err)
			default:
				app.internalServerError(w, r, err)
			}
			return
		}

		ctx = context.WithValue(ctx, userCtx, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserFromContext(r *http.Request) (*store.User, error) {
	user, ok := r.Context().Value(userCtx).(*store.User)
	if !ok {
		return nil, errors.New("not possible to cast user from context")
	}
	return user, nil
}
