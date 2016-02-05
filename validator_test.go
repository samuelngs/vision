package vision

import "testing"

type LoginParameters struct {
	username string `valid:"string;min:3;max:12"`
	password string `valid:"string"`
}

func TestValidate(t *testing.T) {
	Validate(&LoginParameters{
		username: "samuel",
		password: "123456",
	})
}
