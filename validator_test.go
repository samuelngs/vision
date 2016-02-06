package vision

import "testing"

type LoginParameters struct {
	username string `param:"string,min:3;max:12,required"`
	password string `param:"string,strict-required"`
	gender   string `param:"integer,enum:1|2|8,required"`
}

func TestValidate(t *testing.T) {

	_, err := Validate(&LoginParameters{
		username: "samuel",
		password: "password",
		gender:   "1",
	})

	if err != nil {
		t.Fatalf("Unexpected error result. Error: %v", err)
	}
}

func TestFailValidate(t *testing.T) {

	_, err := Validate(&LoginParameters{
		username: "samuel",
		password: " ",
		gender:   "3",
	})

	if err == nil {
		t.Fatalf("Expected to have error result. However no error has been returned.")
	}
}
