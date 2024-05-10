package modelos

import (
	"api/src/seguranca"
	"errors"
	"strings"

	"github.com/badoux/checkmail"
)

type Usuario struct {
	ID            uint64 `json:"idUsuario,omitempty"`
	Nick          string `json:"usuario,omitempty"`
	Senha         string `json:"senha,omitempty"`
	Email         string `json:"email,omitempty"`
	Nome          string `json:"nome,omitempty"`
	Sobrenome     string `json:"sobrenome,omitempty"`
	Sexo          string `json:"sexo,omitempty"`
	Nascimento    string `json:"nascimento,omitempty"`
	Telefone      string `json:"telefone,omitempty"`
	CPF           string `json:"cpf,omitempty"`
	RA            string `json:"ra"`
	Acesso        string `json:"acesso,omitempty"`
	DtInclusao    string `json:"dtinclusao"`
	Campus        uint64 `json:"id_Campus"`
	Especialidade uint64 `json:"id_Especialidade"`
}

// Preparar vai chamar os metodos para validar e formatar o usuario recebido
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	if erro := usuario.formatar(etapa); erro != nil {
		return erro
	}
	return nil
}

// Validar vai validar se os campos estão preenchidos
func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}

	if usuario.Nick == "" {
		return errors.New("O nick é obrigatório e não pode estar em branco")
	}

	if usuario.Email == "" {
		return errors.New("O e-mail é obrigatório e não pode estar em branco")
	}

	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("O e-mail inserido é inválido")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("A senha é obrigatória e não pode estar em branco")
	}

	return nil
}

// formatar tira os espaços nas extremidades
func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}

		usuario.Senha = string(senhaComHash)
	}

	return nil
}
