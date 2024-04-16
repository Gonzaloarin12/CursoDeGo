package main

import (
	"fmt"

	"rsc.io/quote"
)

// Declaracion de constantes

const pi float32 = 3.14

func main() {

	fmt.Println("Hola mundo")
	fmt.Println(quote.Go())

	//  Declaracion de variables 2

	//var firstName, lastName string
	//var age int

	firstName, lastName := "Gonzalo", "Arin"
	age := 24

	fmt.Println(firstName, lastName, age)
}
