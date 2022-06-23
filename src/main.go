package main

import "fmt"

func main() {
	//Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.15

	fmt.Println("pi: ", pi, "pi2: ", pi2)

	//Declaracion de variables enteras que no estaba creada
	base := 12
	//Declaracion de variable que ya estaba creada
	base = 13
	//Declaracion de variable enteras
	var altura int = 15
	var area int

	fmt.Println("base:", base, "altura: ", altura, "area: ", area)

	//Zero values o valores por defecto
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Calcular area de un cuadrado

	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("El area del cuadrado es igual a: ", areaCuadrado)

}
