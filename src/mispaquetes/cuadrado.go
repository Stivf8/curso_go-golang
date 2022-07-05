//Paquete para rectangulo
package mispaquetes

//struct o clase Cuadrado Publica
type Cuadrado struct {
	Base float64
}

//Funcion que calcula el Area del Cuadrado, se referencia el struct Cuadrado en la variable c TIENE EL MISMO NOMBRE QUE EL Area DE Cuadrado -> Area
func (c Cuadrado) Area() float64 {
	return c.Base * c.Base
}
