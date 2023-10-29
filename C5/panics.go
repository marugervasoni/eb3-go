package main

import "fmt"

//función con implementación de recover() ante un panic
func isPair(num int){
	//ante la producción de un panic la función defer permitirá con recover, recuperar el valor del panic, asignarlo a err y evitará la interrupción abrupta del programa
	defer func(){
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()
	
	if (num % 2) != 0 {
		panic("no es un número par")
	}

	fmt.Println(num, "es un numero par")
}

func main() {
	
	num := 3

	isPair(num)

	fmt.Println("Ejecución completada.")
}