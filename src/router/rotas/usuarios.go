package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:                "/createuser",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CreateUser,
		RequerAutenticacao: false,
	},
	{
		URI:                "/findusers",
		Metodo:             http.MethodGet,
		Funcao:             controllers.FindUsers,
		RequerAutenticacao: true,
	},
	{
		URI:                "/findubyid/{userid}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.FindUserByID,
		RequerAutenticacao: true,
	},
	{
		URI:                "/update/{userid}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.UpdateUser,
		RequerAutenticacao: true,
	},
	{
		URI:                "/deleteuser/{userid}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeleteUser,
		RequerAutenticacao: true,
	},

	{
		URI:                "/users/{userid}/follow",
		Metodo:             http.MethodPost,
		Funcao:             controllers.Follow,
		RequerAutenticacao: true,
	},
	{
		URI:                "/users/{userid}/unfollow",
		Metodo:             http.MethodPost,
		Funcao:             controllers.Unfollow,
		RequerAutenticacao: true,
	},
	{
		URI:                "/users/{userid}/findFollowers",
		Metodo:             http.MethodPost,
		Funcao:             controllers.FindFollowers,
		RequerAutenticacao: true,
	},
}
