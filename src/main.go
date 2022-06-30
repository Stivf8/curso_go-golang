package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

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

	//llamo la funcion normalFunction
	normalFunction("Stiven", "Ok!")

	//llamo funcion returnValue
	respuesta := returnValue(2)
	fmt.Println("Desde funcion returnValue retorna: ", respuesta)

	//llamo funcion doubleReturn que retorna dos valores, estos los asigno a dos variables
	value1, value2 := dobleReturn(2)
	//para tomar un solo valor asi la funcion te devuelva dos se realiza de la siguiente manera: (con raya al piso)
	//value1, _ := dobleReturn(2)
	println("retornando uno o varios datos desde funcion dobleReturn", value1, value2)

	//Ciclos en Go solo el for condicional

	for i := 0; i <= 10; i++ {
		fmt.Println("Soy el ciclo for condicional: ", i)
	}

	//este disminuye el i

	for i := 10; i >= 0; i-- {
		fmt.Println("Soy el ciclo for condicional decayendo: ", i)
	}

	// For While (hasta que una condicion se cumpla)
	counter := 0
	for counter <= 10 {
		fmt.Println("Soy ciclo for While", counter)
		counter++
	}

	//For forever quiere decir que va a ser iterando hasta la eternidad

	counterForever := 0
	for {
		fmt.Println("counterFoverever es igual a:", counterForever)
		counterForever++
		//esta condicional se utiliza para salir del forever
		if counterForever == 20 {
			break
		}
	}

	//Condicional IF en GO
	valor1 := 1
	if valor1 == 1 {
		fmt.Println("Soy condicional IF es igual a 1")
	} else {
		fmt.Println("No es igual a 1")
	}

	//Con AND
	valor2 := 2
	if valor1 == 1 && valor2 == 3 {
		fmt.Println("El AND es VERDAD")
	} else {
		fmt.Println("El AND es FALSO")
	}

	//Con OR
	if valor1 == 1 || valor2 == 2 {
		fmt.Println("Soy OR, alguna condicion cumplio")
	} else {
		fmt.Println("Soy OR, ninguna de las condiciones funciono")
	}

	//Convertir texto a numero
	value, err := strconv.Atoi("53")
	//nil es la manera en go para validar si la ejecucion no tuvo errores
	if err != nil {
		//Mostramos el log de que paso, que error ocurrio
		log.Fatal(err)
	}
	fmt.Println("El valor convertido: ", value)

	// 1. se requiere crear una funcion de un numero que recibe
	// esa funcion es par o impar

	//Llamamos a la funcion imparOpar para saber si el numero 8 es par
	imparOpar(80)

	// 2. se requiere una funcion que reciba un usuario
	// un pasword y revise si es valido o no, se espera un retorno
	// true

	//Llamamos a la funcion LogIn para validar los datos del usuario ingresados
	resultadoLogIn := LogIn("manuel@hotmail.com", "manuelitoPte")
	println("El reslutado de inicio de sesion fue: ", resultadoLogIn)

	//Utilizacion de SWITCH cuando estamos ejecutando muchas ejecuciones de IF, es importante validar la posibilidad de usar SWITCH
	//puede ser representado de muchas maneras
	//esta esta:
	//modulo := 5 % 2
	//y la mas utilizada es esta: agregando en el switch toda la operacion
	switch modulo := 4 % 2; modulo {
	case 0:
		fmt.Println("Es Par")
		//en caso de que no se cumplan ninguno de los case impuestos anterior mente el default se ejecuta
	default:
		fmt.Println("Es impar")
	}

	//SWITCH sin condicion, se usa para una misma variable y si es necesario utilziar un if anidado

	value5 := 200
	switch {
	case value5 > 100:
		fmt.Println("es mayor a 100")
	case value5 < 0:
		fmt.Print("es mayor a 0")
	default:
		fmt.Println("no cumple ningun case")

	}

	//Defer, antes de morir la primera funcion va ejecutar la otra o todas las lineas que tenga despues
	//quedanto en este caso Mundo, Hola, pued eutilizarse para cerrar conexiones o para condiciones que si o si deben pasar pero al final, la buena practica es un defer por funcion
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	//Continue y break se utilizan en el for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 2 {
			fmt.Println("es 2")
			//CONTINUA EL CICLO CON NORMALIDAD, se utiliza para digamos un caso en que a pesar de que suceda x cosa, continue la ejecucion del codigo de la funcion por ejemplo
			//para despues procesarlo etc
			continue
		}
		//break es cuando pasa algo determinado, quiero que ese codigo ya no se ejecute
		if i == 8 {
			fmt.Println("es un Break")
			break
		}
	}

	//Array - Arreglos
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println("soy el array mi valor es:", array)
	//len, dice cuantos elementos hay en el arreglo
	//cap, indica la capacidad maxima de ese array
	fmt.Println(len(array), cap(array))

	//Slice, declaracion similiar al array pero no se le indica la cantidad de valores que va a tener
	slice := []int{1, 2, 3, 4, 5, 6}
	println("Slice", slice, len(slice), cap(slice))

	//se utiliza para manejar array slice o listas para poder interactuar con cada uno de los elementos
	//Metodos en el slice
	//Imprime elemento 0
	fmt.Println(slice[0])
	//Imprime hasta el elemento 3, es exclusivo el ultimo parametro
	fmt.Println(slice[:3])
	//desde el indice 2 al 4, excluye el 4 imprime hasta el 3
	fmt.Println(slice[2:4])
	//del 4 en adelante, el primer elemento si es inclusivo
	fmt.Println(slice[4:])
	//Como adicionar elementos al slice
	//Append
	slice = append(slice, 7)
	fmt.Println(slice)
	//agregar otra lista al slice, se debe utilizar el append, pero se le pasa la variable con ...
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

	//COMO RECORRER CADA UNO DE LOS ELEMENTOS DEL SLICE
	//definimos slice tipo string
	sliceString := []string{"hola", "que", "hace"}
	//recorremos el indice con un for y utilizamos el range para recorrelo todo
	//recordar, para omitir una variable, usamos _ en el caso lo pondriamos en la i y con esto lo omitiriamos
	for i, valor := range sliceString {
		fmt.Println("recorriendo el arreglo", i, valor)
	}
	//si se desea saber solo el indice, eliminamos valor y dejamos unicamente el i de indice, algo asi:
	/*
		for i := range sliceString {
			fmt.Println("recorriendo el arreglo", i)
		}*/

	//Ejercicio de si es Palindromo o no es palindromo (se lee de la misma manera de izquiera a derecha que de derecha a izquierda) como ama, y no palindromo amor o roma
	//Para esto llamamos a la funcion esPalindromo
	esPalindromo("Ama")

}

//Ejercicio de si es Palindromo o no es palindromo (se lee de la misma manera de izquiera a derecha que de derecha a izquierda) como ama, y no palindromo casas
//Para esto realizamos una funcion
func esPalindromo(text string) {
	var textAlrevez string
	//con esta linea de codigo nos aseguramos que lo que vamos a procesar este en minuscula para evitar inconvenientes si ingresan valores en mayuscula
	text = strings.ToLower(text)
	//for con sentido inverso
	for i := len(text) - 1; i >= 0; i-- {
		//el =+ es lo mismo que hacer textalrevez = textalrevez
		//aqui pasamos el string de manera alrevez
		textAlrevez += string(text[i])
	}

	if text == textAlrevez {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}

	//sentencia if, si el texto que le pasamos es igual al text pero en sentido contrario es palindromo

}

//Funcion que valida si es par o impar
func imparOpar(num1 int) {
	resultImparOpar := num1 % 2
	if resultImparOpar == 0 {
		fmt.Println("El numero", num1, "es PAR")
	} else {
		fmt.Println("El numero", num1, "es IMPAR")
	}

}

//Funcion login valida si el usario y contrasena ingresada son Ok!
//IMPORTANTE DECIR QUE VA A RETORNAR EN ESTE CASO UN BOOLEAN
func LogIn(user, password string) bool {
	//En teoria estos datos se validan en alguna base de datos, por lo tanto se generan las varialbes de DBpassword y DBuser
	DBpassword := "manuelitoPte"
	DBuser := "manuel@hotmail.com"

	if user == DBuser && password == DBpassword {
		resultadoLogIn := true
		return resultadoLogIn
	} else {
		resultadoLogIn := false
		return resultadoLogIn
	}

}

//FUNCIONES
//cuando son multiples entradas, pueden ser definidas de una vez tipo : a,b int, c, d string
func normalFunction(mensaje, confirmacion string) {

	fmt.Println("hola te saludo desde la funcion NormalFunction", mensaje, confirmacion)
}

//Funciones que retornen 1 valor
func returnValue(valor int) int {
	return valor * 2
}

//Funcion que retoner 2 valores o mas
func dobleReturn(valor1 int) (c, d int) {
	return valor1, valor1 * 2
}
