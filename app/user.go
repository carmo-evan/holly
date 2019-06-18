package app

import "github.com/carmo-evan/holly/model"

// CreateUser takes an email and inserts a new user into the store
func (a *App) CreateUser(email string) (*model.User, error) {
	u := &model.User{Email: email}
	return a.Store.User().Insert(u)
}
