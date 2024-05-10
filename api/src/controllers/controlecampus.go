package controllers

import (
	"api/src/banco"
	"api/src/repositorios"
	"api/src/respostas"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func BuscarControle(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	campusID, erro := strconv.ParseUint(parametros["id_Campus"], 10, 64)
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

	repositorio := repositorios.NewRepositoryControleCampus(db)
	controle, erro := repositorio.BuscarControle(campusID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, controle)

}
