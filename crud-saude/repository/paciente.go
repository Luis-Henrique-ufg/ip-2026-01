package repository

import (
	"crud-saude/banco"
	"crud-saude/models"
	"fmt"
	"strings"
)

// Criar insere um novo paciente no banco
func Criar(p models.Paciente) (models.Paciente, error) {
	query := `INSERT INTO pacientes (nome, idade, cpf, diagnostico) 
	          VALUES (?, ?, ?, ?)`

	resultado, err := banco.DB.Exec(query, p.Nome, p.Idade, p.CPF, p.Diagnostico)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE") {
			return models.Paciente{}, fmt.Errorf("CPF já cadastrado")
		}
		return models.Paciente{}, fmt.Errorf("erro ao salvar paciente")
	}

	id, _ := resultado.LastInsertId()
	p.ID = int(id)
	return p, nil
}

// ListarTodos retorna todos os pacientes do banco
func ListarTodos() ([]models.Paciente, error) {
	rows, err := banco.DB.Query(
		"SELECT id, nome, idade, cpf, diagnostico FROM pacientes",
	)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar pacientes")
	}
	defer rows.Close()

	lista := []models.Paciente{}
	for rows.Next() {
		var p models.Paciente
		err := rows.Scan(&p.ID, &p.Nome, &p.Idade, &p.CPF, &p.Diagnostico)
		if err != nil {
			return nil, fmt.Errorf("erro ao ler dados")
		}
		lista = append(lista, p)
	}

	return lista, nil
}

// BuscarPorID retorna um paciente pelo ID
func BuscarPorID(id int) (models.Paciente, error) {
	var p models.Paciente
	query := "SELECT id, nome, idade, cpf, diagnostico FROM pacientes WHERE id = ?"

	err := banco.DB.QueryRow(query, id).Scan(
		&p.ID, &p.Nome, &p.Idade, &p.CPF, &p.Diagnostico,
	)
	if err != nil {
		return models.Paciente{}, fmt.Errorf("paciente com ID %d não encontrado", id)
	}

	return p, nil
}

// Atualizar modifica os dados de um paciente existente
func Atualizar(id int, p models.Paciente) (models.Paciente, error) {
	query := `UPDATE pacientes SET nome=?, idade=?, cpf=?, diagnostico=? WHERE id=?`

	resultado, err := banco.DB.Exec(query, p.Nome, p.Idade, p.CPF, p.Diagnostico, id)
	if err != nil {
		return models.Paciente{}, fmt.Errorf("erro ao atualizar paciente")
	}

	linhas, _ := resultado.RowsAffected()
	if linhas == 0 {
		return models.Paciente{}, fmt.Errorf("paciente com ID %d não encontrado", id)
	}

	p.ID = id
	return p, nil
}

// Deletar remove um paciente pelo ID
func Deletar(id int) error {
	resultado, err := banco.DB.Exec("DELETE FROM pacientes WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("erro ao deletar paciente")
	}

	linhas, _ := resultado.RowsAffected()
	if linhas == 0 {
		return fmt.Errorf("paciente com ID %d não encontrado", id)
	}

	return nil
}
