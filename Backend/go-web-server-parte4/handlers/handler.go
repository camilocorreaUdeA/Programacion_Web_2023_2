package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/camilocorreaUdeA/go-web-server/controllers"
	"github.com/gorilla/mux"
)

type Handler struct {
	controller *controllers.Controller
}

func NewHandler(controller *controllers.Controller) (*Handler, error) {
	if controller == nil {
		return nil, fmt.Errorf("para instanciar un handler se necesita un controlador no nulo")
	}
	return &Handler{
		controller: controller,
	}, nil
}

func (h *Handler) ActualizarUnAmigo(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("fallo al actualizar un amigo, con error: %s", err.Error())
		http.Error(writer, fmt.Sprintf("fallo al actualizar un amigo, con error: %s", err.Error()), http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	err = h.controller.ActualizarUnAmigo(body, id)
	if err != nil {
		log.Printf("fallo al actualizar un amigo, con error: %s", err.Error())
		http.Error(writer, fmt.Sprintf("fallo al actualizar un amigo, con error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (h *Handler) EliminarUnAmigo(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	err := h.controller.EliminarUnAmigo(id)
	if err != nil {
		log.Printf("fallo al eliminar un amigo, con error: %s", err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(fmt.Sprintf("fallo al eliminar un amigo con id %s", id)))
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (h *Handler) LeerUnAmigo(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	amigo, err := h.controller.LeerUnAmigo(id)
	if err != nil {
		log.Printf("fallo al leer un amigo, con error: %s", err.Error())
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte(fmt.Sprintf("el amigo con id %s no se pudo encontrar", id)))
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(amigo)
}

func (h *Handler) LeerAmigos(writer http.ResponseWriter, req *http.Request) {
	amigos, err := h.controller.LeerAmigos(100, 0)
	if err != nil {
		log.Printf("fallo al leer amigos, con error: %s", err.Error())
		http.Error(writer, "fallo al leer los amigos", http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(amigos)
}

func (h *Handler) CrearAmigo(writer http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("fallo al crear un nuevo amigo, con error: %s", err.Error())
		http.Error(writer, "fallo al crear un nuevo amigo", http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	nuevoId, err := h.controller.CrearAmigo(body)
	if err != nil {
		log.Println("fallo al crear un nuevo amigo, con error:", err.Error())
		http.Error(writer, "fallo al crear un nuevo amigo", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte(fmt.Sprintf("id nuevo amigo: %d", nuevoId)))
}
