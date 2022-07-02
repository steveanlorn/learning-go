package user

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestService_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockStore(ctrl)

	type argument struct {
		username string
		password string
	}

	testsCase := []struct {
		label       string
		argument    argument
		getContext  func() (context.Context, func())
		setMock     func()
		expectedErr bool
	}{
		{
			label: "userCreated",
			argument: argument{
				username: "Admin",
				password: "123",
			},
			getContext: func() (context.Context, func()) {
				return context.Background(), func() {}
			},
			setMock: func() {
				mockStore.EXPECT().
					GetUserByUsername(gomock.Any(), "admin").
					Return(User{}, ErrUserNotFound)
				mockStore.EXPECT().
					CreateUser(gomock.Any(), "admin", gomock.Any()).
					Return(nil)
			},
			expectedErr: false,
		},
		{
			label: "userExist",
			argument: argument{
				username: "Admin",
				password: "123",
			},
			getContext: func() (context.Context, func()) {
				return context.Background(), func() {}
			},
			setMock: func() {
				mockStore.EXPECT().
					GetUserByUsername(gomock.Any(), "admin").
					Return(User{"admin"}, nil)
			},
			expectedErr: true,
		},
		{
			label: "timeout",
			argument: argument{
				username: "Admin",
				password: "123",
			},
			getContext: func() (context.Context, func()) {
				ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
				return ctx, cancel
			},
			setMock: func() {
				mockStore.EXPECT().
					GetUserByUsername(gomock.Any(), "admin").
					DoAndReturn(func(ctx context.Context, username string) error {
						time.Sleep(200 * time.Millisecond)
						return nil
					})
			},
			expectedErr: true,
		},
	}

	t.Logf("Given the need to create a new user")
	for _, tc := range testsCase {
		tf := func(t *testing.T) {
			tc.setMock()

			ctx, cancel := tc.getContext()
			defer cancel()

			service := NewService(mockStore)
			err := service.CreateUser(ctx, tc.argument.username, tc.argument.password)

			if (err != nil) != tc.expectedErr {
				t.Fatalf("\t%s\tShould get an error is %v: %v", failed, tc.expectedErr, err)
			}
			t.Logf("\t%s\tShould get an error is %v", succeed, tc.expectedErr)
		}
		t.Run(tc.label, tf)
	}
}
