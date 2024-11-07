package main

import (
	"bufio"
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"os"
)

type User struct {
	Username     string
	PasswordHash string
	Salt         string
}

func generateSalt() (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(salt), nil
}

func hashPassword(password, salt string) string {
	hasher := sha512.New()
	hasher.Write([]byte(password + salt))
	return hex.EncodeToString(hasher.Sum(nil))
}

func verifyPassword(storedHash, salt, password string) bool {
	return storedHash == hashPassword(password, salt)
}

func createUser(username, password string) (*User, error) {
	salt, err := generateSalt()
	if err != nil {
		return nil, err
	}
	passwordHash := hashPassword(password, salt)
	return &User{
		Username:     username,
		PasswordHash: passwordHash,
		Salt:         salt,
	}, nil
}

func saveUser(user *User) error {
	file, err := os.OpenFile("users.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("%s,%s,%s\n", user.Username, user.PasswordHash, user.Salt))
	return err
}

func loadUsers() ([]User, error) {
	if _, err := os.Stat("users.txt"); os.IsNotExist(err) {
		_, err := os.Create("users.txt")
		if err != nil {
			return nil, fmt.Errorf("erro ao criar o arquivo de usuários: %v", err)
		}
	}

	file, err := os.Open("users.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var users []User
	var username, passwordHash, salt string
	for {
		_, err := fmt.Fscanf(file, "%s,%s,%s\n", &username, &passwordHash, &salt)
		if err != nil {
			break
		}
		users = append(users, User{Username: username, PasswordHash: passwordHash, Salt: salt})
	}
	return users, nil
}

func login(users []User) (*User, error) {
	var username, password string
	fmt.Print("Username: ")
	fmt.Scanln(&username)

	// Lê a senha com a exibição de asteriscos
	fmt.Print("Password: ")
	password = readPasswordWithAsterisks()
	fmt.Println() // Pula a linha após a digitação

	for _, user := range users {
		if user.Username == username && verifyPassword(user.PasswordHash, user.Salt, password) {
			return &user, nil
		}
	}

	return nil, fmt.Errorf("usuário ou senha inválidos")
}

// Função para ler senha com a exibição de asteriscos no terminal
func readPasswordWithAsterisks() string {
	var password string
	reader := bufio.NewReader(os.Stdin)

	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		if char == '\n' || char == '\r' {
			break
		}
		if char == 8 || char == 127 { // Backspace
			if len(password) > 0 {
				password = password[:len(password)-1]
				fmt.Print("\b \b") // Apaga o último asterisco
			}
		} else {
			password += string(char)
			fmt.Print("*") // Exibe o asterisco
		}
	}

	return password
}
