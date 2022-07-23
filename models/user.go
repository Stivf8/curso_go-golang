package models

//como van a funcionar los modelos, su tipo y como se integran
//definimos su struct

type User struct {
	//json suele utilizar minusculas
	//go definimos en mayuscula para que sea publico
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

//patron de diseno Repository
//nueva carpeta con user.go y este a su vez referencia al paquete database archivo postgres.go
