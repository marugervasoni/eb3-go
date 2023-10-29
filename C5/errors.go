package main

import (
	"fmt"
	"errors"
)

type error interface {
	Error()
}

//FORMAS DE PERSONALIZAR E IMPLEMENTAR ERRORES

// A) Error()
//implementación metodo Error:
// func (e *MyError) Error() string {
// 	return "My error info"
// }

// B) errors
//methods:
//errors.New("esto es un error")
//- crear nuevo error con un msj
//-recibe por parametro un msj de error string personalizable
//errors.Unwrap(err error) error
//-desenvolver errores dentro de cadena err
//errors.As(err error, target interface{}) bool
//- interfaz vacia nos permite pasar puntero a un tipo error y retorna variable bool
//-compara dentro cadena err con un puntero a un tipo de errores, al encontrar la primer coincidencia, asigna a target el error y devuelve true
//errors.Is(err, target error) bool
//- devuelve true si encuentra coincidencia dentro de cadena de errores con el error target

// C) fmt.Errorf()
//ejemplo:
// err := fmt.Errorf("Error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)
// fmt.Println(err)

func main() {
	
	//implementación errors.New()
	statusCode := 404
	
	//validamos statusCode
	if statusCode > 399 {
		fmt.Println(errors.New("La petición ha fallado"))
		return
	}
	fmt.Println("El programa finalizó correctamente")

}