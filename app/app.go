package app

import (
	"errors"

	"github.com/carmo-evan/holly/store"
	"github.com/carmo-evan/holly/store/sqlstore"
)

type App struct {
	Store store.Store
}

type Options struct {
	DbFlavor store.DbFlavor
}

func NewApp(options Options) (*App, error) {

	if options.DbFlavor == store.Postgres {
		s, err := sqlstore.NewSQLStore()
		app := &App{Store: s}
		return app, err
	}
	return nil, errors.New("Undetermined db flavor")
}
