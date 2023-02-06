package main

import "fmt"

func main() {
	numero := 455
	if numero == 40 {
		fmt.Println("es 40 xd")

	} else if numero == 45 {
		fmt.Println("es 45 xd")
	} else {
		fmt.Println("Error, es diferente xd")
	}

	edad := 10
	if edad >= 18 {
		fmt.Println("Eres mayor de edad xd")
	} else if edad < 18 && edad >= 0 {
		fmt.Println("Eres menor de edad xd")
	} else {
		fmt.Println("La edad no es válida xd")
	}
}
