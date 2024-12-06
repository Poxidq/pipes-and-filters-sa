package config

import (
	"fmt"
	"log"

	env "github.com/caarlos0/env/v11"
)

type config struct {
	EmailAddress    string `env:"EMAIL_ADDRESS"`
	EmailPassword   string `env:"EMAIL_PASSWORD"`
	EmailRecipients string `env:"EMAIL_RECIPIENTS"`
	ApiPort         string `env:"API_PORT"`
}

var CFG config

func LoadConfig() {
	err := env.Parse(&CFG)

	// res, err := env.ParseAs[config]()
	if err != nil {
		panic(fmt.Sprintf("Error when loading environment %v", err))
	}
	log.Printf("%+v\n", CFG)
	// CFG = res

	if CFG.ApiPort == "" {
		panic("Cannot load .env")
		// CFG.ApiPort = "8080"
	}
}
