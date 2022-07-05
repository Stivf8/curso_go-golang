package mispaquetes

// Computador struct publico
// recordar las iniales para un struct y para sus variables para que sean publicos deben ser con Mayuscula
type Computador struct {
	Ram   int
	Disco int
	Marca string
}

//Funcion encargada de aumentar la ram segun el usuario lo requiera
func (myPC *Computador) AumentarRam(vRam int) {
	myPC.Ram = myPC.Ram + vRam
}
