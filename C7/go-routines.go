package main

import (
	"fmt"
	"time"
)

func main() {

	go contador("Oveja")
	go contador("Pez")

	//añadir espera de 2 segundos para que se complete ejecución de rutinas
	time.Sleep(time.Second * 2)
}

func contador(elem string) {
	for i := 1; true; i++ {
		fmt.Println(i, elem)
		time.Sleep(time.Second / 2) //1/2 segundo
	}
}