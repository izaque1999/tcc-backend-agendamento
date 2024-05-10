package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasCampus = []Rota{
	{
		URI:                "/campus/{idCampus}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarCampus,
		RequerAutenticacao: false,
	},
	{
		URI:                "/campus",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarCampusAll,
		RequerAutenticacao: false,
	},
}
