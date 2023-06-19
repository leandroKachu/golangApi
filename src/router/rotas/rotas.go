package rotas

import (
	"api/src/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	rotas = append(rotas, rotasLogin)
	rotas = append(rotas, routesPost...)

	for _, rota := range rotas {

		if rota.RequerAutenticacao {
			// if true, add middleware validattion first logger to show in terminal and inside the autentication
			r.HandleFunc(rota.URI, middleware.Logger(middleware.Authenticator(rota.Funcao))).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.URI, middleware.Logger(rota.Funcao)).Methods(rota.Metodo)
		}
	}
	return r
}
