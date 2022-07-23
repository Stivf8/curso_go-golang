package database

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"stiven.com/go/rest-ws/models"
)

//implementa el repositorio de una manera concreta para la BD postgres

type PostgresRepository struct {
	db *sql.DB
}

//Definimos metodo para conectarnos a la bd
func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	return &PostgresRepository{db}, nil
}

//Funcion que cierra conexion
func (repo *PostgresRepository) Close() error {
	return repo.db.Close()
}

//como se van a insertar usuarios
func (repo *PostgresRepository) InsertUser(ctx context.Context, user *models.User) error {
	//funcion ExexContext permite realizar la consulta a la bd
	_, err := repo.db.ExecContext(ctx, "INSERT INTO users (email, password) VALUES ($1, $2)", user.Email, user.Password)
	return err
}

//consultar usuarios
func (repo *PostgresRepository) GetUserByID(ctx context.Context, id int64) (*models.User, error) {
	//QueryContext permite realizar la consulta y retornar los valores
	rows, err := repo.db.QueryContext(ctx, "SELECT id, email FROM users WHERE id = $1", id)
	if err != nil {
		//no podemos responder el users de manera directa ya que el dato no es el mismo que se define, entonces se debe trabajar el resultado
		//para que este pueda ser devuelto de manera correcta
		return nil, err
	}
	//ejecutamos funcion al final con defer para cerrar la conexion de los rows retornados
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	//crear la variable que va a ser devuelta
	var user = models.User{}
	//iteramos con for sobre las filas que retorna la bd y parceamos para ir registrando esos datos en user
	for rows.Next() {
		//el Scan toma la fila que se esta procesando y la intenta ingresar en la interface user
		//aqui asumimos que el error es nulo
		if err = rows.Scan(&user.Id, &user.Email); err == nil {
			return &user, nil
		}
	}
	//con la funcion Err nos dira cuando de verdad exista un error al ejecutar, retornando el error
	if err = rows.Err(); err != nil {
		return nil, err
	}
	//en caso de que todo OK retornamos los datos
	return &user, nil

}
