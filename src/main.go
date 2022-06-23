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

	//Numeros enteros
	//int = Depende del OS (32 o 64 bits)
	//int8 = 8bits = -128 a 127
	//int16 = 16bits = -2^15 a 2^15-1
	//int32 = 32bits = -2^31 a 2^31-1
	//int64 = 64bits = -2^63 a 2^63-1

	//Optimizar memoria cuando sabemos que el dato simpre va ser positivo
	//uint = Depende del OS (32 o 64 bits)
	//uint8 = 8bits = 0 a 127
	//uint16 = 16bits = 0 a 2^15-1
	//uint32 = 32bits = 0 a 2^31-1
	//uint64 = 64bits = 0 a 2^63-1

	//numeros decimales
	// float32 = 32 bits = +/- 1.18e^-38 +/- -3.4e^38
	// float64 = 64 bits = +/- 2.23e^-308 +/- -1.8e^308

	//textos y booleanos
	//string = ""
	//bool = true or false

	//numeros complejos
	//Complex64 = Real e Imaginario float32
	//Complex128 = Real e Imaginario float64
	//Ejemplo : c:=10 + 8i

	//Explicacion paquete fmt
	//declaracion de variables
	helloMessage := "Hello"
	worldMessage := "world"

	//Println
	//Imprime en lineas diferentes
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	//Printf
	//Imprime las lineas de manera mas especifica, indicandole el tipo de dato que va a recibir el mensaje que va a mostrar
	//%s es para string, %d para entero, %v es para cualquier valor y el \n es para salto de linea
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos)

	//Sprintf
	//Genera un string pero no lo genera, solo lo guarda como string
	message := fmt.Sprintf("%s tiene mas de %d cursos", nombre, cursos)
	println(message)

	//Tipo de dato que tiene una variable con la opcion : %T

	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)

}
