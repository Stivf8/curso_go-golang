package mispaquetes2

// Computador struct publico
// recordar las iniales para un struct y para sus variables para que sean publicos deben ser con Mayuscula
type Computador struct {
	Ram   int
	Disco int
	Marca string
}

func (myPC *Computador) AumentarRam(vRam int) {
	myPC.Ram = myPC.Ram + vRam
}
