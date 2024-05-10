package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Criar usuario no banco de dados
func CriarConsulta(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var consulta modelos.Consulta
	if erro = json.Unmarshal(bodyRequest, &consulta); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = consulta.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repositorio := repositorios.NewRepositoryConsulta(db)
	consulta.ID, erro = repositorio.Criar(consulta)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, consulta)
}

// Busca um usuario no banco de dados pelo ID
func BuscarConsulta(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	consultaID, erro := strconv.ParseUint(parametros["idConsulta"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NewRepositoryConsulta(db)
	consulta, erro := repositorio.BuscarPorID(consultaID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, consulta)

}

// Busca um usuario no banco de dados pelo ID
func BuscarConsultaEspecialidade(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	especialidadeID, erro := strconv.ParseUint(parametros["id_Especialidade"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NewRepositoryConsulta(db)
	especialidade, erro := repositorio.BuscarEspecialidade(especialidadeID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, especialidade)

}

func BuscarConsultaUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	especialidadeID, erro := strconv.ParseUint(parametros["id_Usuario"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NewRepositoryConsulta(db)
	especialidade, erro := repositorio.BuscarUsuario(especialidadeID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, especialidade)

}

// Atualiza Prioridade e Status da Consulta
func AtualizarPrioridade(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	consultaID, erro := strconv.ParseUint(parametros["idConsulta"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var consulta modelos.Consulta
	if erro = json.Unmarshal(bodyRequest, &consulta); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NewRepositoryConsulta(db)
	if erro = repositorio.AtualizarPrioridade(consultaID, consulta); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// Busca todas as consultas no banco de dados
func BuscarConsultaAll(w http.ResponseWriter, r *http.Request) {

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NewRepositoryConsulta(db)
	consulta, erro := repositorio.Buscar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, consulta)

}

// Busca as consultas por Status no banco de dados
func BuscarStsConsulta(w http.ResponseWriter, r *http.Request) {

	status := (r.URL.Query().Get("stsConsulta"))

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NewRepositoryConsulta(db)
	consulta, erro := repositorio.BuscarStatus(status)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, consulta)

}
