// package app

// import (
// 	"flag"
// 	"log"
// 	"os"

// 	"github.com/Dimasaldian/letsAdopt/app/controllers"
// 	"github.com/joho/godotenv"
// )



// func getEnv(key, fallback string) string {
// 	if value, ok := os.LookupEnv(key); ok {
// 		return value
// 	}

// 	return fallback
// }

// func Run() {
// 	var server = controllers.Server{}
// 	var appConfig = controllers.AppConfig{}
// 	var dbConfig = controllers.DBConfig{}

// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatalf("Error on loading .env file")
// 	}



// 	appConfig.AppName = getEnv("APP_NAME", "LetsAdopt")
// 	appConfig.AppEnv = getEnv("APP_ENV", "development")
// 	appConfig.AppPort = getEnv("APP_PORT", "9000")

// 	dbConfig.DBHost = getEnv("DB_HOST", "localhost")
// 	dbConfig.DBUser = getEnv("DB_USER", "user")
// 	dbConfig.DBPassword= getEnv("DB_PASSWORD", "password")
// 	dbConfig.DBName = getEnv("DB_NAME", "dbname")
// 	dbConfig.DBPort = getEnv("DB_PORT", "5432")

// 	flag.Parse()
// 	arg := flag.Arg(0)
// 	if arg != "" {
// 		server.InitCommands(appConfig, dbConfig)
// 	} else {
// 		server.Initialize(appConfig, dbConfig)
// 		server.Run(":" + appConfig.AppPort)
// 	}
// }
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
    dbConfig.DBPassword= getEnv("DB_PASSWORD")
    dbConfig.DBName = getEnv("DB_NAME")
    dbConfig.DBPort = getEnv("DB_PORT")

    // Hapus atau sesuaikan bagian ini jika tidak diperlukan di Vercel
    // flag.Parse()
    // arg := flag.Arg(0)
    // if arg != "" {
    //     server.InitCommands(appConfig, dbConfig)
    // } else {
        server.Initialize(appConfig, dbConfig)
        server.Run(":" + appConfig.AppPort)
    // }
}