package main

import (
	"tavern/domain/product"
	"tavern/services/order"
	"tavern/services/tavern"
	"github.com/google/uuid"
)

func main(){
	products := productRepository()
	
	os,err := order.NewOrderService(
		order.WithMemoryCustomerRepository(),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		panic(err)
	}

	tav, err := tavern.NewTavern(
		tavern.WithOrderService(os))
	if err != nil {
		panic(err)
	}

	uid, err := os.AddCustomer("Vibin")
	if err != nil {
		panic(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	err = tav.Order(uid, order)
	if err != nil {
		panic(err)
	}

}


func productRepository() []product.Product {
	mem, err := product.NewProduct("Memory", "16GB DDR", 150.00)
	if err != nil {
		panic(err)
	}

	cpu, err := product.NewProduct("CPU", "AMD Ryzen 7", 450.00)
	if err != nil {
		panic(err)
	}

	gpu, err := product.NewProduct("GPU", "NVIDEA 4070", 550.00)
	if err != nil {
		panic(err)
	}

	return []product.Product{
		mem, cpu, gpu,
	}
}