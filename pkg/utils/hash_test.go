package utils_test

import (
	"testing"

	"github.com/v1bh475u/LibMan_MVC/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type passwordTest struct {
	password string
	err      error
}

var passwordTests = []passwordTest{
	{
		password: "test",
		err:      nil,
	},
	{
		password: "well this is totally and i mean totally and for certainty randomly generated password and if you say otherwise then you're just a hater...",
		err:      bcrypt.ErrPasswordTooLong,
	},
}

func TestHashPassword(t *testing.T) {
	for _, test := range passwordTests {
		_, err := utils.HashPassword(test.password)
		if err != test.err {
			t.Fatalf("Error hashing password: %v", err)
		}
	}
}
