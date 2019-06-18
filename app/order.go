package app

import "github.com/carmo-evan/holly/model"

// CreateOrder takes a userID and inserts a new order into the store
func (a *App) CreateOrder(userID string) (*model.Order, error) {
	o := &model.Order{UserID: userID}
	return a.Store.Order().Insert(o)
}

// AddOrderItem creates an orderItem, assigns it to an order, and returns it
func (a *App) AddOrderItem(orderID, skuID string, quantity int64) (*model.OrderItem, error) {
	oi := &model.OrderItem{SKUID: skuID, Price: 1500, OrderID: orderID, Quantity: quantity}
	return a.Store.OrderItem().Insert(oi)
}
