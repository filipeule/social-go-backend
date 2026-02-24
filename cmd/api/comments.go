package main

import (
	"net/http"
	"social/internal/store"
)

type CreateCommentPayload struct {
	Content string `json:"content" validate:"required,max=1000"`
}

func (app *application) createComment(w http.ResponseWriter, r *http.Request) {
	post, err := getPostFromCtx(r)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	var commentPayload CreateCommentPayload
	if err := readJSON(w, r, &commentPayload); err != nil {
		app.badRequestError(w, r, err)
		return
	}

	if err := Validate.Struct(commentPayload); err != nil {
		app.badRequestError(w, r, err)
		return
	}

	comment := &store.Comment{
		PostID: post.ID,
		UserID: post.UserID,
		Content: commentPayload.Content,
	}

	if err := app.store.Comments.Create(r.Context(), comment); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, comment); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
