package main

//LOS CHANNELS PERMITEN COMPARTIR DATOS ENTRE LAS GOROUTINES, manejan de forma nativa la comunicacion.
//los Channels solo se pueden manejar un tipo de dato. y este es asi mismo un una especie de conducto

import "fmt"

func say(text string, c chan<- string) { // canal de entrada de datos
	c <- text
}

func printChannelOutput(c <-chan string) { // canal de salida de datos
	var output string
	output = <-c
	fmt.Println(output)
}

func main() {
	// c := make(chan string) // dinamically accepts goroutines
	c := make(chan string, 1) // one goroutine at time
	fmt.Println("Hello")
	go say("Bye", c)

	//fmt.Println(<-c)
	printChannelOutput(c)
}
