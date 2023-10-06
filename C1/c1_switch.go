package main

import (
	"fmt"
)

func main() {
	meses := []string{"Enero", "Febrero", "Junio", "Agosto", "Abril"}
	ObtenerEstacion(meses)
}

func ObtenerEstacion(meses []string) {
	for _, mes := range meses {
		switch {
		case mes == "Enero" || mes == "Febrero" || mes ==  "Marzo":
			fmt.Println("Verano")
		case mes == "Abril" || mes == "Mayo" || mes == "Junio":
			fmt.Println("Otoño")
		case mes == "Julio" || mes == "Agosto" || mes == "Septiembre":
			fmt.Println("Invierno")
		case mes == "Octubre" || mes == "Noviembre" || mes == "Diciembre":
			fmt.Println("Primavera")
		default:
			fmt.Println("No existe este mes")
		} 
	}
}