package main

import "fmt"

type Alumno struct {
	Nombre		string
	Apellido	string
	Edad 		int
	DNI 		int
	Fecha 		string
}

func main(){
	p := Alumno { 
		"María", "Gervasoni", 27, 39517575, "01/07/2021",
	}
	p.Detalle()

}

func (a *Alumno) Detalle(){
	fmt.Printf("%s %s tiene %d años de edad, Su DNI es %d, e ingresó en fecha %s", a.Nombre, a.Apellido, a.Edad, a.DNI, a.Fecha)
}