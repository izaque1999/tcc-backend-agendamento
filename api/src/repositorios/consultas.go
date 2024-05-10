package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

type Consultas struct {
	db *sql.DB
}

// esta função cria um repositorio de usuarios
func NewRepositoryConsulta(db *sql.DB) *Consultas {
	return &Consultas{db}
}

// Insere uma consulta no banco de dados
func (repositorio Consultas) Criar(consulta modelos.Consulta) (uint64, error) {

	var counter int

	erro := repositorio.db.QueryRow(
		"select count(idConsulta) from consulta where dtMarcacao = ? and hrMarcacao = ? and id_Campus =  ? and id_Especialidade = ?",
		consulta.DtcMarcacao, consulta.HrMarcacao, consulta.IDCampus, consulta.IDEspecialidade,
	).Scan(&counter)
	if erro != nil {
		return 0, erro
	}
	fmt.Println(erro)
	if counter > 0 {
		return 0, fmt.Errorf("Nao foi possivel criar a consulta, data indisponivel")
	}
	fmt.Println("Passou aqui")
	statement, erro := repositorio.db.Prepare(
		"insert into consulta ( descricao, dtMarcacao, hrMarcacao, id_Campus, id_Usuario, id_Especialidade ) values(?, ?, ?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(consulta.Observacao, consulta.DtcMarcacao, consulta.HrMarcacao, consulta.IDCampus, consulta.IDUsuario, consulta.IDEspecialidade)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil

}

func (repositorio Consultas) Buscar() ([]modelos.Consulta, error) {
	linhas, erro := repositorio.db.Query(
		"select idConsulta, descricao, stsConsulta,dtMarcacao, hrMarcacao,id_campus,id_Usuario,id_Especialidade from consulta order by dtMarcacao, HrMarcacao, stsConsulta desc",
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var consultaT []modelos.Consulta

	for linhas.Next() {
		var consulta modelos.Consulta
		if erro = linhas.Scan(
			&consulta.ID,
			&consulta.Observacao,
			&consulta.StatusConsulta,
			&consulta.DtcMarcacao,
			&consulta.HrMarcacao,
			&consulta.IDCampus,
			&consulta.IDUsuario,
			&consulta.IDEspecialidade,
		); erro != nil {
			return nil, erro
		}
		dbCampus := NewRepositoryCampus(repositorio.db)
		campus, erro := dbCampus.BuscarCampusPorID(consulta.IDCampus)
		if erro != nil {
			return nil, erro
		}

		consulta.Campus = campus

		dbUsuario := NewRepositoryUsers(repositorio.db)
		usuario, erro := dbUsuario.BuscarPorID(consulta.IDUsuario)
		if erro != nil {
			return nil, erro
		}

		consulta.NomeUsuario = usuario

		consulta.Campus = campus
		dbEspecialidade := NewRepositoryEspecialidade(repositorio.db)
		especialidade, erro := dbEspecialidade.BuscarEspecialidade(consulta.IDEspecialidade)
		if erro != nil {
			return nil, erro
		}

		consulta.NomeEspecialidade = especialidade
		consultaT = append(consultaT, consulta)
	}

	return consultaT, erro

}

// Buscar a consulta pelo ID
func (repositorio Consultas) BuscarPorID(ID uint64) (modelos.Consulta, error) {
	linhas, erro := repositorio.db.Query(
		"select idConsulta, descricao, stsConsulta, dtMarcacao, hrMarcacao, id_Campus, id_Usuario, id_Especialidade from consulta where idConsulta = ? order by dtMarcacao, HrMarcacao, stsConsulta",
		ID,
	)
	if erro != nil {
		return modelos.Consulta{}, erro
	}
	defer linhas.Close()

	var consulta modelos.Consulta

	if linhas.Next() {
		if erro = linhas.Scan(
			&consulta.ID,
			&consulta.Observacao,
			&consulta.StatusConsulta,
			&consulta.DtcMarcacao,
			&consulta.HrMarcacao,
			&consulta.IDCampus,
			&consulta.IDUsuario,
			&consulta.IDEspecialidade,
		); erro != nil {
			return modelos.Consulta{}, erro
		}
		dbCampus := NewRepositoryCampus(repositorio.db)
		campus, erro := dbCampus.BuscarCampusPorID(consulta.IDCampus)
		if erro != nil {
			return modelos.Consulta{}, erro
		}

		consulta.Campus = campus

		dbUsuario := NewRepositoryUsers(repositorio.db)
		usuario, erro := dbUsuario.BuscarPorID(consulta.IDUsuario)
		if erro != nil {
			return modelos.Consulta{}, erro
		}

		consulta.NomeUsuario = usuario

		consulta.Campus = campus
		dbEspecialidade := NewRepositoryEspecialidade(repositorio.db)
		especialidade, erro := dbEspecialidade.BuscarEspecialidade(consulta.IDEspecialidade)
		if erro != nil {
			return modelos.Consulta{}, erro
		}

		consulta.NomeEspecialidade = especialidade

	}

	return consulta, nil

}

// Buscar a consulta pelo ID Especialidade
func (repositorio Consultas) BuscarEspecialidade(ID uint64) ([]modelos.Consulta, error) {
	linhas, erro := repositorio.db.Query(
		"select idConsulta, descricao, stsConsulta, prioridade, dtMarcacao, hrMarcacao, id_Campus, id_Usuario, id_Especialidade from consulta where id_Especialidade = ?",
		ID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var consultaE []modelos.Consulta

	for linhas.Next() {
		var consulta modelos.Consulta
		if erro = linhas.Scan(
			&consulta.ID,
			&consulta.Observacao,
			&consulta.StatusConsulta,
			&consulta.Prioridade,
			&consulta.DtcMarcacao,
			&consulta.HrMarcacao,
			&consulta.IDCampus,
			&consulta.IDUsuario,
			&consulta.IDEspecialidade,
		); erro != nil {
			return nil, erro
		}
		dbCampus := NewRepositoryCampus(repositorio.db)
		campus, erro := dbCampus.BuscarCampusPorID(consulta.IDCampus)
		if erro != nil {
			return nil, erro
		}

		consulta.Campus = campus

		dbUsuario := NewRepositoryUsers(repositorio.db)
		usuario, erro := dbUsuario.BuscarPorID(consulta.IDUsuario)
		if erro != nil {
			return nil, erro
		}

		consulta.NomeUsuario = usuario

		consulta.Campus = campus
		dbEspecialidade := NewRepositoryEspecialidade(repositorio.db)
		especialidade, erro := dbEspecialidade.BuscarEspecialidade(consulta.IDEspecialidade)
		if erro != nil {
			return nil, erro
		}

		consulta.NomeEspecialidade = especialidade
		consultaE = append(consultaE, consulta)
	}

	return consultaE, erro

}

func (repositorio Consultas) BuscarUsuario(ID uint64) ([]modelos.Consulta, error) {
	linhas, erro := repositorio.db.Query(
		"select idConsulta, descricao, stsConsulta, dtMarcacao, hrMarcacao, id_Campus, id_Usuario, id_Especialidade from consulta where id_Usuario = ?",
		ID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var consultaT []modelos.Consulta

	for linhas.Next() {
		var consulta modelos.Consulta
		if erro = linhas.Scan(
			&consulta.ID,
			&consulta.Observacao,
			&consulta.StatusConsulta,
			&consulta.DtcMarcacao,
			&consulta.HrMarcacao,
			&consulta.IDCampus,
			&consulta.IDUsuario,
			&consulta.IDEspecialidade,
		); erro != nil {
			return nil, erro
		}
		dbCampus := NewRepositoryCampus(repositorio.db)
		campus, erro := dbCampus.BuscarCampusPorID(consulta.IDCampus)
		if erro != nil {
			return nil, erro
		}

		consulta.Campus = campus

		dbUsuario := NewRepositoryUsers(repositorio.db)
		usuario, erro := dbUsuario.BuscarPorID(consulta.IDUsuario)
		if erro != nil {
			return nil, erro
		}

		consulta.NomeUsuario = usuario

		consulta.Campus = campus
		dbEspecialidade := NewRepositoryEspecialidade(repositorio.db)
		especialidade, erro := dbEspecialidade.BuscarEspecialidade(consulta.IDEspecialidade)
		if erro != nil {
			return nil, erro
		}

		consulta.NomeEspecialidade = especialidade
		consultaT = append(consultaT, consulta)
	}

	return consultaT, erro

}

func (repositorio Consultas) AtualizarPrioridade(ID uint64, consulta modelos.Consulta) error {
	statement, erro := repositorio.db.Prepare(
		"update consulta set stsConsulta = ? , prioridade = ? where idConsulta= ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(consulta.StatusConsulta, consulta.Prioridade, ID); erro != nil {
		return erro
	}
	return nil
}

func (repositorio Consultas) BuscarStatus(status string) ([]modelos.Consulta, error) {
	status = fmt.Sprintf("%s", status)
	linhas, erro := repositorio.db.Query(
		"select idConsulta, descricao, stsConsulta, dtMarcacao, hrMarcacao, id_Campus, id_Usuario, id_Especialidade from consulta where stsConsulta = ? order by stsConsulta",
		status,
	)
	if erro != nil {
		fmt.Println("Teste 1")
		return nil, erro
	}
	defer linhas.Close()

	var statusCon []modelos.Consulta

	for linhas.Next() {
		var statusCons modelos.Consulta
		if erro = linhas.Scan(
			&statusCons.ID,
			&statusCons.Observacao,
			&statusCons.StatusConsulta,
			&statusCons.DtcMarcacao,
			&statusCons.HrMarcacao,
			&statusCons.IDCampus,
			&statusCons.IDUsuario,
			&statusCons.IDEspecialidade,
		); erro != nil {
			return nil, erro
		}
		dbCampus := NewRepositoryCampus(repositorio.db)
		campus, erro := dbCampus.BuscarCampusPorID(statusCons.IDCampus)
		if erro != nil {
			return nil, erro
		}

		statusCons.Campus = campus

		dbUsuario := NewRepositoryUsers(repositorio.db)
		usuario, erro := dbUsuario.BuscarPorID(statusCons.IDUsuario)
		if erro != nil {
			return nil, erro
		}

		statusCons.NomeUsuario = usuario

		statusCons.Campus = campus
		dbEspecialidade := NewRepositoryEspecialidade(repositorio.db)
		especialidade, erro := dbEspecialidade.BuscarEspecialidade(statusCons.IDEspecialidade)
		if erro != nil {
			return nil, erro
		}

		statusCons.NomeEspecialidade = especialidade

		statusCon = append(statusCon, statusCons)

	}

	return statusCon, nil

}
