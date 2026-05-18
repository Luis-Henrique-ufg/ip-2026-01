package handlers

import (
	"crud-saude/models"
	"crud-saude/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// ═══════════════════════════════════════════════
//  ROTEADORES
// ═══════════════════════════════════════════════

// Consultas gerencia a listagem e criação de consultas
func Consultas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		listarConsultas(w, r)
	case "POST":
		criarConsulta(w, r)
	default:
		responderErro(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

// ConsultaPorID gerencia busca e cancelamento de uma consulta específica
func ConsultaPorID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, ok := extrairConsultaID(w, r)
	if !ok {
		return
	}

	switch r.Method {
	case "GET":
		buscarConsulta(w, r, id)
	case "DELETE":
		deletarConsulta(w, r, id)
	default:
		responderErro(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

// ConsultasPorPaciente retorna todas as consultas vinculadas a um paciente
func ConsultasPorPaciente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		responderErro(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	idPaciente, ok := extrairPacienteIDDeConsulta(w, r)
	if !ok {
		return
	}

	listarConsultasPorPaciente(w, r, idPaciente)
}

// ═══════════════════════════════════════════════
//  HELPERS PRIVADOS
// ═══════════════════════════════════════════════

func extrairConsultaID(w http.ResponseWriter, r *http.Request) (int, bool) {
	idStr := strings.TrimPrefix(r.URL.Path, "/consultas/")
	if idStr == "" {
		responderErro(w, "ID não informado", http.StatusBadRequest)
		return 0, false
	}
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		responderErro(w, "ID inválido", http.StatusBadRequest)
		return 0, false
	}
	return id, true
}

func extrairPacienteIDDeConsulta(w http.ResponseWriter, r *http.Request) (int, bool) {
	idStr := strings.TrimPrefix(r.URL.Path, "/pacientes/consultas/")
	if idStr == "" {
		responderErro(w, "ID do paciente não informado", http.StatusBadRequest)
		return 0, false
	}
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		responderErro(w, "ID do paciente inválido", http.StatusBadRequest)
		return 0, false
	}
	return id, true
}

func validarConsulta(c models.Consulta) string {
	if c.IDPaciente <= 0 {
		return "O campo 'id_paciente' é obrigatório e deve ser válido"
	}
	if strings.TrimSpace(c.Data) == "" {
		return "O campo 'data' é obrigatório"
	}
	if strings.TrimSpace(c.Medico) == "" {
		return "O campo 'medico' é obrigatório"
	}
	return ""
}

// ═══════════════════════════════════════════════
//  OPERAÇÕES CRUD PRIVADAS
// ═══════════════════════════════════════════════

func criarConsulta(w http.ResponseWriter, r *http.Request) {
	var nova models.Consulta
	if err := json.NewDecoder(r.Body).Decode(&nova); err != nil {
		responderErro(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if msg := validarConsulta(nova); msg != "" {
		responderErro(w, msg, http.StatusBadRequest)
		return
	}

	criada, err := repository.CriarConsulta(nova)
	if err != nil {
		responderErro(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responderSucesso(w, "Consulta agendada com sucesso", criada, http.StatusCreated)
}

func listarConsultas(w http.ResponseWriter, r *http.Request) {
	lista, err := repository.ListarTodasConsultas()
	if err != nil {
		responderErro(w, err.Error(), http.StatusInternalServerError)
		return
	}
	responderSucesso(w, fmt.Sprintf("%d consulta(s) encontrada(s)", len(lista)), lista, http.StatusOK)
}

func buscarConsulta(w http.ResponseWriter, r *http.Request, id int) {
	consulta, err := repository.BuscarConsultaPorID(id)
	if err != nil {
		responderErro(w, err.Error(), http.StatusNotFound)
		return
	}
	responderSucesso(w, "Consulta encontrada", consulta, http.StatusOK)
}

func listarConsultasPorPaciente(w http.ResponseWriter, r *http.Request, idPaciente int) {
	lista, err := repository.ListarConsultasPorPaciente(idPaciente)
	if err != nil {
		responderErro(w, err.Error(), http.StatusInternalServerError)
		return
	}
	responderSucesso(w, fmt.Sprintf("%d consulta(s) encontrada(s) para este paciente", len(lista)), lista, http.StatusOK)
}

func deletarConsulta(w http.ResponseWriter, r *http.Request, id int) {
	err := repository.DeletarConsulta(id)
	if err != nil {
		responderErro(w, err.Error(), http.StatusNotFound)
		return
	}
	responderSucesso(w, fmt.Sprintf("Consulta ID %d cancelada com sucesso", id), nil, http.StatusOK)
}
