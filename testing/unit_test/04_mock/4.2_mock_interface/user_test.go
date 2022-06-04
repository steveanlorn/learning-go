package user

import (
	"context"
	"testing"
	"time"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestService_CreateUser(t *testing.T) {
	type argument struct {
		username string
		password string
	}

	testsCase := []struct {
		label       string
		getContext  func() (context.Context, func())
		mockStore   MockStore
		argument    argument
		expectedErr bool
	}{
		{
			label: "userCreated",
			getContext: func() (context.Context, func()) {
				return context.Background(), func() {}
			},
			mockStore: MockStore{
				FakeGetUserByUsername: func(ctx context.Context, username string) (User, error) {
					return User{}, ErrUserNotFound
				},
				FakeCreateUser: func(ctx context.Context, username string, password string) error {
					return nil
				},
			},
			argument: argument{
				username: "admin",
				password: "123",
			},
			expectedErr: false,
		},
		{
			label: "userExist",
			getContext: func() (context.Context, func()) {
				return context.Background(), func() {}
			},
			mockStore: MockStore{
				FakeGetUserByUsername: func(ctx context.Context, username string) (User, error) {
					return User{Username: username}, nil
				},
				FakeCreateUser: func(ctx context.Context, username string, password string) error {
					return nil
				},
			},
			argument: argument{
				username: "admin",
				password: "123",
			},
			expectedErr: true,
		},
		{
			label: "timeout",
			getContext: func() (context.Context, func()) {
				ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
				return ctx, cancel
			},
			mockStore: MockStore{
				FakeGetUserByUsername: func(ctx context.Context, username string) (User, error) {
					time.Sleep(200 * time.Millisecond)
					return User{}, ErrUserNotFound
				},
				FakeCreateUser: func(ctx context.Context, username string, password string) error {
					return nil
				},
			},
			argument: argument{
				username: "admin",
				password: "123",
			},
			expectedErr: true,
		},
	}

	t.Logf("Given the need to create a new user")
	for _, tc := range testsCase {
		tf := func(t *testing.T) {
			ctx, cancel := tc.getContext()
			defer cancel()

			service := NewService(&tc.mockStore)
			err := service.CreateUser(ctx, tc.argument.username, tc.argument.password)

			if (err != nil) != tc.expectedErr {
				t.Fatalf("\t%s\tShould get an error is %v: %v", failed, tc.expectedErr, err)
			}
			t.Logf("\t%s\tShould get an error is %v", succeed, tc.expectedErr)
		}
		t.Run(tc.label, tf)
	}

}
