package controllers

import (
	"api/src/banco"
	"api/src/repositorios"
	"api/src/respostas"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Busca um campus no banco de dados pelo ID
func BuscarCampus(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	campusID, erro := strconv.ParseUint(parametros["idCampus"], 10, 64)
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

	repositorio := repositorios.NewRepositoryCampus(db)
	campus, erro := repositorio.BuscarCampusPorID(campusID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, campus)

}

//Busca um campus no banco de dados pelo ID
func BuscarCampusAll(w http.ResponseWriter, r *http.Request) {

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NewRepositoryCampus(db)
	campus, erro := repositorio.Buscar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, campus)

}
