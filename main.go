package main

import (
	"fmt"
	"math"
)

func main() {

	var lado1, lado2 float64

	//entrada de datos
	fmt.Print("ingresa Lado 1:")
	fmt.Scanln(&lado1)
	fmt.Print("ingresa Lado 2:")
	fmt.Scanln(&lado2)

	// proceso
	area := (lado1 * lado2) / 2
	hipotenusa := math.Sqrt(math.Pow(lado1, 2) + math.Pow(lado2, 2))
	perimetro := lado1 + lado2 + hipotenusa

	// salida de datos
	fmt.Printf("\n Area: %.2f", area)

	fmt.Printf("\n Perimetro: %.2f", perimetro)
}
