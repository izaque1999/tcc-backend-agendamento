package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasConsulta = []Rota{
	{
		URI:    "/consultas",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarConsulta,

		RequerAutenticacao: false,
	},
	{
		URI:                "/consultas",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarConsultaAll,
		RequerAutenticacao: false,
	},
	{
		URI:                "/consultas/{idConsulta}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarConsulta,
		RequerAutenticacao: false,
	},
	{
		URI:                "/consultas/{idConsulta}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/consultas/esp/{id_Especialidade}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarConsultaEspecialidade,
		RequerAutenticacao: false,
	},

	{
		URI:                "/consultas/{idConsulta}/atualizar-prioridade",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarPrioridade,
		RequerAutenticacao: false,
	},
	{
		URI:                "/consultas/user/{id_Usuario}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarConsultaUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/consultas/sts/",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarStsConsulta,
		RequerAutenticacao: false,
	},
}
