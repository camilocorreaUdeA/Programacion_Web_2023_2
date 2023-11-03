package main

import (
	"encoding/json"
	"io"
	"net/http"
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
	err = json.Unmarshal(body, nuevoAmigo)
	if err != nil {
		http.Error(writer, "fallo al crear un nuevo amigo", http.StatusInternalServerError)
	}
	bd.Amigos[bd.proximoId] = *nuevoAmigo
	bd.proximoId++
	writer.WriteHeader(http.StatusCreated)
}

func main() {
	bd := BaseDatos{
		proximoId: uint64(1),
		Amigos:    make(map[uint64]Amigo),
	}
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

	http.ListenAndServe(":8080", nil)
}
