package main

import "fmt"

//Range, Close y Select en channels
//como acceder a los valores de los channels y como manejar multiples channels

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	// go func()... se lanza de manera concurrente:
	// la línea fmt.Println(<-c) queda en espera hasta que haya algo que leer de <-c.
	// Al ser lanzado concurrentemente, mientras dicha línea espera, se va a la ejecución de fibonacci(c, quit),
	// que guardará valores en c <- x tantas veces se produzca la espera de la línea fmt.Println(<-c).
	// Una vez sale del bucle, pasa a esperar case <-quit, que será satisfecho con quit <- 0, por lo que imprimirá "quit" y finalizará
	// el programa.
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

func main2() {

	//primer canal
	//maneja dos datos al mismo tiempo

	c := make(chan string, 2)
	//insertamos mensaje 1
	c <- "mensaje 1"
	//insertamos mensaje 2
	//c <- "mensaje 2"

	//len lo que hace es decir cuantos datos hay dentro de x, en este caso channel, en este caso dira cuantas goRoutines tiene el channel
	//cap lo que hace es cuanta es la cantidad maxima que puede almacenar el channel

	fmt.Println(len(c), cap(c))

	//como hacer el recorrido de cada uno de los datos de los canales
	//Range y Close
	//Close lo que hace es decirle al RunTime de Go que cierre el canal, y quiere decir que ese canal no va a recibir mas datos, a pesar de que pueda.
	//lo ideal es cerrar el canal cuando ya no se quieran utilizar y sabes que no vas a guardar datos
	close(c)
	//range ayuda a hacer el recorrido de una lista o en este caso el channel
	for message := range c {
		fmt.Println(message)
	}
	//cuando manejamos multiples canales, y no se tiene certeza de cual de los canales va a responder primero, utilizamos Select
	email1 := make(chan string, 1)
	email2 := make(chan string, 1)
	//invocamos la funcion mensaje y la enviamos en una goRoutine, le enviamos el mensaje o Email y le enviamos el canal 1 o email1
	//en este punto no sabemos cual se va a ejecutar primero entonces usamos Select
	go mensaje("mensaje 1", email1)
	go mensaje("mensaje 2", email2)
	//Select
	// se debe conocer la cantidad de datos que el channel va a manejar por lo que se utiliza un for
	// indice va a ser 0 y que i va a ser menor a 2 porque tenemos ejecucion canal 1 canal 2, y cada iteracion suma 1
	for i := 0; i <= 2; i++ {
		select {
		//en caso de que mensaje 1 venga de email 1
		case m1 := <-email1:
			fmt.Println("Email recibido de email1", m1)
			//en caso de que mensaje 2 venga de email 2
		case m2 := <-email2:
			fmt.Println("Email recibido de email2", m2)
		}
	}

}

//funcion que guarda un mensaje en un canal
//recibe dato tipo string y un dato de tipo canal
func mensaje(text string, c chan string) {
	c <- text
}
