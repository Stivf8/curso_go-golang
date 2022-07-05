//paquete para interfaces
package mispaquetes

import "fmt"

//Interfaces por lo general se utilizan para centralizar los metodos o funciones que pueden estar en varios Structs
//Calcular Area de Cuadrado y rectangulo
//SE AGREGA LA INTERFAZ QUE RECIBE LAS DOS FUNCIONES O METODOS "Area"
type figuras2D interface {
	//se le indica que funcion va a recibir y que tipo de dato
	Area() float64
}

//agregamos la funcion que estara encargada de procesar la interfaz Figuras2D
func Calcular(F figuras2D) {
	//le decimos que ejecute la funcion Area()
	fmt.Println("Area: ", F.Area())
}
