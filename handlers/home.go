package handlers

import (
	"encoding/json"
	"net/http"

	"stiven.com/go/rest-ws/server"
)

//este handler procesa la peticion de la pagina principal
//home respones
type HomeResponse struct {
	//enviamos al cliente los mensajes de mensaje y un estado
	//se deben enviar respuestas en tipo json entonces se realiza de esta manera
	//en Go la variable es Message pero cuando se serializa a json toma el nombre de message
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

//definir primer handler
//recibe como parametro el servidor, tipo servidor
//devuelve un HandlerFunc como respuesta
func HomeHandler(s server.Server) http.HandlerFunc {
	//va a retornar una funcion que va a ser un writter, que esta encargado de responder al cliente
	//y tambien tiene un request que es lo que envia el cliente la data
	return func(w http.ResponseWriter, r *http.Request) {
		//escribimos un header a esa peticion http
		//con esto le decimos al cliente que lo que se va a responder es un json asi lo interpreta y parcea
		w.Header().Set("Content-Type", "application/json")
		//escribimos otro header para enviar el estado
		w.WriteHeader(http.StatusOK)
		//CREAMOS LA RESPUESTA QUE LE VAMOS A ENVIAR
		//recibe como parametro el escritor que esta encargado de la respuesta
		//despues lo encodeamos
		json.NewEncoder(w).Encode(HomeResponse{
			Message: "Bienvenido a Stiven en Golang",
			Status:  true,
		})
	}
}
