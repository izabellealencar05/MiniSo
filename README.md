<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <meta charset="UTF-8">
    <title>MiniSO</title>
</head>
<body>

<h2>Funcionalidades que ainda precisam ser feitas</h2>
<ul>
    <li><strong>Cadastro de Usuário:</strong> Caso não existam usuários cadastrados, o sistema solicita a criação de um novo usuário com senha.</li>
    <li><strong>Armazenamento Seguro de Senhas:</strong> As senhas são armazenadas em hash SHA-512 com um salt aleatório, conforme práticas de segurança.</li>
    <li><strong>Login:</strong> Permite o login de usuários registrados. A senha digitada não é exibida (com asteriscos ou oculta).</li>
    <li><strong>Exclusão de Último Usuário:</strong> Se o último usuário for excluído, o sistema retornará ao passo de cadastro de um novo usuário.</li>
    <li><strong>Menu de Login:</strong> Um menu inicial permite ao usuário escolher entre cadastro, login e sair.</li>
    <li><strong>Logout:</strong> Permite que o usuário logado saia de sua conta.</li>
</ul>

<h2>Detalhes de Segurança</h2>
<p>As senhas dos usuários são protegidas através da combinação de salt e hash SHA-512. Isso garante que as senhas não sejam armazenadas em texto puro, dificultando o acesso indevido.</p>

</body>
</html>
