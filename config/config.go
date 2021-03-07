package config

import (
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	TESLA_REST_API_URL string
	DB_USERNAME        string
	DB_PASSWORD        string
	DB_PORT            int
	DB_HOST            string
	DB_NAME            string
	DB_MAXIDLECONNS    int
	DB_MAXOPENCONNS    int
}

func GetConfig() Configuration {
	configuration := Configuration{}
	gonfig.GetConf("config/config.json", &configuration)
	return configuration
}
