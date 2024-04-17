package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

/*Определить структуру для хранения данных о товарах магазина, включая название товара, 
категорию товара, количество товара, цену товара, место хранения. Реализовать методы 
добавления/удаления информации о товаре, методы поиска по названию и категории, поиск 
товара по диапазону цен в данной категории товаров, загрузка/выгрузка данных в файл*/

// Product представляет информацию о товаре.
type Product struct {
	Name       string  `json:"name"`
	Category   string  `json:"category"`
	Quantity   int     `json:"quantity"`
	Price      float64 `json:"price"`
	Storage    string  `json:"storage"`
}

// Inventory представляет список товаров.
type Inventory struct {
	Products []Product `json:"products"`
}

// AddProduct добавляет товар в инвентарь.
func (inv *Inventory) AddProduct(p Product) {
	inv.Products = append(inv.Products, p)
}

// RemoveProduct удаляет товар из инвентаря по названию.
func (inv *Inventory) RemoveProduct(name string) {
	for i, p := range inv.Products {
		if p.Name == name {
			inv.Products = append(inv.Products[:i], inv.Products[i+1:]...)
			break
		}
	}
}

// FindProductByName ищет товар по названию.
func (inv *Inventory) FindProductByName(name string) *Product {
	for _, p := range inv.Products {
		if p.Name == name {
			return &p
		}
	}
	return nil
}

// FindProductsByCategory ищет товары по категории.
func (inv *Inventory) FindProductsByCategory(category string) []Product {
	var result []Product
	for _, p := range inv.Products {
		if p.Category == category {
			result = append(result, p)
		}
	}
	return result
}

// LoadInventory загружает данные из файла.
func LoadInventory(filename string) (*Inventory, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var inv Inventory
	err = json.Unmarshal(data, &inv)
	if err != nil {
		return nil, err
	}
	return &inv, nil
}

// SaveInventory сохраняет данные в файл.
func (inv *Inventory) SaveInventory(filename string) error {
	data, err := json.MarshalIndent(inv, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

func main() {
	// Пример использования:
	inv := Inventory{}
	inv.AddProduct(Product{Name: "Apple", Category: "Fruits", Quantity: 10, Price: 1.5, Storage: "A1"})
	inv.AddProduct(Product{Name: "Banana", Category: "Fruits", Quantity: 20, Price: 0.8, Storage: "B2"})

	// Сохраняем данные в файл.
	err := inv.SaveInventory("inventory.json")
	if err != nil {
		fmt.Println("Ошибка при сохранении данных:", err)
		return
	}

	// Загружаем данные из файла.
	loadedInv, err := LoadInventory("inventory.json")
	if err != nil {
		fmt.Println("Ошибка при загрузке данных:", err)
		return
	}

	// Поиск товара по названию.
	apple := loadedInv.FindProductByName("Apple")
	if apple != nil {
		fmt.Printf("Найден товар: %s, Цена: %.2f\n", apple.Name, apple.Price)
	} else {
		fmt.Println("Товар не найден.")
	}

	// Поиск товаров по категории.
	fruits := loadedInv.FindProductsByCategory("Fruits")
	fmt.Println("Товары в категории 'Fruits':")
	for _, p := range fruits {
		fmt.Printf("%s, Количество: %d\n", p.Name, p.Quantity)
	}
}
