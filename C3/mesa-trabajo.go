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

var Products = []Product{
    {ID: 1, Name: "Puma Run", Price: 75000.00, Description: "Running Sneakers", Category: "Woman"},
    {ID: 2, Name: "Nike 124", Price: 18900.00, Description: "Sport Cap, black color", Category: "No gender"},
    {ID: 3, Name: "Adi zero", Price: 29000.00, Description: "T-shirt Adidas Originals Zero", Category: "Men"},
}

func main() {


}
//Metodo que debera tomar slice products y añadir el producto desde el cual se llama el metodo
func (p* Product) Save(){

}
//deberá imprimir todos los productos guardados en slice products
func (p* Product) GetAll(){
	
}
//pasarle un int como parametro y que retorne el producto correspondiente al parametro pasado
func getById() {

}