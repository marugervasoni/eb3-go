package main

import  (
	"fmt"
	"sync"
)
// Se requiere un programa para calcular si un número es par o impar desde una lista de int. Para esto se deben crear dos funciones:
// par(chan int) debe recibir a través de un channel de int los valores pares enviados y mostrarlos por la terminal.
// impar(chan int) debe recibir a través de un channel los valores impares y mostrarlos por terminal.
// El programa debe comenzar dos goroutines que se ejecuten concurrentemente. En la goroutine principal se itera el slice provisto con 
// valores y, en caso de que el valor sea par, se lo debe enviar al channel correspondiente. De igual manera para los impares.
// Ayuda: Si el resto de una división por dos da cero es un número par (i%2 == 0) y si da 1 es impar (i%2 == 0 ). Para obtener el resto 
// de una operación, utilizar el operador %: i%2 == 1.
var (
	wg sync.WaitGroup
)

func main() {

	numbers := []int{1,2,3,4,27,28,28,10,11}

	pairCh := make(chan int)
	oddCh := make(chan int)

	wg.Add(2)
	go pair(numbers, pairCh)
	go odd(numbers, oddCh)

	fmt.Println("Pair numbers:")
	for num := range pairCh {
		fmt.Println(num)
	}

	fmt.Println("Odd numbers:")
	for num := range oddCh {
		fmt.Println(num)
	}
}

func pair(numbers []int, pairCh chan int) {
	defer wg.Done()
	for _, num := range numbers {
		if num % 2 == 0 {
			pairCh <-num
		} 
	}
	close(pairCh)
} 
func odd(numbers []int, oddCh chan int) {
	defer wg.Done()
	for _, num := range numbers {
		if num % 2 != 0 {
			oddCh <-num
		}
	}
	close(oddCh)
} 