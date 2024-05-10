package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasControle = []Rota{
	{
		URI:                "/controle/{id_Campus}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarControle,
		RequerAutenticacao: false,
	},
}
