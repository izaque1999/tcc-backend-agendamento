package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type Campus struct {
	db *sql.DB
}

//esta função cria um repositorio de campus
func NewRepositoryCampus(db *sql.DB) *Campus {
	return &Campus{db}
}

//Buscar o campus

func (repositorio Campus) Buscar() ([]modelos.CampusUnip, error) {
	linhas, erro := repositorio.db.Query(
		"select * from campus",
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var campusUnip []modelos.CampusUnip

	for linhas.Next() {
		var campus modelos.CampusUnip
		if erro = linhas.Scan(
			&campus.ID,
			&campus.Nome,
			&campus.Descricao,
			&campus.Endereco,
			&campus.Numero,
			&campus.Complemento,
			&campus.CEP,
			&campus.Bairro,
			&campus.Cidade,
			&campus.Estado,
			&campus.Regiao,
			&campus.Telefone,
			&campus.Localizacao,
			&campus.DtInclusao,
		); erro != nil {
			return nil, erro
		}

		campusUnip = append(campusUnip, campus)
	}

	return campusUnip, erro

}

//Buscar o campus pelo ID
func (repositorio Campus) BuscarCampusPorID(ID uint64) (modelos.CampusUnip, error) {
	linhas, erro := repositorio.db.Query(
		"select idCampus, nome, descricao, endereco, numero, complemento, cep, bairro, cidade, estado, regiao, telefone, localizacao, dtinclusao from campus where idCampus = ?",
		ID,
	)
	if erro != nil {
		return modelos.CampusUnip{}, erro
	}
	defer linhas.Close()

	var campusUnip modelos.CampusUnip

	if linhas.Next() {
		if erro = linhas.Scan(
			&campusUnip.ID,
			&campusUnip.Nome,
			&campusUnip.Descricao,
			&campusUnip.Endereco,
			&campusUnip.Numero,
			&campusUnip.Complemento,
			&campusUnip.CEP,
			&campusUnip.Bairro,
			&campusUnip.Cidade,
			&campusUnip.Estado,
			&campusUnip.Regiao,
			&campusUnip.Telefone,
			&campusUnip.Localizacao,
			&campusUnip.DtInclusao,
		); erro != nil {
			return modelos.CampusUnip{}, erro
		}

	}

	return campusUnip, nil

}
