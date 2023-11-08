package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type Amigo struct {
	Nombre      string `json:"nombre"`
	Edad        uint   `json:"edad"`
	Descripcion string `json:"descripcion"`
	EstaCerca   bool   `json:"estaCerca"`
	Direccion   string `json:"direccion"`
}

type BaseDatos struct {
	Amigos    map[uint64]Amigo
	proximoId uint64
}

func (bd *BaseDatos) ActualizarUnAmigo(writer http.ResponseWriter, req *http.Request) {
}

func (bd *BaseDatos) EliminarUnAmigo(writer http.ResponseWriter, req *http.Request) {
}

// amigos?fechaInicio=xxx&fechaFin=yyy
func (bd *BaseDatos) LeerUnAmigo(writer http.ResponseWriter, req *http.Request) {
	/* idString := strings.TrimPrefix(req.URL.Path, "/amigos/") // /amigos/xx -> xx */
	/* queryParams := req.URL.Query()
	inicio := queryParams.Get("fechaInicio") */
	vars := mux.Vars(req)
	idString := vars["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(writer, "fallo al convertir el id", http.StatusInternalServerError)
		return
	}
	amigo, found := bd.Amigos[uint64(id)]
	if !found {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte(fmt.Sprintf("el amigo con id %d no se pudo encontrar", id)))
		return
	}
	amigoJson, err := json.Marshal(amigo)
	if err != nil {
		http.Error(writer, "fallo al codificar los datos", http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(amigoJson)
}

func (bd *BaseDatos) LeerAmigos(writer http.ResponseWriter, req *http.Request) {
	amigos := []Amigo{}
	for _, amigo := range bd.Amigos {
		amigos = append(amigos, amigo)
	}
	jsonAmigos, err := json.Marshal(amigos)
	if err != nil {
		http.Error(writer, "fallo al leer los amigos", http.StatusInternalServerError)
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(jsonAmigos)
}

func (bd *BaseDatos) CrearAmigo(writer http.ResponseWriter, req *http.Request) {
	nuevoAmigo := &Amigo{}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(writer, "fallo al crear un nuevo amigo", http.StatusBadRequest)
	}
	defer req.Body.Close()
	err = json.Unmarshal(body, nuevoAmigo)
	if err != nil {
		http.Error(writer, "fallo al crear un nuevo amigo", http.StatusInternalServerError)
	}
	bd.Amigos[bd.proximoId] = *nuevoAmigo
	bd.proximoId++
	writer.WriteHeader(http.StatusCreated)
}

func LeerUnAmigoHandler(bd *BaseDatos) http.Handler {
	handler := func(writer http.ResponseWriter, req *http.Request) {
		idString := strings.TrimPrefix(req.URL.Path, "/amigos/") // /amigos/xx -> xx
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(writer, "fallo al convertir el id", http.StatusInternalServerError)
			return
		}
		amigo, found := bd.Amigos[uint64(id)]
		if !found {
			writer.WriteHeader(http.StatusNotFound)
			writer.Write([]byte(fmt.Sprintf("el amigo con id %d no se pudo encontrar", id)))
			return
		}
		amigoJson, err := json.Marshal(amigo)
		if err != nil {
			http.Error(writer, "fallo al codificar los datos", http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(amigoJson)
	}
	return http.HandlerFunc(handler)
}

func main() {
	bd := BaseDatos{
		proximoId: uint64(1),
		Amigos:    make(map[uint64]Amigo),
	}

	router := mux.NewRouter()

	router.Handle("/amigos", http.HandlerFunc(bd.LeerAmigos)).Methods(http.MethodGet)
	router.Handle("/amigos", http.HandlerFunc(bd.CrearAmigo)).Methods(http.MethodPost)
	router.Handle("/amigos/{id}", http.HandlerFunc(bd.LeerUnAmigo)).Methods(http.MethodGet)

	http.ListenAndServe(":8080", router)

	// Router o mux para REST API -> gorilla/mux

	// localhost:8080/amigos/1 -> path parameter

	/*http.HandleFunc("/amigos/", func(writer http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			bd.LeerUnAmigo(writer, req)
		default:
			http.Error(writer, "Ese metodo no esta permitido", http.StatusMethodNotAllowed)
		}
	})*/

	/*http.Handle("/amigos/", LeerUnAmigoHandler(&bd))

	http.HandleFunc("/amigos", func(writer http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			bd.LeerAmigos(writer, req)
		case http.MethodPost:
			bd.CrearAmigo(writer, req)
		default:
			http.Error(writer, "Ese metodo no esta permitido", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil) */
}
