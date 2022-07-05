package main

//importa carpeta o paquete donde estan nuestras herramientas
import (
	"curso_go-golang-main/src/mispaquetes"
	"fmt"
)

//Interfaces por lo general se utilizan para centralizar los metodos o funciones que pueden estar en varios Structs
//Calcular Area de Cuadrado y rectangulo
//SE AGREGA LA INTERFAZ QUE RECIBE LAS DOS FUNCIONES O METODOS "Area"

//funcion principal que se ejecuta
func main() {
	//instanciamos los structs
	pCuadrado := mispaquetes.Cuadrado{Base: 2}
	pRectangulo := mispaquetes.Rectangulo{Base: 2, Altura: 4}
	fmt.Println("Cuadrado: ", pCuadrado, "Rectangulo", pRectangulo)
	//utilizamos el metodo publico Calcular para acceder a la interfaz privada Figuras2D la cual ejecuta los metodos compartidos Area() publicos de los paquetes cuadrado y rectangulo
	fmt.Println("Cuadrado:")
	mispaquetes.Calcular(pCuadrado)
	fmt.Println("Rectangulo:")
	mispaquetes.Calcular(pRectangulo)
	//lista de interfaces
	fmt.Println("Lista de interfaces:")
	//En go en las lsitas o slides o arrays necesita definir el parametro que va a contener esa lista, int o string etc, no varios tipos de datos en una misma
	//la manera de simular una lista con multiples tipos de datos es: []interface{} al definir el slides
	miinterfaz := []interface{}{"Hola", 12, 4.2}
	//imprimimos el slides con ... (toma cada uno de los elementos e imprime de manera individual)
	fmt.Println(miinterfaz...)
}
