package banco

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

// DB é a conexão compartilhada com o banco de dados
var DB *sql.DB

// Iniciar abre a conexão e cria as tabelas necessárias
func Iniciar() {
	var err error

	DB, err = sql.Open("sqlite", "./banco.db")
	if err != nil {
		log.Fatal("Erro ao abrir banco:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Erro ao conectar no banco:", err)
	}

	criarTabelas()
	fmt.Println("  Banco de dados conectado! ✓")
}

func criarTabelas() {
	queryPacientes := `
	CREATE TABLE IF NOT EXISTS pacientes (
		id          INTEGER PRIMARY KEY AUTOINCREMENT,
		nome        TEXT    NOT NULL,
		idade       INTEGER NOT NULL,
		cpf         TEXT    NOT NULL UNIQUE,
		diagnostico TEXT
	);`

	_, err := DB.Exec(queryPacientes)
	if err != nil {
		log.Fatal("Erro ao criar tabela pacientes:", err)
	}

	queryConsultas := `
	CREATE TABLE IF NOT EXISTS consultas (
		id          INTEGER PRIMARY KEY AUTOINCREMENT,
		id_paciente INTEGER NOT NULL,
		data        TEXT    NOT NULL,
		medico      TEXT    NOT NULL,
		observacoes TEXT,
		FOREIGN KEY(id_paciente) REFERENCES pacientes(id) ON DELETE CASCADE
	);`

	_, err = DB.Exec(queryConsultas)
	if err != nil {
		log.Fatal("Erro ao criar tabela consultas:", err)
	}
}
