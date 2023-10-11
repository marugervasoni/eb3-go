package main

import "fmt"

func main() {

	// 1 
	estudiantes := []string{"Benajmin","Nahuel","Brenda","Marcos","Pedro","Axel","Alez","Dolores","Federico","Hernán","Leandro","Eduardo","Duvraschka"}
	fmt.Println("Ejercicio 1------------------------------------")
	fmt.Println("A\t",estudiantes)
	estudiantes = append(estudiantes,"Gabriela")
	fmt.Println("B\t",estudiantes)

	// 2
	employees := map[string]int{"Benjamin":20, "Nahuel":26, "Brenda":19, "Darío":44, "Pedro":18}
	fmt.Println("Ejercicio 2------------------------------------")
	fmt.Println("A\t","La edad de Benjamin es: ",employees["Benjamin"])
	
	var countEmplOlder21 int

	for _, employee := range employees {
		if employee > 21 {
			countEmplOlder21++ 
		}	
	}
	fmt.Println("B\t","Cantidad de empleados mayores de 21: ",countEmplOlder21)

	employees["Federico"] = 25
	fmt.Println("C\t", employees)

	delete(employees, "Pedro")
	fmt.Println("D\t", employees)	

}