package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//configuracion que el servidor requiere para poder ejecutarse
type Config struct {
	//puerto donde se ejecuta
	Port string
	//clave secreta para generar los tokens
	JWTSecret string
	//url de conexion a la bd
	DatabaseURL string
}

//definir interfaz para el tipo server - es una manera de definir metodos del server
//para que "algo" sea tomado en cuenta como servidor, debe de tener Config osea que retorne dicha configuracion
type Server interface {
	//el metodo Config retornara algun dato(*) de tipo Config
	Config() *Config
}

//definir broker que va a ser encargado de manejar esos servidores

type Broker struct {
	//este va a tener la configuracion con las propiedades definidas anteriormente
	config *Config
	//va a tener un ruteador que define las rutas que la API tendra
	//utilizamos la dependencia mux
	router *mux.Router
}

//para que el Broker satisfaga la interface Server es de la siguiente manera
//debe tener el metodo Config
//receiver function
func (b *Broker) Config() *Config {
	//necesitamos devolver la configuracion del Broker, para que se comporte como un tipo Server
	return b.config
}

//definimos el constructor para nuestro Struct
//recibe contexto para encontrar posibles problemas de nuestro codigo
//y el segundo parametro es la configuracion, utilizamos el apuntador para el dato
//luego retornamos dos valores, el primero tipo Broker y el segundo el error
func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	//REVISAMOS LA CONFIGURACION PARA ASEGURAR QUE NO SON CAMPOS VACIOS
	if config.Port == "" || config.JWTSecret == "" || config.DatabaseURL == "" {
		//retornamos error
		//retornamos primero nil porque en este caso no tendriamos el Broker definido
		return nil, errors.New("Algo fallo en la configuracion del Servidor, revisa los parametros de puerto, clave privada y conexion a base de datos")

	}

	//retornamos el broker instanciandolo
	//retornamos la configuracion
	//y un rotuer nuevo, que lo que hara es definir una nueva instancia, para eso usamos la dependencia mux
	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
	}
	//retornamos el broker
	return broker, nil
}

//agregamos metodo al Broker para que se pueda ejecutar
//recibe funcion binder como parametro
//la funcion recibe una variable tipo Server
//la funcion recibe luego una variable de routeador
func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	//traer un nuevo router con la libreria mux, definir objeto
	b.router = mux.NewRouter()
	//b es tipo server y a router tipo rotuer
	binder(b, b.router)
	//imprimimos mensaje para verificar
	log.Println("Inicializando Servidor en puerto:", b.Config().Port)
	//ejecutamos el servidor
	//revisamos si hay errores
	//usamos http con el metodo ListenAndServe que recibe el puerto y el router donde estaran mis rutas las que yo defina
	//despues revisamos que nuestro error sea diferente a nil, si es nil es todo ok y si no envia log
	if err := http.ListenAndServe(b.config.Port, b.router); err != nil {
		//imprime cadena de texto
		log.Fatal("ListenAndServe: ", err)
	}

}
