package db

func InitDB() {
	ConnectDB()
	InitQueryBuilder()
	InitRegisterRepo()
	InitProductRepo()
	InitOrderRepo()
}
