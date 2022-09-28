package config

import (
	"log"
	"os"
	"strconv"
	"sync"
	"wallet-api/internal/model"

	"github.com/lib/pq"
)

var config *model.Config
var doOnce sync.Once

type GetConfigFn func() *model.Config

func GetConfig() *model.Config {

	doOnce.Do(func() {
		config = &model.Config{
			SentryDNS:          os.Getenv("SENTRY_DNS"),
			Environment:        os.Getenv("ENV"),
			IsHeroku:           os.Getenv("IS_HEROKU"),
			DbConnString:       os.Getenv("DB_CONNSTRING"),
			AwsAccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
			AwsSecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
			Region:             os.Getenv("AWS_REGION"),
			Bucket:             os.Getenv("AWS_BUCKET"),
		}

		if config.IsHeroku == "1" {
			var err error
			config.DbConnString, err = pq.ParseURL(os.Getenv("DATABASE_URL"))
			if err != nil {
				log.Fatal("Error parsing DB Connection")
			}
			config.DbConnString += " sslmode=require"
		}

		if config.Environment == "" {
			config.Environment = "PRODUCTION"
		}

		restPortString := os.Getenv("PORT")
		if restPortString == "" {
			restPortString = "8030"
		}

		restPort, err := strconv.Atoi(restPortString)
		if err != nil {
			panic(err.Error())
		}

		config.RestPort = restPort
	})

	return config
}
