package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
	edad     int
}

func (p persona) saludar(saludo string) {
	fmt.Println(saludo + ", " + p.nombre)

}

func (pers persona) cumple() int {
	return pers.edad + 1

}

func main() {
	persona1 := persona{nombre: "Piter", apellido: "Valiente", edad: 23}
	persona2 := persona{nombre: "Angel", apellido: "De León,", edad: 26}

	fmt.Println("La persona1 se llama: ", persona1.nombre)
	fmt.Println("La persona2 se llama: ", persona2.nombre)
	fmt.Println("Los datos completos de la persona1 es: ", persona1)

	persona1.edad = 24
	fmt.Println("La persona1 su edad cambio de 23 a : ", persona1.edad, " años")

	persona1.saludar("Hello")
	persona2.saludar("Hola")

	fmt.Println(persona1.cumple())

	edad_persona2 := persona2.cumple()
	fmt.Println(edad_persona2)
}
