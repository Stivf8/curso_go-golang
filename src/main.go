package main

//go get para instalar paquetes externos
//ejemplo para instalar tour
//go install -v golang.org/x/website/tour@latest
//otro ejemplo
//-v nos muestra el output si no esta instalado, el -u para que siempre lo muestre asi este instalado
//go get -v -u golang.org/x/tour
//Echo un framework para desarrollo web en go
//Para instalar echo se realizo la ejecucion de los siguientes pasos
//ir a la carpeta donde queremos iniciar el servidor o servicio
//alli descargar el echo con el siguiente comando
//go get github.com/labstack/echo
//despues iniciar el modulo en esa misma carpeta con la siguiente instruccion
//buena practica indicar la ruta de github de tu repo, indicando igualmente el nombre de tu carpeta y repo al final
//go mod init github.com/Stivf8/curso_go-golang
// tambien estan los comandos go mod verify para verificar los modulos o go mod edit
//luego en la documentacion de echo esta como incializar
//https://echo.labstack.com/guide/
//pero aqui abajo ya tenemos una manera
//y para inicar el servicio es
//go run src/main.go en este caso
//GO MODULES SOLUCIONA EL EMPAQUETAMIENTO TODO EL CODIGO QUE USAMOS, librerias de terceros etc
//Esto se implementa con un comando el cual es:
//go mod vendor
//aqui creara una nueva carpeta en donde deberia estar todo el codigo externo que necesitemos usar
//la buena practia es borrar los paquetes que no vamos a usar
//las librerias estan en el archivo go.mod y el go.sum es la verificacion de todo lo que usamos, para eliminar libreria es
//go mod tidy, esto quita las librerias que el programa no va a usar
import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	// Instanciar echo
	e := echo.New()
	// Ruta
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
