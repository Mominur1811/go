package db

func InitDB() {
	ConnectDB()
	InitQueryBuilder()
	InitRedis()
	InitRegisterRepo()
	InitProductRepo()
	InitOrderRepo()
	//ClearRedisClient()
}
