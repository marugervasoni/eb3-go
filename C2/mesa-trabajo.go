package main

import "fmt"


func main() {
	
	catA := calculateSalary("A")
	salaryCatA :=  catA(160)
	catB := calculateSalary("B")
	salaryCatB :=  catB(160)
	catC := calculateSalary("C")
	salaryCatC :=  catC(160)
	
	fmt.Println("Salary for employee of category A is:", salaryCatA)
	fmt.Println("Salary for employee of category B is:", salaryCatB)
	fmt.Println("Salary for employee of category C is:", salaryCatC)
}

//calcular el salario de los empleados de una empresa marinera, basandose en la cantidad de 
//horas trabajadas por mes y categoría 
//funcion que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría,
//y que devuelva su salario
//(Categ A cobra 3000 por hora + un 50% adicional de su salario; B 1500 x hora + 20% adicional; y C 1000 por hora solamente)
func categoryA(hours float64) float64 {
	return (hours * 3000) * 1.5
}
func categoryB(hours float64) float64 {
	return (hours * 1500) * 1.2
}
func categoryC(hours float64) float64 {
	return  hours * 1000
}

func calculateSalary(category string) func(minutes float64) float64 {
	
	switch category {
	case "A":
		return categoryA
	case "B":
		return categoryB
	case "C":
		return categoryC
	}
	return nil
}
