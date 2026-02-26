package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type Follower struct {
	UserID     int64  `json:"user_id"`
	FollowerID int64  `json:"follower_id"`
	CreatedAt  string `json:"created_at"`
}

type FollowStore struct {
	db *sql.DB
}

func (s *FollowStore) Follow(ctx context.Context, followerID, userID int64) error {
	query := `
		insert into followers (user_id, follower_id) values ($1, $2)
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, userID, followerID)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return ErrConflict
		}

		return err
	}

	return nil
}

func (s *FollowStore) Unfollow(ctx context.Context, followerID, userID int64) error {
	query := `
		delete from followers where user_id = $1 and follower_id = $2
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, userID, followerID)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return ErrConflict
		}

		return err
	}

	return nil
}