package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

var users []User

func main() {
	http.HandleFunc("/create", create)
	http.HandleFunc("/list", list)
	http.HandleFunc("/update", update)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func update(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPut {
		var update_user User
		err := json.NewDecoder(r.Body).Decode(&update_user)

		if err != nil {

			http.Error(w, "Erroe user to update", http.StatusBadRequest)
			return
		}

		var detect_user *User

		for i, user := range users {

			if user.ID == update_user.ID {
				detect_user = &users[i]
				break
			}
		}

		if detect_user == nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		detect_user.Name = update_user.Name
		detect_user.Password = update_user.Password

		// Return a success response
		fmt.Fprintf(w, "User updated successfully: %s", detect_user.Name)
		return

	}
}

func create(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err == nil {
			users = append(users, user)
			fmt.Fprintf(w, "Register Successfuly: %s", user.Name)
			return
		}
		http.Error(w, "Error users to json", http.StatusInternalServerError)
		return
	}
	http.Error(w, "method not allowed", http.StatusInternalServerError)

}
func list(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		userJson, err := json.Marshal(users)
		if err == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(userJson)
			return
		}
		http.Error(w, "Error users to json", http.StatusInternalServerError)
		return
	}
	http.Error(w, "Method not allowed ", http.StatusMethodNotAllowed)
}
