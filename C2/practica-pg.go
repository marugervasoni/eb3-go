package main

import "fmt"

func main() {

	sueldoEmp1 := 55000.00
	sueldoEmp2 := 155000.00
	s1 := calcularImpuesto(sueldoEmp1)
	s2 := calcularImpuesto(sueldoEmp2)
	fmt.Println("Ejercicio 1")
	fmt.Println("Empleado 1: tiene sueldo de: $",sueldoEmp1,"y sufrirá un descuento de: $",s1)
	fmt.Println("Empleado 2: tiene sueldo de: $",sueldoEmp2,"y sufrirá un descuento de: $",s2)

	e1 := calcularPromedio(9, 9, 10, 9, 9)
	e2 := calcularPromedio(6, 6, 7, 8, 8, 7, 7)
	fmt.Println("Ejercicio 2")
	fmt.Println("Estudiante 1: tiene un promedio de:", e1)
	fmt.Println("Estudiante 2: tiene un promedio de:", e2)
}

// Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento 
// de depositar el sueldo. Para cumplir el objetivo es necesario crear una función que 
// devuelva el impuesto de un salario, teniendo en cuenta que si la persona gana más de
// $50.000 se le descontará un 17 % del sueldo y si gana más de $150.000 se le descontará,
// además, un 10 % (27 % en total).
func calcularImpuesto(sueldo float64) float64 {
	var impuesto float64
	if sueldo > 50000 {
		impuesto = sueldo * 17 / 100
	} else if sueldo > 150000 {
		impuesto = sueldo * 27 / 100
	} else {
		impuesto = 0
	}
	return impuesto
}


// Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones. Se solicita
// generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio. 
// No se pueden introducir notas negativas.
func calcularPromedio(notas ...int) int {
	var total int
	for _, nota := range notas {
		total += nota
	}
	return total / len(notas)
}
