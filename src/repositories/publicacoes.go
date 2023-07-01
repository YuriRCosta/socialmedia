package repositories

import (
	"api/src/models"
	"database/sql"
)

// Publicacoes representa um repositório de pubilcacoes
type Publicacoes struct {
	db *sql.DB
}

// NovoRepositorioDePublicacoes cria um repositório de pubilcacoes
func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

// Criar insere uma publicacao no banco de dados
func (repositorio Publicacoes) Criar(publicacao models.Publicacao) (uint64, error) {
	statement, err := repositorio.db.Prepare(
		"insert into publicacoes (titulo, conteudo, autor_id) values (?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if err != nil {
		return 0, err
	}

	ultimoIDInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ultimoIDInserido), nil
}

// Buscar traz as publicacoes que atendem um filtro de nome ou nick
func (repositorio Publicacoes) Buscar(usuarioID uint64) ([]models.Publicacao, error) {
	linhas, err := repositorio.db.Query(
		`select distinct p.*, u.nick from publicacoes p
		 inner join usuarios u on u.id = p.autor_id
		 inner join seguidores s on p.autor_id = s.usuario_id
		 where u.id = ? or s.seguidor_id = ?
		 order by 1 desc`,
		usuarioID, usuarioID,
	)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var publicacoes []models.Publicacao

	for linhas.Next() {
		var publicacao models.Publicacao

		if err = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); err != nil {
			return nil, err
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

// BuscarPorID traz as publicacoes que atendem um filtro de nome ou nick
func (repositorio Publicacoes) BuscarPorID(publicacaoID uint64) (models.Publicacao, error) {
	linhas, err := repositorio.db.Query(
		`select p.*, u.nick from publicacoes p
		 inner join usuarios u on u.id = p.autor_id
		 where p.id = ?`,
		publicacaoID,
	)
	if err != nil {
		return models.Publicacao{}, err
	}
	defer linhas.Close()

	var publicacoes models.Publicacao

	if linhas.Next() {
		if err = linhas.Scan(
			&publicacoes.ID,
			&publicacoes.Titulo,
			&publicacoes.Conteudo,
			&publicacoes.AutorID,
			&publicacoes.Curtidas,
			&publicacoes.CriadaEm,
			&publicacoes.AutorNick,
		); err != nil {
			return models.Publicacao{}, err
		}
	}

	return publicacoes, nil
}

// Atualizar altera as informações de uma publicacao no banco de dados
func (repositorio Publicacoes) Atualizar(publicacaoID uint64, publicacao models.Publicacao) error {
	statement, err := repositorio.db.Prepare(
		"update publicacoes set titulo = ?, conteudo = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID); err != nil {
		return err
	}

	return nil
}

// Deletar exclui uma publicacao do banco de dados
func (repositorio Publicacoes) Deletar(publicacaoID uint64) error {
	statement, err := repositorio.db.Prepare("delete from publicacoes where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(publicacaoID); err != nil {
		return err
	}

	return nil
}

// BuscarPorUsuario traz as publicacoes de um usuario especifico
func (repositorio Publicacoes) BuscarPorUsuario(usuarioID uint64) ([]models.Publicacao, error) {
	linhas, err := repositorio.db.Query(
		`select p.*, u.nick from publicacoes p
		 inner join usuarios u on u.id = p.autor_id
		 where u.id = ?`,
		usuarioID,
	)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var publicacoes []models.Publicacao

	for linhas.Next() {
		var publicacao models.Publicacao

		if err = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); err != nil {
			return nil, err
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

// Curtir adiciona uma curtida na publicacao
func (repositorio Publicacoes) Curtir(publicacaoID uint64) error {
	statement, err := repositorio.db.Prepare(
		"update publicacoes set curtidas = curtidas + 1 where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(publicacaoID); err != nil {
		return err
	}

	return nil
}

// Descurtir remove uma curtida na publicacao
func (repositorio Publicacoes) Descurtir(publicacaoID uint64) error {
	statement, err := repositorio.db.Prepare(
		"update publicacoes set curtidas = case when curtidas > 0 then curtidas - 1 else 0 end where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(publicacaoID); err != nil {
		return err
	}

	return nil
}
