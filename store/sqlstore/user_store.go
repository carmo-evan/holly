package sqlstore

import (
	"fmt"
	"time"

	"github.com/carmo-evan/holly/model"
	uuid "github.com/satori/go.uuid"
)

// SQLUserStore is a wrapper with User methods for the SqlStore
type SQLUserStore struct {
	SQLStore *SQLStore
}

// Insert commits new p to the database
func (su *SQLUserStore) Insert(u *model.User) (*model.User, error) {
	u.UserID = uuid.NewV4().String()
	u.CreatedAt = time.Now().UnixNano()
	u.UpdatedAt = u.CreatedAt
	err := su.SQLStore.Tx.Insert(u)
	if err != nil {
		err = fmt.Errorf("[SQLUserStore] error in calling Insert: %v", err)
	}

	return u, err
}

// Update saves changes made to p to the database
func (su *SQLUserStore) Update(u *model.User) (*model.User, error) {
	u.UpdatedAt = time.Now().UnixNano()
	_, err := su.SQLStore.Tx.Update(u)
	return u, err
}

// Get returns the User with the provided id
func (su *SQLUserStore) Get(id string) (*model.User, error) {
	u, err := su.SQLStore.Tx.Get(model.User{}, id)
	return u.(*model.User), err
}

// CreateTable executes a gorp mapping for model.User
func (su *SQLUserStore) CreateTable() error {
	// add a table, setting the table name to 'Users' and
	// suecifying that the Id property is an auto incrementing PK
	su.SQLStore.dbMap.AddTableWithName(model.User{}, "users").SetKeys(false, "user_id")

	// create the table. in a Userion system you'd generally
	// use a migration tool, or create the tables via scripts
	err := su.SQLStore.dbMap.CreateTablesIfNotExists()
	return err
}
