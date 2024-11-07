<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <meta charset="UTF-8">
    <title>MiniSO - Sistema de Gerenciamento de Usuários</title>
</head>
<body>

<h1>MiniSO - Sistema de Gerenciamento de Usuários</h1>

<p>Este projeto implementa um sistema de gerenciamento de usuários em um ambiente MiniSO com suporte a funcionalidades de criação, login, logout e exclusão de contas. Ele foi projetado com segurança aprimorada para o armazenamento de senhas.</p>

<h2>Funcionalidades</h2>
<ul>
    <li><strong>Cadastro de Usuário:</strong> Caso não existam usuários cadastrados, o sistema solicita a criação de um novo usuário com senha.</li>
    <li><strong>Armazenamento Seguro de Senhas:</strong> As senhas são armazenadas em hash SHA-512 com um salt aleatório, conforme práticas de segurança.</li>
    <li><strong>Login:</strong> Permite o login de usuários registrados. A senha digitada não é exibida (com asteriscos ou oculta).</li>
    <li><strong>Exclusão de Último Usuário:</strong> Se o último usuário for excluído, o sistema retornará ao passo de cadastro de um novo usuário.</li>
    <li><strong>Menu de Login:</strong> Um menu inicial permite ao usuário escolher entre cadastro, login e sair.</li>
    <li><strong>Logout:</strong> Permite que o usuário logado saia de sua conta.</li>
</ul>

<h2>Como Executar</h2>
<ol>
    <li>Clone o repositório: <code>git clone &lt;url-do-repositorio&gt;</code></li>
    <li>Instale as dependências (se houver).</li>
    <li>Execute o script principal com o comando <code>python miniSO.py</code>.</li>
</ol>

<h2>Detalhes de Segurança</h2>
<p>As senhas dos usuários são protegidas através da combinação de salt e hash SHA-512. Isso garante que as senhas não sejam armazenadas em texto puro, dificultando o acesso indevido.</p>

<h2>Licença</h2>
<p>Este projeto está licenciado sob a licença MIT. Consulte o arquivo <code>LICENSE</code> para mais detalhes.</p>

</body>
</html>
