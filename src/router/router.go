package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

// Gerar vao retornar um router com as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
