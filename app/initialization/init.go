package initialization

func Init() {
	initConfig()
	initLog()
	initDb()
	initRedis()
}
