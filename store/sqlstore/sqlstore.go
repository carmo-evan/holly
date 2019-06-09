package sqlstore

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/carmo-evan/holly/store"
	"gopkg.in/gorp.v1"
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
	ProductStore store.ProductStore
	SKUStore     store.SKUStore
	Tx           *gorp.Transaction
	dbMap        *gorp.DbMap
}

func NewSQLStore() (store.Store, error) {
	dbMap := &gorp.DbMap{Db: getDb(), Dialect: gorp.PostgresDialect{}}
	dbMap.TraceOn("[gorp]", log.New(os.Stdout, "Holly:", log.Lmicroseconds))
	tx, err := dbMap.Begin()
	store := &SQLStore{dbMap: dbMap, Tx: tx}
	return store, err
}

func (s *SQLStore) Product() store.ProductStore {
	if s.ProductStore == nil {
		s.ProductStore = &SQLProductStore{SQLStore: s}
		s.ProductStore.CreateTable()
	}
	return s.ProductStore
}

func (s *SQLStore) SKU() store.SKUStore {
	if s.SKUStore == nil {
		s.SKUStore = &SQLSKUStore{SQLStore: s}
		s.SKUStore.CreateTable()
	}
	return s.SKUStore
}

func (s *SQLStore) Close() error {
	return s.dbMap.Db.Close()
}

func (s *SQLStore) Commit() error {
	return s.Tx.Commit()
}

func getDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("new DB")
	return db
}
