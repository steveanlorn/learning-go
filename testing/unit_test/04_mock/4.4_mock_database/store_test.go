package store

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestStore_CreateUser(t *testing.T) {
	username := "admin"
	password := "123"
	expectedUserID := 2
	expectedErr := false

	t.Log("Given the need to create user in database successfully")
	db, mock, err := sqlmock.New(
		sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual),
		sqlmock.MonitorPingsOption(true),
	)

	if err != nil {
		t.Fatalf("\t%s\tShould able to open a stub database connection: %v", failed, err)
	}
	defer db.Close()

	{
		mock.ExpectPing()
		mock.ExpectBegin()
		mock.ExpectPrepare(queryInsertUser).
			ExpectQuery().
			WithArgs("admin", "123").
			WillReturnRows(sqlmock.NewRows([]string{"user_id"}).AddRow(2))
		mock.ExpectCommit()
	}

	store := NewStore(db)
	userID, err := store.CreateUser(context.Background(), username, password)
	if (err != nil) != expectedErr {
		t.Errorf("\t%s\tShould get an error: %v: got %v", failed, expectedErr, err)
	} else {
		t.Logf("\t%s\tShould get an error: %v: got %v", succeed, expectedErr, err != nil)
	}

	if userID != expectedUserID {
		t.Errorf("\t%s\tShould get an user ID: %d: got %d", failed, expectedUserID, userID)
	} else {
		t.Logf("\t%s\tShould get an user ID: %d: got %d", succeed, expectedUserID, userID)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("\t%s\tShould met with db expectation: %s", failed, err)
	} else {
		t.Logf("\t%s\tShould met with db expectation", succeed)
	}
}
