package mispaquetes

import "fmt"

// Computador struct publico
// recordar las iniales para un struct y para sus variables para que sean publicos deben ser con Mayuscula
type Computador struct {
	Ram   int
	Disco int
	Marca string
}

//Funcion encargada de aumentar la ram segun el usuario lo requiera
func (myPC *Computador) AumentarRam(vRam int){
	myPC.Ram = myPC.Ram + vRam
}

//Se agrega funcion que trabaja el output del struc Computador
func (myPC Computador) String() string {
	//se agrega logica para trabajar la salida de la funcion
	//recordar %d para int y %s string
	return fmt.Sprintf("Tengo %d GB de ram, %d GB de disco y es marca %s", myPC.Ram, myPC.Disco, myPC.Marca)
}
