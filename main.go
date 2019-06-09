package main

import (
	"sync"

	"github.com/carmo-evan/holly/model"
	store "github.com/carmo-evan/holly/store"
	"github.com/carmo-evan/holly/store/sqlstore"
	_ "github.com/lib/pq"
)

func main() {
	str, err := sqlstore.NewSQLStore()
	if err != nil {
		panic(err)
	}
	defer str.Commit()
	defer str.Close()
	str.Product()
	str.SKU()
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go doStuff(&wg, i, str)
	}
	wg.Wait()
}

func doStuff(wg *sync.WaitGroup, i int, str store.Store) {
	defer wg.Done()

	istr := string(i%84 + 1)
	p := &model.Product{Name: istr, DisplayName: istr, Description: istr}
	p, err := str.Product().Insert(p)

	if err != nil {
		panic(err)
	}

	s := &model.SKU{Name: p.Name, Description: p.Description, ProductID: p.ProductID}
	s, err = str.SKU().Insert(s)

	if err != nil {
		panic(err)
	}

	p.DefaultSKUID = s.SKUID
	p.DisplayName = p.DisplayName + "Updated"
	str.Product().Update(p)
	if err != nil {
		panic(err)
	}
}
