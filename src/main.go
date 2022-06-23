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

	//Operadores aritmeticos

	x := 10
	y := 50

	//Suma
	result := x + y
	fmt.Println("Suma: ", result)

	//Resta
	result = x - y
	fmt.Println("Resta: ", result)

	//Multiplicacion
	result = x * y
	fmt.Println("Multiplicacion: ", result)

	//Division
	result = y / x
	fmt.Println("Division: ", result)

	//Modulo
	result = y % x
	fmt.Println("Modulo: ", result)

	//Incremental
	x++
	fmt.Println("Incremental: ", x)

	//Decremental
	x--
	fmt.Println("Decremental: ", x)

	//Retos
	//Rectangulo, trapecio y de un circulo

	//Area de rectangulo
	largo := 20
	ancho := 10
	areaRectangulo := largo * ancho
	fmt.Println("El area del rectangulo es: ", areaRectangulo)

	//Area de un trapecio
	//la suma de las bases por la altura divido dos
	base1Trapecio := 4
	base2Trapecio := 10
	alturaTrapecio := 4
	areaTrapecio := ((base1Trapecio + base2Trapecio) * alturaTrapecio) / 2
	fmt.Println("El area de unt trapecio es: ", areaTrapecio, "Centimetros cuadrados")

	//Area de un circulo
	//En la linea 7 ya contamos con el valor de pi
	//Tener en cuenta el tipo de dato, float, entero etc
	radioCirculo := 5.0
	areaCirculo := (pi * (radioCirculo * radioCirculo))
	fmt.Println("El area de un circulo es: ", areaCirculo, "Centimetros cuadrados")

}
