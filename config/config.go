package config

type AppConfig struct {
	MYSQL_USER     string
	MYSQL_HOST     string
	MYSQL_PORT     int
	MYSQL_DBNAME   string
	MYSQL_PASSWORD string
}

func GetConfig() *AppConfig {
	return &AppConfig{
		MYSQL_DBNAME:   "db_bankku",
		MYSQL_HOST:     "database-3.cj67dqd97qqz.us-east-1.rds.amazonaws.com",
		MYSQL_PORT:     3306,
		MYSQL_USER:     "admin",
		MYSQL_PASSWORD: "vizucode",
	}
}
