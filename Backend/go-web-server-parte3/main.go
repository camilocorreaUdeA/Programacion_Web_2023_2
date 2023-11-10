package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	repositorio "github.com/camilocorreaUdeA/go-web-server/repository" /* importando el paquete de repositorio */
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

/*
es mejor conservar un estándar entre las etiquetas de json y db para no tener problemas al parsear
de json a db en el método ActualizarUnAmigo
*/
type Amigo struct {
	Id          int    `db:"id" json:"id"`
	Nombre      string `db:"nombre" json:"nombre"`
	Edad        uint   `db:"edad" json:"edad"`
	Descripcion string `db:"descripcion" json:"descripcion"`
	EstaCerca   bool   `db:"esta_cerca" json:"esta_cerca"`
	Direccion   string `db:"direccion" json:"direccion"`
}

/*
la estructura BaseDatos solo tiene un campo que es un objeto del tipo de la interfaz
Repository definida en el paquete repository. Cualquier instancia concreta de la
interfaz puede ejecutar las implementaciones de los métodos Create, Read, List,
Update y Delete.
La interfaz puede ser instanciada para cualquier estructura, ya que es un genérico
o template de Go. Para este ejemplo está instanciada para el tipo de la estructura
Amigo
*/
type BaseDatos struct {
	repo repositorio.Repository[Amigo]
}

/*
este método decodifica lo que viene en el objeto json del cuerpo de la solicitud
en un map[string]any para que se pueda pasar directamente al metodo Update del
repositorio. Ejemplo:
Cuerpo: {"edad":25, "esta_cerca":true} ->decodificando-> map[string]any{"edad":25, "esta_cerca":true}
*/
func (bd *BaseDatos) ActualizarUnAmigo(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("fallo al actualizar un amigo, con error: %s", err.Error())
		http.Error(writer, fmt.Sprintf("fallo al actualizar un amigo, con error: %s", err.Error()), http.StatusBadRequest)
		return
	}
	defer req.Body.Close()
	nuevosValoresAmigo := make(map[string]any)
	err = json.Unmarshal(body, &nuevosValoresAmigo)
	if err != nil {
		log.Printf("fallo al actualizar un amigo, con error: %s", err.Error())
		http.Error(writer, fmt.Sprintf("fallo al actualizar un amigo, con error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	if len(nuevosValoresAmigo) == 0 {
		log.Printf("fallo al actualizar un amigo, con error: %s", err.Error())
		http.Error(writer, fmt.Sprintf("fallo al actualizar un amigo, con error: %s", err.Error()), http.StatusBadRequest)
		return
	}

	query := "UPDATE amigos SET %s WHERE id=:id;"
	columns := []string{}
	/* el resultado de este ciclo for es: columns = []string{"col1=:col1", "col2=:col2",..., "coln=:coln"} */
	for key := range nuevosValoresAmigo {
		columns = append(columns, fmt.Sprintf("%s=:%s", key, key))
	}
	/* columnString = "col1=:col1, col2=:col2,..., coln=:coln" */
	columnsString := strings.Join(columns, ",")
	/* realQuery = "UPDATE amigos SET col1=:col1, col2=:col2,..., coln=:coln WHERE id=:id;" */
	realQuery := fmt.Sprintf(query, columnsString)
	/* nuevosValoresAmigo = map[string]any{
		"id":id,
		"col1":valor_col1,
		"col2":valor_col2,
		....,
		"coln":valor_coln,
	} */
	nuevosValoresAmigo["id"] = id
	/* el metodo Update utiliza un query con nombres, ejemplo: update amigos set edad=:edad where id=:id;
	los nombres sirven para que busque esas claves en el mapa y pase al query los valores correspondientes
	a las claves del mapa que coinciden con los nombres del query (:nombre)
	*/
	err = bd.repo.Update(context.TODO(), realQuery, nuevosValoresAmigo)
	if err != nil {
		log.Printf("fallo al actualizar un amigo, con error: %s", err.Error())
		http.Error(writer, fmt.Sprintf("fallo al actualizar un amigo, con error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (bd *BaseDatos) EliminarUnAmigo(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	/* aquí el valor del parámetro id se pasa como $1 porque es un parámetro que se pasa directamente al método Delete del repositorio
	y no como un valor nombrado en un map */
	query := "DELETE FROM amigos WHERE id=$1;"
	err := bd.repo.Delete(context.TODO(), query, id)
	if err != nil {
		log.Printf("fallo al eliminar un amigo, con error: %s", err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(fmt.Sprintf("fallo al eliminar un amigo con id %s", id)))
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (bd *BaseDatos) LeerUnAmigo(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	/* aquí el valor del parámetro id se pasa como $1 porque es un parámetro que se pasa directamente al método Read del repositorio
	y no como un valor nombrado en un map */
	query := "SELECT id, nombre, edad, descripcion, esta_cerca, direccion FROM amigos WHERE id=$1;"
	amigo, err := bd.repo.Read(context.TODO(), query, id)
	if err != nil {
		log.Printf("fallo al leer un amigo, con error: %s", err.Error())
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte(fmt.Sprintf("el amigo con id %s no se pudo encontrar", id)))
		return
	}
	amigoJson, err := json.Marshal(amigo)
	if err != nil {
		log.Printf("fallo al leer un amigo, con error: %s", err.Error())
		http.Error(writer, "fallo al codificar los datos", http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(amigoJson)
}

/*
para pasar valores a los parámetros limit y offset desde la solicitud a la API
va a tener que usar query params (parámetros opcionales en la URL). Ejemplo:
.../amigos?limit=xx&offset=zz (donde xx y zz son los valores para limit y offset)
Luego los debe extraer de la solicitud con la función request.URL.Query
esa función devuelve un map[string]any donde las claves son los nombres que pasó
en la URL y los valores lo que puso en xx y zz. Ejemplo:
/amigos?limit=xx&offset=zz -> map[string]any{"limit":xx, "offset":zz}
*/
func (bd *BaseDatos) LeerAmigos(writer http.ResponseWriter, req *http.Request) {
	/* aquí el valor del parámetro limit se pasa como $1 porque es un parámetro que se pasa directamente al método List del repositorio
	y no como un valor nombrado en un map
	aquí el valor del parámetro offset se pasa como $2 porque es un parámetro que se pasa directamente al método List del repositorio
	y no como un valor nombrado en un map y va segundo en orden después del parámetro limit */
	query := "SELECT id, nombre, edad, descripcion, esta_cerca, direccion FROM amigos limit $1 offset $2"
	amigos, _, err := bd.repo.List(context.TODO(), query, 100, 0)
	if err != nil {
		log.Printf("fallo al leer amigos, con error: %s", err.Error())
		http.Error(writer, "fallo al leer los amigos", http.StatusInternalServerError)
		return
	}
	jsonAmigos, err := json.Marshal(amigos)
	if err != nil {
		log.Printf("fallo al leer amigos, con error: %s", err.Error())
		http.Error(writer, "fallo al leer los amigos", http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(jsonAmigos)
}

func (bd *BaseDatos) CrearAmigo(writer http.ResponseWriter, req *http.Request) {
	nuevoAmigo := &Amigo{}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("fallo al crear un nuevo amigo, con error: %s", err.Error())
		http.Error(writer, "fallo al crear un nuevo amigo", http.StatusBadRequest)
		return
	}
	defer req.Body.Close()
	err = json.Unmarshal(body, nuevoAmigo)
	if err != nil {
		log.Printf("fallo al crear un nuevo amigo, con error: %s", err.Error())
		http.Error(writer, "fallo al crear un nuevo amigo", http.StatusInternalServerError)
		return
	}

	/* se construye un map para pasarle al método Create del repositorio con los valores recibidos
	en el objeto json del cuerpo de la solicitud POST (decodificado en un objeto del tipo Amigo, claramente) */
	valoresColumnasNuevoAmigo := map[string]any{
		"nombre":      nuevoAmigo.Nombre,
		"edad":        nuevoAmigo.Edad,
		"descripcion": nuevoAmigo.Descripcion,
		"esta_cerca":  nuevoAmigo.EstaCerca,
		"direccion":   nuevoAmigo.Direccion,
	}
	/* se agrega el comando RETURNING al query INSERT para que devuelva el id de la fila recien creada */
	query := "INSERT INTO amigos (nombre, edad, descripcion, esta_cerca, direccion) VALUES (:nombre, :edad, :descripcion, :esta_cerca, :direccion) returning id;"
	nuevoId, err := bd.repo.Create(context.TODO(), query, valoresColumnasNuevoAmigo)
	if err != nil {
		log.Println("fallo al crear un nuevo amigo, con error:", err.Error())
		http.Error(writer, "fallo al crear un nuevo amigo", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte(fmt.Sprintf("id nuevo amigo: %d", nuevoId)))
}

/*
función para conectarse a la instancia de PostgreSQL, en general sirve para cualquier base de datos SQL.
Necesita la URL del host donde está instalada la base de datos y el tipo de base datos (driver)
*/
func ConectarDB(url, driver string) (*sqlx.DB, error) {
	pgUrl, _ := pq.ParseURL(url)
	db, err := sqlx.Connect(driver, pgUrl) // driver: postgres
	if err != nil {
		log.Printf("fallo la conexion a PostgreSQL, error: %s", err.Error())
		return nil, err
	}

	log.Printf("Nos conectamos bien a la base de datos db: %#v", db)
	return db, nil
}

func main() {
	/* creando un objeto de conexión a PostgreSQL */
	db, err := ConectarDB("aquí_va_la_URL_de_conexión_de_tu_instancia_de_PostgreSQL", "postgres")
	if err != nil {
		log.Fatalln("error conectando a la base de datos", err.Error())
		return
	}

	/* creando una instancia del tipo Repository del paquete repository
	se debe especificar el tipo de struct que va a manejar la base de datos
	para este ejemplo es Amigo y se le pasa como parámetro el objeto de
	conexión a PostgreSQL */
	repo, err := repositorio.NewRepository[Amigo](db)
	if err != nil {
		log.Fatalln("fallo al crear una instancia de repositorio", err.Error())
		return
	}

	/* incializamos un objeto del tipo BaseDatos, con su campo repo igual al objeto del tipo
	de la interfaz Repository del paquete repository */
	bd := BaseDatos{
		repo: repo,
	}

	/* router (multiplexador) a los endpoints de la API (implementado con el paquete gorilla/mux) */
	router := mux.NewRouter()

	/* rutas a los endpoints de la API */
	router.Handle("/amigos", http.HandlerFunc(bd.LeerAmigos)).Methods(http.MethodGet)
	router.Handle("/amigos", http.HandlerFunc(bd.CrearAmigo)).Methods(http.MethodPost)
	router.Handle("/amigos/{id}", http.HandlerFunc(bd.LeerUnAmigo)).Methods(http.MethodGet)
	router.Handle("/amigos/{id}", http.HandlerFunc(bd.ActualizarUnAmigo)).Methods(http.MethodPatch)
	router.Handle("/amigos/{id}", http.HandlerFunc(bd.EliminarUnAmigo)).Methods(http.MethodDelete)

	/* servidor escuchando en localhost por el puerto 8080 y entrutando las peticiones con el router */
	http.ListenAndServe(":8080", router)
}
