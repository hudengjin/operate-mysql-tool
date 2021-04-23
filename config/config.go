package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct {
	AppName string
	Debug bool
	LogPath string
	LogMaxSize int
	LogMaxAge int
	LogMaxBackups int
	LogIsCompress bool
	MySqlHost string
	MySqlPort string
	MySqlUsername string
	MySqlPassword string
	MySqlDataBase string
}

func GetEnv() *Env {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Cannot load .env file, please check the file is exist!")
	}
	return &Env{
		AppName: os.Getenv("APP_NAME"),
		Debug: true,
		LogPath: os.Getenv("LOG_PATH"),
		LogMaxSize: parseInt(os.Getenv("LOG_MAX_SIZE"), 128),
		LogMaxAge: parseInt(os.Getenv("LOG_MAX_AGE"), 7),
		LogMaxBackups: parseInt(os.Getenv("LOG_MAX_BACKUPS"), 2),
		LogIsCompress: false,
		MySqlHost: os.Getenv("MYSQL_HOST"),
		MySqlPort: os.Getenv("MYSQL_PORT"),
		MySqlUsername: os.Getenv("MYSQL_USERNAME"),
		MySqlPassword: os.Getenv("MYSQL_PASSWORD"),
		MySqlDataBase: os.Getenv("MYSQL_DB_NAME"),
	}
}

func parseInt(intString string, defaultInt int) int {
	intValue, err := strconv.Atoi(intString)
	if err != nil {
		return defaultInt
	}
	return intValue
}