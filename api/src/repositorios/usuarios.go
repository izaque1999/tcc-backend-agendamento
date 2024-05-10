package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

type Usuarios struct {
	db *sql.DB
}

// esta função cria um repositorio de usuarios
func NewRepositoryUsers(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um usuário no banco de dados
func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (usuario, senha, email, nome, sobrenome, sexo, nascimento, telefone, cpf, ra, acesso) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nick, usuario.Senha, usuario.Email, usuario.Nome, usuario.Sobrenome, usuario.Sexo, usuario.Nascimento, usuario.Telefone, usuario.CPF, usuario.RA, usuario.Acesso)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil

}

func (repositorio Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)

	linhas, erro := repositorio.db.Query(
		"select idUsuario, usuario, senha, email, nome, sobrenome, sexo, nascimento, telefone, cpf, ra, acesso, dtinclusao from usuarios where nome LIKE ? or usuario LIKE ?",
		nomeOuNick, nomeOuNick,
	)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nick,
			&usuario.Senha,
			&usuario.Email,
			&usuario.Nome,
			&usuario.Sobrenome,
			&usuario.Sexo,
			&usuario.Nascimento,
			&usuario.Telefone,
			&usuario.CPF,
			&usuario.RA,
			&usuario.Acesso,
			&usuario.DtInclusao,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)

	}

	return usuarios, erro

}

// Buscar o usuario pelo ID
func (repositorio Usuarios) BuscarPorID(ID uint64) (modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		"select idUsuario, usuario, senha, email, nome, sobrenome, sexo, nascimento, telefone, cpf, ra, acesso, dtinclusao from usuarios where idUsuario = ?",
		ID,
	)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nick,
			&usuario.Senha,
			&usuario.Email,
			&usuario.Nome,
			&usuario.Sobrenome,
			&usuario.Sexo,
			&usuario.Nascimento,
			&usuario.Telefone,
			&usuario.CPF,
			&usuario.RA,
			&usuario.Acesso,
			&usuario.DtInclusao,
		); erro != nil {
			return modelos.Usuario{}, erro
		}

	}

	return usuario, nil

}

// Atualizar alterar informações de um usuario no banco de dados
func (repositorio Usuarios) Atualizar(ID uint64, usuario modelos.Usuario) error {
	statement, erro := repositorio.db.Prepare(
		"update usuarios set id_Campus = ? , id_Especialidade = ? where idUsuario = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuario.Campus, usuario.Especialidade, ID); erro != nil {
		return erro
	}

	return nil
}

// Deletar exclui as informações de um usuário no banco de dados
func (repositorio Usuarios) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from usuarios where idUsuario = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query("select idUsuario, usuario, senha, email, nome, sobrenome, sexo, nascimento, telefone, cpf, ra, acesso, dtinclusao from usuarios where email = ?", email)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nick,
			&usuario.Senha,
			&usuario.Email,
			&usuario.Nome,
			&usuario.Sobrenome,
			&usuario.Sexo,
			&usuario.Nascimento,
			&usuario.Telefone,
			&usuario.CPF,
			&usuario.RA,
			&usuario.Acesso,
			&usuario.DtInclusao,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil

}

func (repositorio Usuarios) BuscarSenha(usuarioID uint64) (string, error) {
	linha, erro := repositorio.db.Query("select senha from usuarios where idUsuario = ?", usuarioID)
	if erro != nil {
		return "", erro
	}
	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.Senha); erro != nil {
			return "", erro
		}
	}

	return usuario.Senha, nil
}

// AtualizarSenha altera a senha de um usuário
func (repositorio Usuarios) AtualizarSenha(usuarioID uint64, senha string) error {
	statement, erro := repositorio.db.Prepare("update usuarios set senha = ? where idUsuario = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(senha, usuarioID); erro != nil {
		return erro
	}

	return nil
}
