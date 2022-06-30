//le indicamos el nombre del paquete
package mispaquetes

import "fmt"

//agregamos el mismo struct, pero esta vez con mayuscula en las iniciales para indicar que es publico
// CarroPublico Carro con acceso publico
type CarroPublico struct {
	Marca string
	Year  int
}

//Struct private
//iniciales en minuscula lo deja ya en privado
type carroPrivado struct {
	modelo string
	color  string
}

//funciones publicas con la incial mayuscula, PARA LAS PRIVADAS ES LA PRIMERA minuscula
func ImprimirMensaje(text string) {
	defer fmt.Println(text, "Funcion publica xd")
}
