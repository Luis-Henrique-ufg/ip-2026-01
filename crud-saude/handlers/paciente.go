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

func Pacientes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		listarPacientes(w, r)
	case "POST":
		criarPaciente(w, r)
	default:
		responderErro(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

func PacientePorID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, ok := extrairID(w, r)
	if !ok {
		return
	}

	switch r.Method {
	case "GET":
		buscarPaciente(w, r, id)
	case "PUT":
		atualizarPaciente(w, r, id)
	case "DELETE":
		deletarPaciente(w, r, id)
	default:
		responderErro(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

// ═══════════════════════════════════════════════
//  HELPERS PRIVADOS
// ═══════════════════════════════════════════════

func extrairID(w http.ResponseWriter, r *http.Request) (int, bool) {
	idStr := strings.TrimPrefix(r.URL.Path, "/pacientes/")
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

func responderErro(w http.ResponseWriter, mensagem string, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(models.Resposta{Mensagem: mensagem})
}

func responderSucesso(w http.ResponseWriter, mensagem string, dados interface{}, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(models.Resposta{Mensagem: mensagem, Dados: dados})
}

func validar(p models.Paciente) string {
	if strings.TrimSpace(p.Nome) == "" {
		return "O campo 'nome' é obrigatório"
	}
	if strings.TrimSpace(p.CPF) == "" {
		return "O campo 'cpf' é obrigatório"
	}
	if p.Idade <= 0 || p.Idade > 150 {
		return "Idade deve ser entre 1 e 150"
	}
	return ""
}

// ═══════════════════════════════════════════════
//  OPERAÇÕES CRUD PRIVADAS
// ═══════════════════════════════════════════════

func criarPaciente(w http.ResponseWriter, r *http.Request) {
	var novo models.Paciente
	if err := json.NewDecoder(r.Body).Decode(&novo); err != nil {
		responderErro(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if msg := validar(novo); msg != "" {
		responderErro(w, msg, http.StatusBadRequest)
		return
	}

	criado, err := repository.Criar(novo)
	if err != nil {
		status := http.StatusInternalServerError
		if strings.Contains(err.Error(), "CPF") {
			status = http.StatusConflict
		}
		responderErro(w, err.Error(), status)
		return
	}

	responderSucesso(w, "Paciente criado com sucesso", criado, http.StatusCreated)
}

func listarPacientes(w http.ResponseWriter, r *http.Request) {
	lista, err := repository.ListarTodos()
	if err != nil {
		responderErro(w, err.Error(), http.StatusInternalServerError)
		return
	}
	responderSucesso(w, fmt.Sprintf("%d paciente(s) encontrado(s)", len(lista)), lista, http.StatusOK)
}

func buscarPaciente(w http.ResponseWriter, r *http.Request, id int) {
	paciente, err := repository.BuscarPorID(id)
	if err != nil {
		responderErro(w, err.Error(), http.StatusNotFound)
		return
	}
	responderSucesso(w, "Paciente encontrado", paciente, http.StatusOK)
}

func atualizarPaciente(w http.ResponseWriter, r *http.Request, id int) {
	var dados models.Paciente
	if err := json.NewDecoder(r.Body).Decode(&dados); err != nil {
		responderErro(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if msg := validar(dados); msg != "" {
		responderErro(w, msg, http.StatusBadRequest)
		return
	}

	atualizado, err := repository.Atualizar(id, dados)
	if err != nil {
		responderErro(w, err.Error(), http.StatusNotFound)
		return
	}

	responderSucesso(w, "Paciente atualizado com sucesso", atualizado, http.StatusOK)
}

func deletarPaciente(w http.ResponseWriter, r *http.Request, id int) {
	err := repository.Deletar(id)
	if err != nil {
		responderErro(w, err.Error(), http.StatusNotFound)
		return
	}
	responderSucesso(w, fmt.Sprintf("Paciente ID %d removido com sucesso", id), nil, http.StatusOK)
}
