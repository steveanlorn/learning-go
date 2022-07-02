package user

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var (
	// ErrUserNotFound ...
	ErrUserNotFound = errors.New("user is nor found")
)

//go:generate mockgen -package user -source ./user.go -destination ./storemock_test.go

// Store denotes store layer to get User data.
type Store interface {
	// GetUserByUsername denotes a getter for user by username.
	// Concrete type of Store must return ErrUserNotFound to denote a not found user.
	GetUserByUsername(ctx context.Context, username string) (User, error)
	CreateUser(ctx context.Context, username string, password string) error
}

// Service ...
type Service struct {
	store Store
}

func NewService(store Store) *Service {
	service := Service{
		store: store,
	}
	return &service
}

// User ...
type User struct {
	Username string
}

// CreateUser ...
func (s *Service) CreateUser(ctx context.Context, username string, password string) error {

	errChan := make(chan error, 1)

	go func() {
		sanitizedUsername := sanitizeUsername(username)

		_, err := s.store.GetUserByUsername(ctx, sanitizedUsername)
		switch err {
		case ErrUserNotFound:
			break
		case nil:
			errChan <- fmt.Errorf("user with username %s is already exist in store", username)
			return
		default:
			errChan <- err
			return
		}

		hashedPassword, err := hashPassword([]byte(password))
		if err != nil {
			errChan <- err
			return
		}

		err = s.store.CreateUser(ctx, sanitizedUsername, hashedPassword)
		if err != nil {
			errChan <- err
			return
		}

		errChan <- nil
	}()

	select {
	case err := <-errChan:
		return err
	case <-ctx.Done():
		return fmt.Errorf("timeout while creating user: %v", ctx.Err())
	}
}

func sanitizeUsername(username string) string {
	sanitizedUsername := strings.ToLower(username)
	return sanitizedUsername
}

func hashPassword(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
