package main

//de esta manera importamos paquetes, indicando SOLO la carpeta
import (
	mispaquetes2 "curso_go-golang-main/src/mispaquetes"
	"fmt"
)

func main() {
	//Instanciamos el struct con la variable myPC que seria el objeto o clase
	//igualmente al instanciar le agregamos parametros
	myPC := mispaquetes2.Computador{Ram: 16, Disco: 200, Marca: "MSI"}
	fmt.Println(myPC)
	//se le pasa el parametro en cuanto sera aumentada la Ram
	myPC.AumentarRam(10)

	fmt.Println(myPC)
}
