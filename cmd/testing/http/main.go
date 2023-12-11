package main

import (
	"log"
	"net/http"
)

// User — основной объект для теста.
type User struct {
	ID        string
	FirstName string
	LastName  string
}

func main() {
	users := make(map[string]User)
	u1 := User{
		ID:        "u1",
		FirstName: "Misha",
		LastName:  "Popov",
	}
	u2 := User{
		ID:        "u2",
		FirstName: "Sasha",
		LastName:  "Popov",
	}
	users["u1"] = u1
	users["u2"] = u2

	http.HandleFunc("/users", UserViewHandler(users))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
