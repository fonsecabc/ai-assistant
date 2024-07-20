package configs

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Variables struct {
	OpenAIAPIKey       string `envconfig:"OPEN_AI_API_KEY" required:"true"`
	OpenAIModel        string `envconfig:"OPEN_AI_MODEL" required:"true"`
	OpenAIMaxTokens    int    `envconfig:"OPEN_AI_MAX_TOKENS" required:"true"`
	HubspotAccessToken string `envconfig:"HUBSPOT_ACCESS_TOKEN" required:"true"`
	HubspotAPIBaseURL  string `envconfig:"HUBSPOT_API_BASE_URL" required:"true"`
	ServerPort         string `envconfig:"SERVER_PORT" required:"true"`
}

var Env Variables

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	err = envconfig.Process("", &Env)
	if err != nil {
		panic(err)
	}
}
