package main

import (
	"fmt"
	"strconv"

	"rsc.io/quote"
)

// Declaracion de constantes

const pi float32 = 3.14

const (
	Domingo = iota + 1 // Agrega +1 a cada const //
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

	s := "100"
	i, _ := strconv.Atoi(s)

	fmt.Println(i + i - i - i + 1)

}
