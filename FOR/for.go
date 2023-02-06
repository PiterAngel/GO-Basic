package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	nombre := "Piter"
	fmt.Println(string(nombre[0]))

	for i := 0; i < 5; i++ {
		fmt.Println(i, string(nombre[i]))
	}

}
