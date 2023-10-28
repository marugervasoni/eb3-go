package main

import "fmt"

//estructura product
//slice global de Product, Products instanciado con valores
//2 metodos asociados a Product Save() y GetAll()
//1 funcion getById()
//ejecutar al menos una vez cada metodo y funcion en main

type Product struct {
	ID			int
	Name		string
	Price		float64
	Description	string
	Category	string
}

var Products []Product


func main() {

	prod1:= Product{ID: 1, Name: "Puma Run", Price: 75000.00, Description: "Running Sneakers", Category: "Woman"}
	prod2:= Product{ID: 2, Name: "Nike 124", Price: 18900.00, Description: "Sport Cap, black color", Category: "No gender"}

	//guardar productos
	prod1.Save()
	prod2.Save()

	//listar productos
	prod1.GetAll()

	//buscar producto por id
	id := 3
	product :=getById(id)
	if product != nil {
		fmt.Printf("Product with ID %d founded: %s\n", id, product.Name)
	} else {
		fmt.Printf("Producto with ID %d not founded \n", id)
	}
}

//Metodo que debera tomar slice products y añadir el producto desde el cual se llama el metodo
func (p *Product) Save(){
	Products = append(Products, *p) 
}

//deberá imprimir todos los productos guardados en slice products
func (p* Product) GetAll(){
	fmt.Println("Products list")
	for _, product := range Products {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f, Description: %s, Category: %s\n", product.ID, product.Name, product.Price, product.Description, product.Category)
	}
}

//pasarle un int como parametro y que retorne el producto correspondiente al parametro pasado
func getById(id int) *Product {
	for _, product := range Products {
		if product.ID == id {
			return &product
		}	
	}
	return nil
}