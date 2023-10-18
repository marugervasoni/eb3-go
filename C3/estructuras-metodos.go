package main

import "fmt"

//declaración Estructura
type Persona struct {
	Nombre	string
	Edad 	int
}

func main(){
	//declaramos variable de tipo Persona
	p := Persona {
		Nombre:	"Gabriel",
		Edad:	24,
	}
	p.Descripcion()

}

//declaración Método
func (p *Persona) Descripcion(){
	fmt.Printf("%s tiene %d años de edad", p.Nombre, p.Edad)
}