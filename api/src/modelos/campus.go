package modelos

import (
	"database/sql"
	"errors"
	"strings"
	"time"
)

type CampusUnip struct {
	ID          uint64         `json:"idCampus,omitempty"`
	Nome        string         `json:"nome,omitempty"`
	Descricao   string         `json:"descricao,omitempty"`
	Endereco    string         `json:"endereco,omitempty"`
	Numero      string         `json:"numero,omitempty"`
	Complemento sql.NullString `json:"complemento,omitempty"`
	CEP         string         `json:"cep,omitempty"`
	Bairro      string         `json:"bairro,omitempty"`
	Cidade      string         `json:"cidade,omitempty"`
	Estado      string         `json:"estado,omitempty"`
	Regiao      string         `json:"regiao,omitempty"`
	Telefone    string         `json:"telefone,omitempty"`
	Localizacao string         `json:"localizacao,omitempty"`
	DtInclusao  time.Time      `json:"dtinclusao"`
}

// Preparar vai chamar os métodos para validar e formatar a publicação recebida
func (campusUnip *CampusUnip) Preparar(etapa string) error {
	if erro := campusUnip.validar(); erro != nil {
		return erro
	}

	campusUnip.formatar()
	return nil
}

func (campusUnip *CampusUnip) validar() error {
	if campusUnip.Nome == "" {
		return errors.New("O título é obrigatório e não pode estar em branco")
	}

	if campusUnip.Descricao == "" {
		return errors.New("O conteúdo é obrigatório e não pode estar em branco")
	}

	if campusUnip.Endereco == "" {
		return errors.New("O conteúdo é obrigatório e não pode estar em branco")
	}

	if campusUnip.Numero == "" {
		return errors.New("O conteúdo é obrigatório e não pode estar em branco")
	}

	if campusUnip.CEP == "" {
		return errors.New("O conteúdo é obrigatório e não pode estar em branco")
	}

	if campusUnip.Bairro == "" {
		return errors.New("O conteúdo é obrigatório e não pode estar em branco")
	}

	if campusUnip.Cidade == "" {
		return errors.New("O conteúdo é obrigatório e não pode estar em branco")
	}

	if campusUnip.Estado == "" {
		return errors.New("O conteúdo é obrigatório e não pode estar em branco")
	}

	if campusUnip.Regiao == "" {
		return errors.New("O conteúdo é obrigatório e não pode estar em branco")
	}

	if campusUnip.Telefone == "" {
		return errors.New("O conteúdo é obrigatório e não pode estar em branco")
	}

	if campusUnip.Localizacao == "" {
		return errors.New("O conteúdo é obrigatório e não pode estar em branco")
	}

	return nil
}

func (campusUnip *CampusUnip) formatar() {
	campusUnip.Nome = strings.TrimSpace(campusUnip.Nome)
	campusUnip.Descricao = strings.TrimSpace(campusUnip.Descricao)
	campusUnip.Endereco = strings.TrimSpace(campusUnip.Endereco)
	campusUnip.Numero = strings.TrimSpace(campusUnip.Numero)
	campusUnip.CEP = strings.TrimSpace(campusUnip.CEP)
	campusUnip.Bairro = strings.TrimSpace(campusUnip.Bairro)
	campusUnip.Cidade = strings.TrimSpace(campusUnip.Cidade)
	campusUnip.Estado = strings.TrimSpace(campusUnip.Estado)
	campusUnip.Regiao = strings.TrimSpace(campusUnip.Regiao)
	campusUnip.Telefone = strings.TrimSpace(campusUnip.Telefone)
	campusUnip.Localizacao = strings.TrimSpace(campusUnip.Localizacao)
}
