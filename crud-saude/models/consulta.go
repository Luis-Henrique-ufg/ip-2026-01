package models

// Consulta representa um agendamento de consulta no sistema de saúde
type Consulta struct {
	ID          int      `json:"id"`
	IDPaciente  int      `json:"id_paciente"`
	Paciente    Paciente `json:"paciente"`
	Data        string   `json:"data"`
	Medico      string   `json:"medico"`
	Observacoes string   `json:"observacoes"`
}
