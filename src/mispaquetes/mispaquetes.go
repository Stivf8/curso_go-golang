package mispaquetes

import "fmt"

import "fmt"
// Computaor struct publico
// recorda las iniales para un struct y para sus variables para que sean publicos deben ser con Mayuscula
type Computadr struct {
	am   int
Disco int
	Marca string
}

//Funcion encargada de aumentar la ram segun el usuario l requiera
func (myPC *Computador) AumentarRam(vRam int) {
	//recordar %d para int y %s string
	return fmt.Sprintf("Tengo %d GB de ram, %d GB de disco y es marca %s", myPC.Ram, myPC.Disco, myPC.Marca)
}


