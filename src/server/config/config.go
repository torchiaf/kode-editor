package config

import (
	"os"

	models "server/models"
	"server/utils"
)

type Config struct {
	IsDev     bool
	App       string
	Namespace string
	Routes    []models.Route
}

func isDevEnv() bool {
	env := os.Getenv("DEV_ENV")
	if len(env) > 0 {
		return true
	}
	return false
}

func GetConfig() Config {

	routes := utils.ParseFile[models.Routes]("assets/routes")

	c := Config{
		IsDev:     isDevEnv(),
		App:       "code-editor",
		Namespace: os.Getenv("POD_NAMESPACE"),
		Routes:    routes.Routes,
	}

	return c
}
