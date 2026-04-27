package cache

import (
	"context"
	"social/internal/store"

	"github.com/stretchr/testify/mock"
)

func NewMockCacheStorage() Storage {
	return Storage{
		Users: &MockCacheStore{},
	}
}

type MockCacheStore struct{
	mock.Mock
}

func (m *MockCacheStore) Get(ctx context.Context, userID int64) (*store.User, error) {
	args := m.Called(userID)
	return nil, args.Error(1)
}

func (m *MockCacheStore) Set(ctx context.Context, user *store.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockCacheStore) Delete(ctx context.Context, userID int64) {
	m.Called(userID)
}