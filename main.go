package main

import (
	"github.com/carmo-evan/holly/model"
	store "github.com/carmo-evan/holly/store/sqlstore"
	_ "github.com/lib/pq"
)

func main() {
	p := &model.Product{}
	s := store.NewSQLStore()
	s.Product().Save(p)
	s.DB.Close()
}
