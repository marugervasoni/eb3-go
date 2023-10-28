package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	filename="./categorias.csv"
)
func main(){

	readfile(filename)

}

func readfile(name string) {

	file, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}

	//split para limpiar los salto de línea
	data := strings.Split(string(file), "/n")
	//var total float64

	fmt.Println(data)

	//teniendo data lo recorremos
	for i := 0; i < len(data) -1; i++ {
		line := strings.Split(string(data[i]), ",")

		fmt.Println(line)
	}

}