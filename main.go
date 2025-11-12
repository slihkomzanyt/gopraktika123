package main

import (
	"fmt"
)

type Product struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
}

type Inventory struct {
	products map[int]Product
}

func (inv *Inventory) SellProduct(id, qty int) error {
	product, exists := inv.products[id]
	if !exists {
		return fmt.Errorf("товар не найден")
	}
	if product.Quantity < qty {
		return fmt.Errorf("недостаточно товара")
	}
	product.Quantity -= qty
	inv.products[id] = product
	return nil
}

func (inv *Inventory) AddProduct(p Product) {
	inv.products[p.ID] = p
}

func main() {

	inv := &Inventory{
		products: make(map[int]Product),
	}

	inv.AddProduct(Product{
		ID:       1,
		Name:     "Laptop",
		Price:    1000.0,
		Quantity: 5,
	})

	inv.AddProduct(Product{
		ID:       2,
		Name:     "Phone",
		Price:    500.0,
		Quantity: 10,
	})

	err := inv.SellProduct(1, 3)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Продажа успешна!")
	}

	err = inv.SellProduct(1, 5)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
}
