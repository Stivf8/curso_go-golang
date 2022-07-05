//Paquete para rectangulo
package mispaquetes

//Clase o struct para rectangulo

//struct de Area del rectangulo
type Rectangulo struct {
	Base   float64
	Altura float64
}

//Funcion que calcula el Area del rectangulo, se referencia el struct Cuadrado en la variable r TIENE EL MISMO NOMBRE QUE EL Area DE Cuadrado -> Area
func (r Rectangulo) Area() float64 {
	return r.Base * r.Altura
}
