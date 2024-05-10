package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type ControleCampusEsp struct {
	db *sql.DB
}

//esta função cria um repositorio de campus
func NewRepositoryControleCampus(db *sql.DB) *ControleCampusEsp {
	return &ControleCampusEsp{db}
}

func (repositorio ControleCampusEsp) BuscarControle(IDCampus uint64) ([]modelos.ControleCampusEsp, error) {
	linhas, erro := repositorio.db.Query(
		"select c.id_Campus, c.id_Especialidade from controlecampusesp c inner join campus p on p.idCampus = c.id_Campus inner join especialidade e on e.idEspecialidade = c.id_Especialidade where c.id_Campus = ? order by c.id_Campus",
		IDCampus,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var (
		controleCampusEsp []modelos.ControleCampusEsp
		mapCampus         = map[uint64]modelos.ControleCampusEsp{}
	)

	for linhas.Next() {
		var (
			idCampus        uint64
			idEspecialidade uint64
		)

		if erro = linhas.Scan(
			&idCampus,
			&idEspecialidade,
		); erro != nil {
			return nil, erro
		}

		dbCampus := NewRepositoryCampus(repositorio.db)
		campus, erro := dbCampus.BuscarCampusPorID(idCampus)
		if erro != nil {
			return nil, erro
		}

		dbEspecialidade := NewRepositoryEspecialidade(repositorio.db)
		especialidade, erro := dbEspecialidade.BuscarEspecialidade(idEspecialidade)
		if erro != nil {
			return nil, erro
		}

		if _, ok := mapCampus[idCampus]; ok {
			data := mapCampus[idCampus]
			data.EspecialidadesCampus = append(data.EspecialidadesCampus, especialidade)
			mapCampus[idCampus] = data
		} else {
			mapCampus[idCampus] = modelos.ControleCampusEsp{
				Campus:               campus,
				EspecialidadesCampus: []modelos.Especialidade{especialidade},
			}
		}
	}

	for _, campus := range mapCampus {
		controleCampusEsp = append(controleCampusEsp, campus)
	}

	return controleCampusEsp, erro

}
