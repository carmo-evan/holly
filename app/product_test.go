package app

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	err := os.Setenv("HOLLY_ENV", "DEV")
	if err != nil {
		panic(err)
	}
	code := m.Run()
	os.Exit(code)
}

func TestCreateProduct(t *testing.T) {
	app, err := NewApp()
	if err != nil {
		t.Error(err)
	}
	p, err := app.CreateProduct("testProduct", "Test Product", "Product for testing")

	if p.ProductID == "" {
		t.Fatalf("Did not create ID for product")
	}

	if err != nil {
		t.Error(err)
	}
}
