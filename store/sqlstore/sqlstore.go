package store

import (
	"fmt"

	"github.com/carmo-evan/holly/model"
	"github.com/jinzhu/gorm"
)

const (
	host   = "localhost"
	port   = 3000
	user   = "postgres"
	dbname = "holly"
)

/*SQLStore interface implements the store.Store interface, with a method to retrieve
the specific store for each one of the app's entities */
type SQLStore struct {
	productStore *SQLProductStore
	DB           *gorm.DB
}

func (s *SQLStore) Product() *SQLProductStore {
	if s.productStore == nil {
		s.DB.AutoMigrate(&model.Product{})
		s.productStore = &SQLProductStore{SQLStore: *s}
	}

	return s.productStore
}

func NewSQLStore() *SQLStore {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic("failed to connect database")
	}
	return &SQLStore{DB: db}
}
