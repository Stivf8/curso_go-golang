package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"stiven.com/go/rest-ws/handlers"
	"stiven.com/go/rest-ws/server"
)

func main() {
	//cargamos el archivo env con la libreria godotenv
	err := godotenv.Load(".env")
	//si hay error cargando el archivo lo imprimimos
	if err != nil {
		log.Fatalf("Error loading .env file %v\n", err)
	}
	//aqui definimos del archivo que variables va a tomar con os
	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	//instancias el nuevo server, primer parametro es el conexto donde se ejecuta, y el segundo es la configuracion del server
	s, err := server.NewServer(context.Background(), &server.Config{
		Port:        PORT,
		JWTSecret:   JWT_SECRET,
		DatabaseURL: DATABASE_URL,
	})
	//si hay error li imprimimos
	if err != nil {
		log.Fatalf("Error creating server %v\n", err)
	}
	//iniciamos el servidor con el metodo Start y enviamos la funcion BindRoutes
	s.Start(BindRoutes)
}

//tiene la misma estructura que el Start en el paquete Server
func BindRoutes(s server.Server, r *mux.Router) {
	//le indicamos la ruta al handler y que tipo es
	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)
}
