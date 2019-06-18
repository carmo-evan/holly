package main

import (
	"sync"

	app "github.com/carmo-evan/holly/app"
	"github.com/carmo-evan/holly/model"
	_ "github.com/lib/pq"
)

func main() {
	app, err := app.NewApp()
	if err != nil {
		panic(err)
	}
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

	schan := make(chan *model.SKU, 1)

	go func() {
		_, s, err := app.CreateProductWithDefaultSku(1999, istr, istr, istr)
		if err != nil {
			panic(err)
		}
		schan <- s
	}()

	ochan := make(chan *model.Order, 1)

	go func() {
		u, err := app.CreateUser(istr + "@mailinator.com")
		if err != nil {
			panic(err)
		}
		o, err := app.CreateOrder(u.UserID)
		if err != nil {
			panic(err)
		}
		ochan <- o
	}()

	o := <-ochan
	s := <-schan
	app.AddOrderItem(o.OrderID, s.SKUID, 1)
}
