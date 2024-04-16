package main

import (
	"fmt"

	"rsc.io/quote"
)

// Declaracion de constantes

const pi float32 = 3.14

const (
	Domingo = iota + 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {

	fmt.Println("Hola mundo")
	fmt.Println(quote.Go())

	//  Declaracion de variables

	//var firstName, lastName string
	//var age int

	firstName, lastName := "Gonzalo", "Arin"
	age := 24

	fmt.Println(firstName, lastName, age)
	fmt.Println(Domingo, Jueves)
}
