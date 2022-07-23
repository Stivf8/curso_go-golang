package repository

import (
	"context"

	"stiven.com/go/rest-ws/models"
)

//definimos una nueva interfaz

type UserRepository interface {
	//definimos plantilla de metodos de la interface, que indica que deben ser implementados pero no los implementa
	//tendra el insert, recibe el contexto, el usuario del modelo y retorna error en caso de
	InsertUser(ctx context.Context, user *models.User) error
	//definimos otro metodo, que recibe contexto, el id del usuario y retorna el Usuario en dado caso el error
	GetUserByID(ctx context.Context, id int64) (*models.User, error)
	//cerrar conexionas a base de datos
	//devuelve error
	Close() error
}

//inyeccion de dependencia
//define en que se va a implementar para ser abstractos
var implementation UserRepository

//este le indica al metodo en que se va a implementar la consulta en este caso
func SetRepository(repository UserRepository) {
	implementation = repository
}

//definimos metodo insert user, recibe contexto y usuario, este devuelve lo que la implementacion hace
func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user)
}

//definimos metodo get user, recibe contexto y id, este devuelve lo que la implementacion hace
func GetUserByID(ctx context.Context, id int64) (*models.User, error) {
	return implementation.GetUserByID(ctx, id)
}

//funcion close, va a devolver lo que la implementacion esta haciendo
func Close() error {
	return implementation.Close()
}
