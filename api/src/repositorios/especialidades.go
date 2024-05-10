package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type Especialidades struct {
	db *sql.DB
}

// esta função cria um repositorio de usuarios
func NewRepositoryEspecialidade(db *sql.DB) *Especialidades {
	return &Especialidades{db}
}

// Buscar o usuario pelo ID
func (repositorio Especialidades) BuscarEspecialidade(ID uint64) (modelos.Especialidade, error) {
	linhas, erro := repositorio.db.Query(
		"select idEspecialidade, nome from especialidade where idEspecialidade = ?",
		ID,
	)
	if erro != nil {
		return modelos.Especialidade{}, erro
	}
	defer linhas.Close()

	var especialidades modelos.Especialidade

	if linhas.Next() {
		if erro = linhas.Scan(
			&especialidades.ID,
			&especialidades.Nome,
		); erro != nil {
			return modelos.Especialidade{}, erro
		}

	}

	return especialidades, nil

}

func (repositorio Especialidades) BuscarEspAll() ([]modelos.Especialidade, error) {
	linhas, erro := repositorio.db.Query(
		"select idEspecialidade, nome from especialidade",
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var esp []modelos.Especialidade

	for linhas.Next() {
		var especialidades modelos.Especialidade
		if erro = linhas.Scan(
			&especialidades.ID,
			&especialidades.Nome,
		); erro != nil {
			return nil, erro
		}
		esp = append(esp, especialidades)

	}

	return esp, erro

}
