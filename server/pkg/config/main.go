package config

import (
	"log"
	"os"
	"path"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/tahsinature/future-proof-gin/pkg/utilities"
)

type EntryArgsType struct {
	Run       bool `arg:"-r,help:Run the server"`
	Seed      bool `arg:"--seed,help:Seed the database"`
	SyncForce bool `arg:"--syncf,help:Sync Force the database"`
}

var (
	DB        dbConfig
	App       appConfig
	JWT       jwtConfig
	Redis     redisConfig
	EntryArgs EntryArgsType
	Telegram  telegramConfig
	Other     otherConfig
)

func Validate() {
	_, file, _, _ := runtime.Caller(0)
	serverRootPath := path.Join(file, "..", "..", "..")

	if envPath := path.Join(serverRootPath, ".env"); godotenv.Load(envPath) != nil {
		log.Fatal("error: failed to load the env file")
	}

	DB = dbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBname:   os.Getenv("DB_NAME"),
	}

	App = appConfig{
		Port:         os.Getenv("PORT"),
		Environment:  os.Getenv("ENV"),
		SSL:          os.Getenv("SSL") == "TRUE",
		IsProduction: os.Getenv("ENV") == "PRODUCTION",
		APIVersion:   os.Getenv("API_VERSION"),
	}

	JWT = jwtConfig{
		AccessSecret:  os.Getenv("ACCESS_SECRET"),
		RefreshSecret: os.Getenv("REFRESH_SECRET"),
	}

	Redis = redisConfig{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       os.Getenv("REDIS_DB"),
	}

	Telegram = telegramConfig{
		BOT_TOKEN:         os.Getenv("TELEGRAM_BOT_TOKEN"),
		CHAT_ID_ME_SERVER: os.Getenv("TELEGRAM_CHAT_ID_ME_SERVER"),
	}

	Other = otherConfig{
		RootPath: serverRootPath,
	}

	utilities.ValidateMultipleStruct(DB, App, JWT, Redis, Other, Telegram)
}
