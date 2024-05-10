package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasEspecialidades = []Rota{
	{
		URI:                "/especialidade/{idEspecialidade}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarEspecialidade,
		RequerAutenticacao: false,
	},
	{
		URI:                "/especialidades",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarEspecialidadeAll,
		RequerAutenticacao: false,
	},
}
