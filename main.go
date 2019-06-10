package main

import (
	"sync"

	app "github.com/carmo-evan/holly/app"
	"github.com/carmo-evan/holly/model"
	store "github.com/carmo-evan/holly/store"
	_ "github.com/lib/pq"
)

func main() {
	ops := app.Options{DbFlavor: store.Postgres}
	app, err := app.NewApp(ops)
	if err != nil {
		panic(err)
	}
	defer app.Store.Commit()
	defer app.Store.Close()

	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go doStuff(&wg, i, app)
	}
	wg.Wait()
}

func doStuff(wg *sync.WaitGroup, i int, app *app.App) {
	defer wg.Done()

	istr := string(i%84 + 1)
	p := &model.Product{Name: istr, DisplayName: istr, Description: istr}
	p, err := app.Store.Product().Insert(p)

	if err != nil {
		panic(err)
	}

	s := &model.SKU{Name: p.Name, Description: p.Description, ProductID: p.ProductID, Price: 1999}
	s, err = app.Store.SKU().Insert(s)

	if err != nil {
		panic(err)
	}

	p.DefaultSKUID = s.SKUID
	p.DisplayName = p.DisplayName + "Updated"
	p, err = app.Store.Product().Update(p)
	if err != nil {
		panic(err)
	}

	oi := &model.OrderItem{SkuID: s.SKUID, Price: 1500}
	oi, err = app.Store.OrderItem().Insert(oi)
	if err != nil {
		panic(err)
	}
	o := &model.Order{OrderItemIDs: oi.OrderItemID, Total: oi.Price}
	o, err = app.Store.Order().Insert(o)
	oi.OrderID = o.OrderID
	app.Store.OrderItem().Update(oi)
}
