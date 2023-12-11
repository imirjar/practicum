package main

import "testing"

func TestFullName(t *testing.T) {
	u := User{
		FirstName: "Misha",
		LastName:  "Popov",
	}

	if err := u.FullName(); err != "Misha Popov" {
		t.Errorf("Какая-то шляпа!")
	}
}
