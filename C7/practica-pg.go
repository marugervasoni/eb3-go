package main

import (
	"fmt"
//	"time"
)

// Calcular precio
// Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento. Para ello requieren realizar un programa que se 
// encargue de calcular el precio total de Productos, Servicios y Mantenimientos. Debido a la fuerte demanda, y para optimizar la velocidad, 
// requieren que el cálculo de la sumatoria se realice en paralelo mediante tres goroutines.
// Se requieren tres estructuras:
// Productos: nombre, precio, cantidad.
// Servicios: nombre, precio, minutos trabajados.
// Mantenimiento: nombre, precio.
// Y se requieren tres funciones:
// Sumar Productos: recibe un array de producto y devuelve el precio total (precio por cantidad).
// Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio por media hora trabajada. Si no llega a trabajar 30 minutos 
// se le cobra como si hubiese trabajado media hora).
// Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.
// Las tres se deben ejecutar concurrentemente y, al final, se debe mostrar por pantalla el monto final (sumando el total de las tres).
type Product struct {
	Name 	string
	Price 	float64
	Amount 	float64
}
type Service struct {
	Name 			string
	Price			float64
	MinutesWorked	float64
}
type Maintenance struct {
	Name 	string
	Price 	float64
}

func addProducts(products []Product, result chan float64) {
	var totalPrice float64 = 0
	for _, product := range products {
		totalPrice += product.Price * product.Amount		
	}
	result <- totalPrice
}
func addServices(services []Service, result chan float64) {
	var totalPrice float64 = 0
	for _, service := range services {
		billableTime := service.MinutesWorked
		if  billableTime < 30 {
			billableTime = 30
		} 	
		totalPrice += service.Price * billableTime / 30.0
	}
	result <- totalPrice
}
func addMaintenance(maintenances []Maintenance, result chan float64) {
	var totalPrice float64 = 0
	for _, maintenance := range maintenances {
		totalPrice += maintenance.Price		
	}
	result <- totalPrice
}


func main() {

	p1 := Product{"Product1", 1000,5}
	p2 := Product{"Product2", 500,2}

	s1 := Service{"Service1", 150,27}
	s2 := Service{"Service2", 200,120}

	m1 := Maintenance{"Maintenance1", 350}
	m2 := Maintenance{"Maintenance2", 425}

	products := []Product{p1,p2}
	services := []Service{s1,s2}
	maintenances := []Maintenance{m1,m2}

	//channels for each totalPrice result
	pch := make(chan float64)
	sch := make(chan float64)
	mch := make(chan float64)

	go addProducts(products, pch)
	go addServices(services, sch)
	go addMaintenance(maintenances, mch)

	var finalMount float64
	finalMount = <-pch + <-sch + <-mch
	//fmt.Printf("Total Price is: \n -Products \t $ %f \n -Services \t $ %f \n -Maintenances \t $ %f \n The total is: $ %f", pch,sch,mch,finalMount)
	fmt.Printf("Total Price is: $ %f", finalMount)
}
