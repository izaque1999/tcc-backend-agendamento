package controllers

import (
	"api/src/banco"
	"api/src/repositorios"
	"api/src/respostas"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Busca uma especialidade no banco de dados pelo ID
func BuscarEspecialidade(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	especialidadeID, erro := strconv.ParseUint(parametros["idEspecialidade"], 10, 64)
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

	repositorio := repositorios.NewRepositoryEspecialidade(db)
	campus, erro := repositorio.BuscarEspecialidade(especialidadeID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, campus)

}

func BuscarEspecialidadeAll(w http.ResponseWriter, r *http.Request) {

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NewRepositoryEspecialidade(db)
	especialidade, erro := repositorio.BuscarEspAll()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, especialidade)

}
