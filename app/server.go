package app

import (
	"log"
	"os"

	"github.com/Dimasaldian/letsAdopt/app/controllers"
)

func getEnv(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	// Tidak perlu fallback value, biarkan error muncul jika variabel tidak ada
	log.Fatalf("Environment variable %s not found", key)
	return ""
}

func Run() {
	var server = controllers.Server{}
	var appConfig = controllers.AppConfig{}
	var dbConfig = controllers.DBConfig{}

	appConfig.AppName = getEnv("APP_NAME")
	appConfig.AppEnv = getEnv("APP_ENV")
	appConfig.AppPort = getEnv("APP_PORT")

	dbConfig.DBHost = getEnv("DB_HOST")
	dbConfig.DBUser = getEnv("DB_USER")
	dbConfig.DBPassword = getEnv("DB_PASSWORD")
	dbConfig.DBName = getEnv("DB_NAME")
	dbConfig.DBPort = getEnv("DB_PORT")

	server.Initialize(appConfig, dbConfig)
	server.Run(":" + appConfig.AppPort)
	// }
}
