package main

import "fmt"

func main() {
	remera := 2000
	jean := 10000
	campera := 15000
	gorra := 5000 
	camisa := 12000
	fmt.Println("---------Ejercicio 1: -------") 
	obtenerDescuento(remera)
	obtenerDescuento(jean)
	obtenerDescuento(campera)
	obtenerDescuento(gorra)
	obtenerDescuento(camisa)

	//0: false; 1: true
	cliente1 := map[string]int{"edad": 25, "empleado":1, "añosAntiguedad":2, "sueldo": 95000}
	cliente2 := map[string]int{"edad": 17, "empleado":0, "añosAntiguedad":0, "sueldo": 0}
	cliente3 := map[string]int{"edad": 35, "empleado":1, "añosAntiguedad":10, "sueldo": 300000}
	fmt.Println("---------Ejercicio 2: -------")
	autorizarPrestamo(cliente1)
	autorizarPrestamo(cliente2)
	autorizarPrestamo(cliente3)
}

// Una tienda de ropa quiere ofrecer a sus clientes un descuento sobre sus productos.
// Para ello necesitan una aplicación que les permita calcular el descuento basándose 
// en dos variables: su precio y el descuento en porcentaje. La tienda espera obtener 
// como resultado el valor con el descuento aplicado y luego imprimirlo en consola.
func obtenerDescuento(producto int){
	descuento := producto * 10 / 100 
	producto -= descuento
	fmt.Println("El precio con descuento es: $",producto)
}

// Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los 
// mismos. El banco tiene ciertas reglas para saber a qué cliente se le puede otorgar: 
// solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados
//  y tengan más de un año de antigüedad en su trabajo. Dentro de los préstamos que otorga, 
//  no les cobrará interés a los que su sueldo sea mejor a $100.000.
// Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de 
// acuerdo a cada caso.
func autorizarPrestamo(cliente map[string]int) {
	if cliente["edad"] > 22 && cliente["empleado"] == 1 && cliente["añosAntiguedad"] > 1 {
		if cliente["sueldo"] > 100000 {
			fmt.Println("Felicidades, eres elegible para un préstamo sin interés.")
		} else {
			fmt.Println("Eres elegible para un préstamo, pero se aplicará interés.")
		}
	} else {
		fmt.Println("Lo sentimos, no cumples con los requisitos para un préstamo en este momento.")
	}
}
