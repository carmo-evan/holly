package sqlstore

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/carmo-evan/holly/config"
	"github.com/carmo-evan/holly/model"
	"github.com/carmo-evan/holly/store"
	_ "github.com/lib/pq"
	"gopkg.in/gorp.v1"
)

var gorpFlavorMap = map[string]gorp.Dialect{
	"postgres": gorp.PostgresDialect{},
}

/*SQLStore interface implements the *SQLStore interface, with a method to retrieve
the specific store for each one of the app's entities */
type SQLStore struct {
	ProductStore   store.ProductStore
	SKUStore       store.SKUStore
	OrderStore     store.OrderStore
	UserStore      store.UserStore
	OrderItemStore store.OrderItemStore
	Tx             *gorp.Transaction
	dbMap          *gorp.DbMap
}

// NewStore creates the underlying DB connection, gorp.dbMap and gorp.Transaction
func NewStore(env *config.EnvConfig) (store.Store, error) {
	dbMap := &gorp.DbMap{Db: getDb(env), Dialect: gorpFlavorMap[env.DbCredentials.Flavor]}
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

// User initializes the underlying table and returns its store layer
func (s *SQLStore) User() store.UserStore {
	return s.UserStore
}

// Close is a wrapper around the underlying DB Close method
func (s *SQLStore) Close() error {
	s.Tx.Commit()
	return s.dbMap.Db.Close()
}

func getDb(env *config.EnvConfig) *sql.DB {

	db, err := sql.Open(env.DbCredentials.Flavor, fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		env.DbCredentials.Host, env.DbCredentials.Port, env.DbCredentials.User, env.DbCredentials.Dbname))
	if err != nil {
		panic(err)
	}
	fmt.Println("new DB")
	return db
}

func (s *SQLStore) createTables() error {
	s.dbMap.AddTableWithName(model.Product{}, "products").SetKeys(false, "product_id")
	s.dbMap.AddTableWithName(model.SKU{}, "skus").SetKeys(false, "sku_id")
	s.dbMap.AddTableWithName(model.Order{}, "orders").SetKeys(false, "order_id")
	s.dbMap.AddTableWithName(model.User{}, "users").SetKeys(false, "user_id")
	s.dbMap.AddTableWithName(model.OrderItem{}, "order_items").SetKeys(false, "order_item_id")
	return s.dbMap.CreateTablesIfNotExists()
}

func (s *SQLStore) createStores() {
	s.OrderItemStore = &SQLOrderItemStore{SQLStore: s}
	s.OrderStore = &SQLOrderStore{SQLStore: s}
	s.ProductStore = &SQLProductStore{SQLStore: s}
	s.SKUStore = &SQLSKUStore{SQLStore: s}
	s.UserStore = &SQLUserStore{SQLStore: s}
}
