package main

const (
	catA = "A"
	catB = "B"
	catC = "C"
)

func main() {
	
	calculateSalary(120, catA)
	calculateSalary(120, catB)
	calculateSalary(120, catC)
}

//calcular el salario de los empleados de una empresa marinera, basandose en la cantidad de 
//horas trabajadas por mes y categoría

//funcion que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría,
//y que devuelva su salario
 

func calculateSalary(minutes int, category string) float64 {
	var salary float64
	var hours float64 = minutes / 60
	switch category {
		case catA:
			salary = (hours * 3000) * 0.5
		case catB:
			salary = (hours * 1500) * 0.2
		case catC:
			salary = hours * 1000
	}
	return salary


}
