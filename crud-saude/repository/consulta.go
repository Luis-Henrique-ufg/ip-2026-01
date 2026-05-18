package repository

import (
	"crud-saude/banco"
	"crud-saude/models"
	"fmt"
)

// CriarConsulta insere um novo agendamento no banco de dados
func CriarConsulta(c models.Consulta) (models.Consulta, error) {
	// Primeiro, verifica se o paciente existe
	var existe int
	err := banco.DB.QueryRow("SELECT COUNT(1) FROM pacientes WHERE id = ?", c.IDPaciente).Scan(&existe)
	if err != nil || existe == 0 {
		return models.Consulta{}, fmt.Errorf("paciente com ID %d não existe", c.IDPaciente)
	}

	query := `INSERT INTO consultas (id_paciente, data, medico, observacoes) 
	          VALUES (?, ?, ?, ?)`

	resultado, err := banco.DB.Exec(query, c.IDPaciente, c.Data, c.Medico, c.Observacoes)
	if err != nil {
		return models.Consulta{}, fmt.Errorf("erro ao agendar consulta")
	}

	id, _ := resultado.LastInsertId()
	c.ID = int(id)

	// Busca os dados do paciente para retornar a consulta completa
	paciente, err := BuscarPorID(c.IDPaciente)
	if err == nil {
		c.Paciente = paciente
	}

	return c, nil
}

// ListarTodasConsultas retorna todas as consultas com dados do paciente inclusos
func ListarTodasConsultas() ([]models.Consulta, error) {
	query := `SELECT c.id, c.id_paciente, c.data, c.medico, c.observacoes, 
	                 p.id, p.nome, p.idade, p.cpf, p.diagnostico
	          FROM consultas c
	          INNER JOIN pacientes p ON c.id_paciente = p.id
	          ORDER BY c.data DESC`

	rows, err := banco.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar consultas")
	}
	defer rows.Close()

	lista := []models.Consulta{}
	for rows.Next() {
		var c models.Consulta
		var p models.Paciente

		err := rows.Scan(
			&c.ID, &c.IDPaciente, &c.Data, &c.Medico, &c.Observacoes,
			&p.ID, &p.Nome, &p.Idade, &p.CPF, &p.Diagnostico,
		)
		if err != nil {
			return nil, fmt.Errorf("erro ao ler dados de consulta")
		}

		c.Paciente = p
		lista = append(lista, c)
	}

	return lista, nil
}

// BuscarConsultaPorID busca uma consulta pelo ID
func BuscarConsultaPorID(id int) (models.Consulta, error) {
	query := `SELECT c.id, c.id_paciente, c.data, c.medico, c.observacoes, 
	                 p.id, p.nome, p.idade, p.cpf, p.diagnostico
	          FROM consultas c
	          INNER JOIN pacientes p ON c.id_paciente = p.id
	          WHERE c.id = ?`

	var c models.Consulta
	var p models.Paciente

	err := banco.DB.QueryRow(query, id).Scan(
		&c.ID, &c.IDPaciente, &c.Data, &c.Medico, &c.Observacoes,
		&p.ID, &p.Nome, &p.Idade, &p.CPF, &p.Diagnostico,
	)
	if err != nil {
		return models.Consulta{}, fmt.Errorf("consulta com ID %d não encontrada", id)
	}

	c.Paciente = p
	return c, nil
}

// ListarConsultasPorPaciente retorna todas as consultas vinculadas a um paciente específico
func ListarConsultasPorPaciente(idPaciente int) ([]models.Consulta, error) {
	query := `SELECT c.id, c.id_paciente, c.data, c.medico, c.observacoes, 
	                 p.id, p.nome, p.idade, p.cpf, p.diagnostico
	          FROM consultas c
	          INNER JOIN pacientes p ON c.id_paciente = p.id
	          WHERE c.id_paciente = ?
	          ORDER BY c.data DESC`

	rows, err := banco.DB.Query(query, idPaciente)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar consultas do paciente")
	}
	defer rows.Close()

	lista := []models.Consulta{}
	for rows.Next() {
		var c models.Consulta
		var p models.Paciente

		err := rows.Scan(
			&c.ID, &c.IDPaciente, &c.Data, &c.Medico, &c.Observacoes,
			&p.ID, &p.Nome, &p.Idade, &p.CPF, &p.Diagnostico,
		)
		if err != nil {
			return nil, fmt.Errorf("erro ao ler dados de consulta")
		}

		c.Paciente = p
		lista = append(lista, c)
	}

	return lista, nil
}

// DeletarConsulta remove uma consulta pelo ID
func DeletarConsulta(id int) error {
	resultado, err := banco.DB.Exec("DELETE FROM consultas WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("erro ao cancelar consulta")
	}

	linhas, _ := resultado.RowsAffected()
	if linhas == 0 {
		return fmt.Errorf("consulta com ID %d não encontrada", id)
	}

	return nil
}
