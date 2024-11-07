package main

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
)

// Função para ler a senha sem exibir
func readPassword() (string, error) {
	// Lê a senha sem exibi-la no terminal
	password, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}
	fmt.Println() // Apenas para adicionar uma linha após a leitura da senha
	return string(password), nil
}

func main() {
	// Carrega os usuários existentes
	users, err := loadUsers()
	if err != nil {
		fmt.Println("Erro ao carregar usuários:", err)
		os.Exit(1)
	}

	var currentUser *User
	if len(users) == 0 {
		// Caso não exista nenhum usuário, cria um novo usuário
		var username, password string
		fmt.Print("Criação de usuário - Digite o nome: ")
		fmt.Scanln(&username)
		fmt.Print("Digite a senha: ")
		password, err = readPassword() // Lê a senha sem exibir
		if err != nil {
			fmt.Println("Erro ao ler a senha:", err)
			os.Exit(1)
		}

		user, err := createUser(username, password)
		if err != nil {
			fmt.Println("Erro ao criar o usuário:", err)
			os.Exit(1)
		}
		saveUser(user)
		currentUser = user
	} else {
		// Caso já existam usuários, solicita login
		currentUser, err = login(users)
		if err != nil {
			fmt.Println("Erro ao fazer login:", err)
			os.Exit(1)
		}
	}

	// Mensagem de boas-vindas
	fmt.Printf("Bem-vindo, %s!\n", currentUser.Username)

	// Usando bufio.Reader para ler comandos com espaços
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erro ao ler comando:", err)
			continue
		}
		command = strings.TrimSpace(command) // Remove a quebra de linha

		// Verifica se o comando não está vazio
		if command != "" {
			executeCommand(command, currentUser)
		} else {
			fmt.Println("Comando vazio.")
		}
	}
}

func executeCommand(command string, currentUser *User) {
	parts := strings.Fields(command) // Divide o comando em partes

	if len(parts) == 0 {
		fmt.Println("Comando vazio.")
		return
	}

	switch parts[0] {
	case "listar":
		listFiles(currentUser)
	case "criar":
		if len(parts) < 3 {
			fmt.Println("Uso: criar <arquivo/diretorio> <nome>")
			return
		}
		if parts[1] == "arquivo" {
			createFile(parts[2], currentUser)
		} else if parts[1] == "diretorio" {
			createDirectory(parts[2], currentUser)
		} else {
			fmt.Println("Comando inválido. Use 'arquivo' ou 'diretorio'.")
		}
	case "apagar":
		if len(parts) < 3 {
			fmt.Println("Uso: apagar <arquivo/diretorio> <nome>")
			return
		}
		if parts[1] == "arquivo" {
			deleteFile(parts[2], currentUser)
		} else if parts[1] == "diretorio" {
			if len(parts) == 4 && parts[3] == "--force" {
				// Chama o processo de criação e alocação de memória antes de apagar o diretório
				process := createProcess("apagar diretório forçado", 64) // Aloca 64 bytes para o processo
				if process != nil {
					deleteDirectoryForce(parts[2], currentUser)
				}
			} else {
				deleteDirectory(parts[2], currentUser)
			}
		} else {
			fmt.Println("Comando inválido. Use 'arquivo' ou 'diretorio'.")
		}
	default:
		fmt.Println("Comando desconhecido:", parts[0])
	}
}
