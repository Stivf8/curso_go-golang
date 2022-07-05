package main

import "fmt"

//Interfaces por lo general se utilizan para centralizar los metodos o funciones que pueden estar en varios Structs
//calcular area de cuadrado y rectangulo
//SE AGREGA LA INTERFAZ QUE RECIBE LAS DOS FUNCIONES O METODOS "area"
type figuras2D interface {
	//se le indica que funcion va a recibir y que tipo de dato
	area() float64
}

//struct o clase cuadrado privada
type cuadrado struct {
	base float64
}

//struct de area del rectangulo
type rectangulo struct {
	base   float64
	altura float64
}

//Funcion que calcula el area del rectangulo, se referencia el struct cuadrado en la variable r TIENE EL MISMO NOMBRE QUE EL AREA DE CUADRADO -> area
func (r rectangulo) area() float64 {
	return r.base * r.altura
}

//Funcion que calcula el area del cuadrado, se referencia el struct cuadrado en la variable c TIENE EL MISMO NOMBRE QUE EL AREA DE CUADRADO -> area
func (c cuadrado) area() float64 {
	return c.base * c.base
}

//agregamos la funcion que estara encargada de procesar la interfaz figuras2D
func calcular(f figuras2D) {
	//le decimos que ejecute la funcion area()
	fmt.Println("Area: ", f.area())
}

//funcion principal que se ejecuta
func main() {
	//instanciamos structs en nueva variable y definimos que el valor base de la clase cuadrado va a ser 2
	miCuadrado := cuadrado{base: 2}
	miRectangulo := rectangulo{base: 2, altura: 4}
	//Para ejecutar la interfaz para las dos funciones o metodos que comparten y retonran el mismo valor, se utiliza la funcion que procesa la interfaz calcular
	//como parametro le pasamos la instancia del struct o clase, lo cual es mas optimo
	calcular(miCuadrado)
	calcular(miRectangulo)

	//lista de interfaces
	//go en las lsitas o slides o arrays necesita definir el parametro que va a contener esa lista, int o string etc, no varios tipos de datos en una misma
	//la manera de simular una lista con multiples tipos de datos es: []interface{} al definir el slides
	miinterfaz := []interface{}{"Hola", 12, 4.2}
	//imprimimos el slides con ... (toma cada uno de los elementos e imprime de manera individual)
	fmt.Println(miinterfaz...)
}
