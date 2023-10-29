package main

import "fmt"

func main() {

	//POINTERS 
	//creation of pointer: *+dataType
	var p *int
	
	//creation of variable
	var v int = 50

	//asignation of variable v direction to variable p
	p = &v

	//direction of memory of pointer p:
	fmt.Println("direction of memory in pointer p: ", p)


	//direction operator: & + variable
	fmt.Println("direction in memory of variable v is: ", &v)

	//desreference operator: * + variable
	fmt.Println("with desreference of a variable, I have the value of my variable p (who has v direction): ", *p)

}