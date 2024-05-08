package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"gobasic/web/message"
	"net/http"

	_ "github.com/lib/pq"
)

type User struct {
	ID       int    `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	Password string `db:"password" json:"password"`
}

const (
	host     = "localhost"
	port     = 5432
	userx    = "root"
	password = "admin"
	dbname   = "root"
)

var users []User

func Create(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		var new_user User
		err := json.NewDecoder(r.Body).Decode(&new_user)

		if err == nil {
			users = append(users, new_user)
			id := new_user.ID
			name := new_user.Name
			pass := new_user.Password

			psql := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, userx, password, dbname)

			DB, err1 := sql.Open("postgres", psql)

			CheckError(err1)

			defer DB.Close()

			insertStmt := `INSERT INTO employee(employeeid, name, password) VALUES($1, $2, $3)`

			_, err2 := DB.Exec(insertStmt, id, name, pass)

			CheckError(err2)

			fmt.Fprintf(w, "Register Successfuly: %s", new_user.Name)
			return
		}
		message.Send_Error(w, http.StatusInternalServerError, "Error in json message", "")

		return
	}
	message.Send_Error(w, http.StatusMethodNotAllowed, "Method not allowed", "")
}

func CheckError(err error) {

	if err != nil {
		panic(err)
	}

}
