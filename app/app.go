package app

import (
	"github.com/carmo-evan/holly/config"
	"github.com/carmo-evan/holly/store"
	"github.com/carmo-evan/holly/store/sqlstore"
)

type App struct {
	Store store.Store
}

func NewApp() (*App, error) {
	var err error

	app := &App{}
	envConfig := config.GetEnvConfig()
	switch envConfig.DbCredentials.Flavor {
	case "postgres":
		s, err := sqlstore.NewStore(envConfig)
		app.Store = s
		return app, err
	}
	return app, err
}
