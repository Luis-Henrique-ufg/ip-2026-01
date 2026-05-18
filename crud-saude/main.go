package main

import (
	"crud-saude/banco"
	"crud-saude/handlers"
	"fmt"
	"net/http"
)

func main() {
	banco.Iniciar()
	defer banco.DB.Close()

	// ── API ──────────────────────────────────────
	http.HandleFunc("/pacientes", handlers.Pacientes)
	http.HandleFunc("/pacientes/", handlers.PacientePorID)
	http.HandleFunc("/consultas", handlers.Consultas)
	http.HandleFunc("/consultas/", handlers.ConsultaPorID)
	http.HandleFunc("/pacientes/consultas/", handlers.ConsultasPorPaciente)

	// ── FRONTEND ─────────────────────────────────
	// Serve todos os arquivos da pasta ./static
	http.Handle("/", http.FileServer(http.Dir("./static")))

	fmt.Println("╔══════════════════════════════════════╗")
	fmt.Println("║   Sistema de Saúde rodando!          ║")
	fmt.Println("║   Acesse: http://localhost:8080      ║")
	fmt.Println("╚══════════════════════════════════════╝")

	http.ListenAndServe(":8080", nil)
}
