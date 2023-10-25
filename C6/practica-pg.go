package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main(){

	var data = [][]string{
		{"ID_PRODUCTO","PRECIO","CANTIDAD"},
		{"1","1000","5"},
		{"2","5400","10"},
		{"3","15000","3"},
	}
	creaYEscribirArchivo("productos.csv",data)

}

//funcionalidad para guardar un archivo de texto (CSV) con info de productos comprados, separados por ;
	//debe tener id del producto, precio y cantidad

func creaYEscribirArchivo(name string, data [][]string) {
	file, err := os.Create(name)
	checkError("Cannot create file", err)
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	for _, value := range data {
		err := w.Write(value)
		checkError("Cannot write to file", err)
	}
}
func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
