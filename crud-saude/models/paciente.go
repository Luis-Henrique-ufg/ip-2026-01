package models

// Paciente representa um paciente no sistema de saúde
type Paciente struct {
	ID          int    `json:"id"`
	Nome        string `json:"nome"`
	Idade       int    `json:"idade"`
	CPF         string `json:"cpf"`
	Diagnostico string `json:"diagnostico"`
}

// Resposta é o formato padrão de todas as respostas da API
type Resposta struct {
	Mensagem string      `json:"mensagem"`
	Dados    interface{} `json:"dados,omitempty"`
}
