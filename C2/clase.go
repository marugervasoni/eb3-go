package main

import (
	"fmt"
	"errors"
)

const (
	minimum = "minimum"
	maximum = "maximum"
	average = "average"
)

func main(){

	minFunc, err := operation(minimum)
	maxFunc, err := operation(maximum)
	averageFunc, err := operation(average)

	if err != nil {
		fmt.Println("se produjo un error")
	}

	// se llaman a las funciones asignadas previamente a las variables minFunc, maxFunc, y averageFunc con conjuntos de números como argumentos
	minValues := minFunc(1,2,3,4,10,2,4,5)
	maxValues := maxFunc(2,3,4,10,2,4,5)
	averageValues := averageFunc(2,3,4,10,2,4,5)	

	fmt.Println("minimal value: ", minValues)
	fmt.Println("maximum value: ", maxValues)
	fmt.Println("average value: ", averageValues)

}

//funciones: 1 por cada operación
func minFunc(values...int) int {
	min := 100
	for _, value := range values {
		if value < min {
			min = value
		}
	}
	return min
}

func maxFunc(values...int) int {
	var max int
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return max
}

func averageFunc(values...int) int {
	var sum int
	for _, value := range values {
		sum += value
	}
	
	avg := sum / len(values)
	
	return avg
}

//función orquestadora de operaciones que a su vez retorna la función corresp a la operación indicada en la invocación de operation()
func operation(name string) (func(...int) int, error) {

	switch name {
	case minimum:
		return minFunc, nil
	case maximum:
		return maxFunc, nil
	case average:
		return averageFunc, nil
	default:
		return nil, errors.New("No se encontró la operación")
	}

}