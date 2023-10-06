package main

import "fmt"

func main() {
	
	s,r,m,d := operaciones(6,2)

	fmt.Println("Suma:\t\t",s)
	fmt.Println("Resta:\t\t",r)
	fmt.Println("Multiplicación:\t",m)
	fmt.Println("División:\t",d)
}

func operaciones(valor1, valor2 float64) (float64, float64, float64, float64){
	suma := valor1 + valor2
	resta := valor1 - valor2
	multip := valor1 * valor2
	var divis float64

	if valor2 != 0 {
		divis = valor1 /valor2
	}

	return suma, resta, multip, divis
}
