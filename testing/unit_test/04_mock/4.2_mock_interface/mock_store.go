package user

import (
	"context"
)

type MockStore struct {
	FakeGetUserByUsername func(ctx context.Context, username string) (User, error)
	FakeCreateUser        func(ctx context.Context, username string, password string) error
}

func (m *MockStore) GetUserByUsername(ctx context.Context, username string) (User, error) {
	user, err := m.FakeGetUserByUsername(ctx, username)
	return user, err
}

func (m *MockStore) CreateUser(ctx context.Context, username string, password string) error {
	err := m.FakeCreateUser(ctx, username, password)
	return err
}
