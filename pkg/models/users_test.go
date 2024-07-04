package models_test

import (
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/v1bh475u/LibMan_MVC/pkg/models"
	"github.com/v1bh475u/LibMan_MVC/pkg/types"
)

type InsertUser struct {
	User          types.User
	expectedError *mysql.MySQLError
}

var InsertUserTests = []InsertUser{
	{
		User: types.User{
			Username: "test",
			Password: "test",
			Role:     "user",
		},
		expectedError: nil,
	},
	{
		User: types.User{
			Username: "test1",
			Password: "test1",
			Role:     "test",
		},
		expectedError: &mysql.MySQLError{Number: 0x4f1, SQLState: [5]uint8{0x30, 0x31, 0x30, 0x30, 0x30}, Message: "Data truncated for column 'Role' at row 1"},
	},
}

func TestInsertUser(t *testing.T) {
	err := godotenv.Load("../../config/.env")
	if err != nil {
		t.Fatalf("Error loading .env file")
	}
	for _, test := range InsertUserTests {
		err := models.InsertUser(test.User)
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number != test.expectedError.Number {
				t.Fatalf("Error inserting user: %v", err)
			}
		} else if err != nil {
			t.Fatalf("Expected MySQL error, got %v", err)
		}
	}
}
