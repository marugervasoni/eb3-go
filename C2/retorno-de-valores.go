package main

import "fmt"

func main() {
	
	fmt.Println("suma:",operacionAritmetica(6,2, Suma))
	fmt.Println("resta:", operacionAritmetica(6,2, Resta))
	fmt.Println("multiplicación:",operacionAritmetica(6,2, Multip))
	fmt.Println("división:", operacionAritmetica(6,2, Divis))

}

const (
	Suma = "+"
	Resta = "-"
	Multip = "*"
	Divis = "/"
)

func operacionAritmetica(valor1, valor2 float64, operador string) float64 {
	switch operador {
	case Suma:
		return valor1 + valor2
	case Resta:
		return valor1 - valor2
	case Multip:
		return valor1 * valor2
	case Divis:
		return valor1 / valor2
	}
	return 0
}
