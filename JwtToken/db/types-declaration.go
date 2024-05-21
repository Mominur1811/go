package db

type User struct {
	UserName  string `db:"username"         json:"username"`
	Password  string `db:"password"         json:"password"`
	ContactNo string `db:"contactno"        json:"contactno"`
}

type LoginData struct {
	UserName string `db:"username"         json:"username"`
	Password string `db:"password"         json:"password"`
}

const (
	host     = "localhost"
	port     = "5432"
	dbName   = "Product_Info"
	sslMode  = "disable"
	user     = "root"
	password = "admin"
)
