// ════════════════════════════════════════════════
//  CONFIGURAÇÃO
// ════════════════════════════════════════════════

// Endereço base da API Go
const API = 'http://localhost:8080';

// ════════════════════════════════════════════════
//  NAVEGAÇÃO ENTRE SEÇÕES
// ════════════════════════════════════════════════

function mostrarSecao(nome) {
    // Esconde todas as seções
    document.querySelectorAll('section').forEach(s => s.classList.add('escondido'));

    // Remove o "ativo" de todos os botões de navegação
    document.querySelectorAll('.nav-btn').forEach(b => b.classList.remove('ativo'));

    // Mostra a seção escolhida
    document.getElementById('secao-' + nome).classList.remove('escondido');

    // Marca o botão como ativo
    document.getElementById('btn-nav-' + nome).classList.add('ativo');

    // Carrega os dados da seção
    if (nome === 'pacientes') carregarPacientes();
    if (nome === 'consultas') {
        carregarConsultas();
        carregarPacientesNoSelect();
    }
}

// ════════════════════════════════════════════════
//  MENSAGENS DE FEEDBACK
// ════════════════════════════════════════════════

function mostrarMensagem(texto, tipo) {
    const el = document.getElementById('mensagem');
    el.textContent = texto;
    el.className = 'mensagem ' + tipo; // 'sucesso' ou 'erro'
    el.style.display = 'block';

    // Esconde automaticamente após 4 segundos
    setTimeout(() => { el.style.display = 'none'; }, 4000);
}

// ════════════════════════════════════════════════
//  PACIENTES — CRUD COMPLETO
// ════════════════════════════════════════════════

// Guarda o ID do paciente sendo editado (null = criando novo)
let editandoID = null;

// ── LISTAR ────────────────────────────────────
async function carregarPacientes() {
    document.getElementById('lista-pacientes').innerHTML =
        '<p class="carregando">Carregando...</p>';

    try {
        // Faz GET /pacientes
        const res = await fetch(`${API}/pacientes`);
        const json = await res.json();
        const lista = json.dados || [];

        const div = document.getElementById('lista-pacientes');

        // Se não há pacientes, mostra mensagem
        if (lista.length === 0) {
            div.innerHTML = '<p class="vazio">📭 Nenhum paciente cadastrado ainda.</p>';
            return;
        }

        // Monta a tabela com os dados
        div.innerHTML = `
            <div class="tabela-wrapper">
                <table>
                    <thead>
                        <tr>
                            <th>ID</th>
                            <th>Nome</th>
                            <th>Idade</th>
                            <th>CPF</th>
                            <th>Diagnóstico</th>
                            <th>Ações</th>
                        </tr>
                    </thead>
                    <tbody>
                        ${lista.map(p => `
                            <tr>
                                <td><span class="badge-id">#${p.id}</span></td>
                                <td><strong>${p.nome}</strong></td>
                                <td>${p.idade} anos</td>
                                <td>${p.cpf}</td>
                                <td>${p.diagnostico || '<em style="color:#a0aec0">—</em>'}</td>
                                <td>
                                    <div class="acoes">
                                        <button
                                            class="btn-editar"
                                            onclick="prepararEdicao(${p.id}, '${escapar(p.nome)}', ${p.idade}, '${p.cpf}', '${escapar(p.diagnostico)}')">
                                            ✏️ Editar
                                        </button>
                                        <button
                                            class="btn-deletar"
                                            onclick="deletarPaciente(${p.id}, '${escapar(p.nome)}')">
                                            🗑️ Deletar
                                        </button>
                                    </div>
                                </td>
                            </tr>
                        `).join('')}
                    </tbody>
                </table>
            </div>
        `;

    } catch (err) {
        document.getElementById('lista-pacientes').innerHTML =
            '<p class="vazio">❌ Erro ao conectar com o servidor.</p>';
    }
}

// ── CREATE / UPDATE ───────────────────────────
document.getElementById('form-paciente').addEventListener('submit', async function (e) {
    // Impede o comportamento padrão do form (recarregar a página)
    e.preventDefault();

    // Pega os valores dos campos
    const paciente = {
        nome: document.getElementById('nome').value.trim(),
        idade: parseInt(document.getElementById('idade').value),
        cpf: document.getElementById('cpf').value.trim(),
        diagnostico: document.getElementById('diagnostico').value.trim()
    };

    try {
        let res;

        if (editandoID) {
            // ── ATUALIZAR (PUT) ──
            res = await fetch(`${API}/pacientes/${editandoID}`, {
                method: 'PUT',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(paciente)
            });
        } else {
            // ── CRIAR (POST) ──
            res = await fetch(`${API}/pacientes`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(paciente)
            });
        }

        const json = await res.json();

        if (res.ok) {
            mostrarMensagem(json.mensagem, 'sucesso');
            cancelarEdicao();         // limpa o formulário
            carregarPacientes();      // recarrega a tabela
        } else {
            mostrarMensagem(json.mensagem, 'erro');
        }

    } catch (err) {
        mostrarMensagem('Erro ao conectar com o servidor.', 'erro');
    }
});

// ── PREPARAR EDIÇÃO ───────────────────────────
function prepararEdicao(id, nome, idade, cpf, diagnostico) {
    // Guarda qual paciente está sendo editado
    editandoID = id;

    // Preenche o formulário com os dados atuais
    document.getElementById('nome').value = nome;
    document.getElementById('idade').value = idade;
    document.getElementById('cpf').value = cpf;
    document.getElementById('diagnostico').value = diagnostico;

    // Muda o visual do formulário para modo edição
    document.getElementById('titulo-form-paciente').textContent = `✏️ Editando paciente #${id}`;
    document.getElementById('btn-submit-paciente').textContent = '💾 Salvar Alterações';
    document.getElementById('btn-cancelar').style.display = 'block';

    // Sobe a página para o formulário
    window.scrollTo({ top: 0, behavior: 'smooth' });
}

// ── CANCELAR EDIÇÃO ───────────────────────────
function cancelarEdicao() {
    editandoID = null;

    // Limpa o formulário
    document.getElementById('form-paciente').reset();

    // Restaura o visual original
    document.getElementById('titulo-form-paciente').textContent = 'Cadastrar Paciente';
    document.getElementById('btn-submit-paciente').textContent = '➕ Cadastrar';
    document.getElementById('btn-cancelar').style.display = 'none';
}

// ── DELETAR ───────────────────────────────────
async function deletarPaciente(id, nome) {
    // Pede confirmação antes de deletar
    if (!confirm(`Tem certeza que deseja remover o paciente "${nome}"?`)) return;

    try {
        const res = await fetch(`${API}/pacientes/${id}`, { method: 'DELETE' });
        const json = await res.json();

        mostrarMensagem(json.mensagem, res.ok ? 'sucesso' : 'erro');
        carregarPacientes();

    } catch (err) {
        mostrarMensagem('Erro ao conectar com o servidor.', 'erro');
    }
}

// ════════════════════════════════════════════════
//  CONSULTAS — CREATE + READ + DELETE
// ════════════════════════════════════════════════

// ── PREENCHER SELECT COM PACIENTES ────────────
async function carregarPacientesNoSelect() {
    try {
        const res = await fetch(`${API}/pacientes`);
        const json = await res.json();
        const lista = json.dados || [];

        const select = document.getElementById('select-paciente');
        select.innerHTML = '<option value="">Selecione um paciente *</option>';

        lista.forEach(p => {
            const option = document.createElement('option');
            option.value = p.id;
            option.textContent = `${p.nome} (CPF: ${p.cpf})`;
            select.appendChild(option);
        });

    } catch (err) {
        console.error('Erro ao carregar pacientes no select:', err);
    }
}

// ── LISTAR CONSULTAS ──────────────────────────
async function carregarConsultas() {
    document.getElementById('lista-consultas').innerHTML =
        '<p class="carregando">Carregando...</p>';

    try {
        const res = await fetch(`${API}/consultas`);
        const json = await res.json();
        const lista = json.dados || [];

        const div = document.getElementById('lista-consultas');

        if (lista.length === 0) {
            div.innerHTML = '<p class="vazio">📭 Nenhuma consulta agendada ainda.</p>';
            return;
        }

        div.innerHTML = `
            <div class="tabela-wrapper">
                <table>
                    <thead>
                        <tr>
                            <th>ID</th>
                            <th>Paciente</th>
                            <th>Data</th>
                            <th>Médico</th>
                            <th>Observações</th>
                            <th>Ações</th>
                        </tr>
                    </thead>
                    <tbody>
                        ${lista.map(c => `
                            <tr>
                                <td><span class="badge-id">#${c.id}</span></td>
                                <td><strong>${c.paciente.nome}</strong></td>
                                <td>${formatarData(c.data)}</td>
                                <td>${c.medico}</td>
                                <td>${c.observacoes || '<em style="color:#a0aec0">—</em>'}</td>
                                <td>
                                    <div class="acoes">
                                        <button
                                            class="btn-deletar"
                                            onclick="deletarConsulta(${c.id})">
                                            🗑️ Deletar
                                        </button>
                                    </div>
                                </td>
                            </tr>
                        `).join('')}
                    </tbody>
                </table>
            </div>
        `;

    } catch (err) {
        document.getElementById('lista-consultas').innerHTML =
            '<p class="vazio">❌ Erro ao conectar com o servidor.</p>';
    }
}

// ── CRIAR CONSULTA ────────────────────────────
document.getElementById('form-consulta').addEventListener('submit', async function (e) {
    e.preventDefault();

    const consulta = {
        id_paciente: parseInt(document.getElementById('select-paciente').value),
        data: document.getElementById('data').value,
        medico: document.getElementById('medico').value.trim(),
        observacoes: document.getElementById('observacoes').value.trim()
    };

    // Valida seleção do paciente
    if (!consulta.id_paciente) {
        mostrarMensagem('Selecione um paciente.', 'erro');
        return;
    }

    try {
        const res = await fetch(`${API}/consultas`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(consulta)
        });
        const json = await res.json();

        if (res.ok) {
            mostrarMensagem(json.mensagem, 'sucesso');
            document.getElementById('form-consulta').reset();
            carregarConsultas();
        } else {
            mostrarMensagem(json.mensagem, 'erro');
        }

    } catch (err) {
        mostrarMensagem('Erro ao conectar com o servidor.', 'erro');
    }
});

// ── DELETAR CONSULTA ──────────────────────────
async function deletarConsulta(id) {
    if (!confirm(`Tem certeza que deseja remover a consulta #${id}?`)) return;

    try {
        const res = await fetch(`${API}/consultas/${id}`, { method: 'DELETE' });
        const json = await res.json();

        mostrarMensagem(json.mensagem, res.ok ? 'sucesso' : 'erro');
        carregarConsultas();

    } catch (err) {
        mostrarMensagem('Erro ao conectar com o servidor.', 'erro');
    }
}

// ════════════════════════════════════════════════
//  UTILITÁRIOS
// ════════════════════════════════════════════════

// Formata "2024-03-15" como "15/03/2024"
function formatarData(data) {
    if (!data) return '—';
    const partes = data.split('-');
    if (partes.length !== 3) return data;
    return `${partes[2]}/${partes[1]}/${partes[0]}`;
}

// Escapa aspas simples para uso seguro dentro de onclick=""
function escapar(texto) {
    if (!texto) return '';
    return texto.replace(/'/g, "\\'");
}

// ════════════════════════════════════════════════
//  INICIALIZAÇÃO
// ════════════════════════════════════════════════

// Carrega a seção inicial ao abrir a página
carregarPacientes();