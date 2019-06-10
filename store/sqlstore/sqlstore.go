package sqlstore

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/carmo-evan/holly/model"
	"github.com/carmo-evan/holly/store"
	"gopkg.in/gorp.v1"
)

const (
	host   = "localhost"
	port   = 3000
	user   = "postgres"
	dbname = "holly"
)

/*SQLStore interface implements the *SQLStore interface, with a method to retrieve
the specific store for each one of the app's entities */
type SQLStore struct {
	ProductStore   store.ProductStore
	SKUStore       store.SKUStore
	OrderStore     store.OrderStore
	OrderItemStore store.OrderItemStore
	Tx             *gorp.Transaction
	dbMap          *gorp.DbMap
}

// NewSQLStore creates the underlying DB connection, gorp.dbMap and gorp.Transaction
func NewSQLStore() (store.Store, error) {
	dbMap := &gorp.DbMap{Db: getDb(), Dialect: gorp.PostgresDialect{}}
	dbMap.TraceOn("[gorp]", log.New(os.Stdout, "Holly:", log.Lmicroseconds))
	tx, err := dbMap.Begin()
	store := &SQLStore{dbMap: dbMap, Tx: tx}
	store.createStores()
	store.createTables()
	return store, err
}

// Product initializes the underlying table and returns its store layer
func (s *SQLStore) Product() store.ProductStore {
	return s.ProductStore
}

// SKU initializes the underlying table and returns its store layer
func (s *SQLStore) SKU() store.SKUStore {
	return s.SKUStore
}

// Order initializes the underlying table and returns its store layer
func (s *SQLStore) Order() store.OrderStore {
	return s.OrderStore
}

// OrderItem initializes the underlying table and returns its store layer
func (s *SQLStore) OrderItem() store.OrderItemStore {
	return s.OrderItemStore
}

// Close is a wrapper around the underlying DB Close method
func (s *SQLStore) Close() error {
	return s.dbMap.Db.Close()
}

// Commit is a wrapper around the underlying transaction commit method
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

func (s *SQLStore) createTables() error {
	s.dbMap.AddTableWithName(model.Product{}, "products").SetKeys(false, "product_id")
	s.dbMap.AddTableWithName(model.SKU{}, "skus").SetKeys(false, "sku_id")
	s.dbMap.AddTableWithName(model.Order{}, "orders").SetKeys(false, "order_id")
	s.dbMap.AddTableWithName(model.OrderItem{}, "order_items").SetKeys(false, "order_item_id")
	return s.dbMap.CreateTablesIfNotExists()
}

func (s *SQLStore) createStores() {
	s.OrderItemStore = &SQLOrderItemStore{SQLStore: s}
	s.OrderStore = &SQLOrderStore{SQLStore: s}
	s.ProductStore = &SQLProductStore{SQLStore: s}
	s.SKUStore = &SQLSKUStore{SQLStore: s}
}
