package main

import "fmt"

func main() {
	//                  0        1           2
	frutas := []string{"Pera", "Manzana", "Naranja"}
	fmt.Println(frutas[0])

	frutas = append(frutas, "Melocotón", "Melón")
	frutas[0] = "Sandía"

	for i := 0; i < len(frutas); i++ {
		fmt.Println(i, frutas[i])
	}

	for i := 0; i < len(frutas); i++ {
		if frutas[i] == "Melocotón" {
			fmt.Println("Hay una coincidencia en la lista")

		}
	}
}
