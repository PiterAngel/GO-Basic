package main

import "fmt"

func main() {
	mi_funcion()
	saludar("Piter")
	sumar(5, 15)
	nuevasuma := sumar(100, 20)
	fmt.Println(nuevasuma)
}

func mi_funcion() {
	fmt.Println("Esta es una línea de código")
	fmt.Println("Esta es otra línea de código")
}

func saludar(nombre string) {
	fmt.Println("Hola " + nombre)
}

func sumar(numero1 int, numero2 int) int {
	suma := numero1 + numero2
	fmt.Println(suma)
	return suma
}
