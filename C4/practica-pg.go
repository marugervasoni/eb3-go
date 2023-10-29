package main

import "fmt"

//CONSIGNA
// Una empresa de redes sociales requiere implementar una estructura usuarios con funciones que vayan agregando información a la misma. Para optimizar y ahorrar memoria
// requieren que la estructura usuarios ocupe el mismo lugar en memoria para el main del programa y para las funciones. La estructura debe tener los campos: 
//nombre, apellido, edad, correo y contraseña. Y deben implementarse las funciones:
// cambiarNombre: permite cambiar el nombre y apellido.
// cambiarEdad: permite cambiar la edad.
// cambiarCorreo: permite cambiar el correo.
// cambiarContraseña: permite cambiar la contraseña.

type User struct {
	Name 		string
	Lastname 	string
	Age 		int
	Email 		string
	Password	string
}

func (u User) details(){
	fmt.Printf("--USER: %s %s \n Age: %d \n Email: %s \n Password: %s \n", u.Name, u.Lastname, u.Age, u.Email, u.Password)
}
func (u *User) changeName(newName, newLastname string) {
	u.Name = newName
	u.Lastname = newLastname
}
func (u *User) changeAge(newAge int) {
	u.Age = newAge
}
func (u *User) changeEmail(newEmail string) {
	u.Email = newEmail
}
func (u *User) changePass(newPassword string) {
	u.Password = newPassword
}

func main(){

	user := User{
		Name: "Mariana",
		Lastname: "Gervasoni",
		Age: 21,
		Email: "gervasonim@gmail.com",
		Password: "123456",
	}
	user.details()
	user.changeAge(27)
	user.changeName("Maria Rosa", "Gervasoni")
	user.changeEmail("mariar.gervasoni@gmail.com")
	user.changePass("password1")
	user.details()
}