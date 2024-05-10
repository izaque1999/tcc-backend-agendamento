package modelos

import (
	"errors"
	"strings"
)

type Consulta struct {
	ID                uint64 `json:"idConsulta,omitempty"`
	Observacao        string `json:"descricao,omitempty" `
	StatusConsulta    string `json:"stsConsulta,omitempty"`
	Prioridade        string `json:"prioridade"`
	DtcMarcacao       string `json:"dtMarcacao,omitempty"`
	HrMarcacao        string `json:"hrMarcacao,omitempty"`
	IDCampus          uint64 `json:"id_Campus"`
	Campus            CampusUnip
	IDUsuario         uint64 `json:"id_Usuario,omitempty"`
	NomeUsuario       Usuario
	IDEspecialidade   uint64 `json:"id_Especialidade"`
	NomeEspecialidade Especialidade
}

// Preparar vai chamar os métodos para validar e formatar a publicação recebida
func (consulta *Consulta) Preparar() error {
	if erro := consulta.validar(); erro != nil {
		return erro
	}

	consulta.formatar()
	return nil
}

func (consulta *Consulta) validar() error {

	if consulta.DtcMarcacao == "" {
		return errors.New("O conteúdo é obrigatório e não pode estar em branco")
	}

	return nil
}

func (consulta *Consulta) formatar() {

	consulta.StatusConsulta = strings.TrimSpace(consulta.StatusConsulta)
	consulta.DtcMarcacao = strings.TrimSpace(consulta.DtcMarcacao)
}
