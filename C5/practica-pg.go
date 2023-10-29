package main

import (
	"fmt"
	"os"
)

// Ejercicio 1 - Impuestos de salario
// En la función main, definir una variable salary y asignarle un valor de tipo “int”. Luego, crear un error personalizado con un struct que implemente Error()
// con el mensaje “Error: el salario ingresado no alcanza el mínimo imponible” y lanzarlo en caso de que salary sea menor a 150.000. De lo contrario, imprimir 
// por consola el mensaje “Debe pagar impuesto”.

type salaryError struct {
	message	string
}
func (e salaryError) Error() string {
	return e.message
}

// Ejercicio 2 - Datos de clientes
// Un estudio contable necesita acceder a los datos de sus empleados para poder realizar distintas liquidaciones. Para ello, cuentan con todo el detalle necesario 
// en un archivo TXT. 
// Desarrollar el código necesario para leer los datos de un archivo llamado “customers.txt”. Sin embargo, debemos tener en cuenta que la empresa no nos ha pasado
// el archivo a leer por el programa. Dado que no contamos con el archivo necesario, se obtendrá un error. En tal caso, el programa deberá arrojar un panic al intentar
// leer un archivo que no existe, mostrando el mensaje “El archivo indicado no fue encontrado o está dañado”.

// Más allá de eso, deberá siempre imprimirse por consola “Ejecución finalizada”.

func main(){

	// 1)
	salary := 100000

	//Cuando el salario es menor a 150,000, creamos una instancia de esta estructura con el mensaje de error correspondiente y la lanzamos
	if salary < 150000 {
		err := salaryError{"Error: el salario ingresado no alcanza el mínimo imponible"}
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
	mensaje()

	//2
	defer func() {
        if err := recover(); err != nil {
			fmt.Println("Recover",err)
			mensaje()
        }
    }()
	_, err := os.Open("customers.txt")

	if err != nil {
		panic(err)	
	}
}
func mensaje() {
	fmt.Println("Ejecución finalizada")
}