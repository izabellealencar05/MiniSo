package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func listFiles(currentUser *User) {
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Erro ao listar arquivos:", err)
		return
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func createFile(filePath string, currentUser *User) {
	dir := filepath.Dir(filePath)
	if dir != "." {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			fmt.Println("Erro ao criar diretório:", err)
			return
		}
	}

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Erro ao criar arquivo:", err)
		return
	}
	defer file.Close()

	content := "Conteúdo aleatório gerado"
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
	}
	fmt.Println("Arquivo criado:", filePath)
}

func deleteFile(filePath string, currentUser *User) {
	err := os.Remove(filePath)
	if err != nil {
		fmt.Println("Erro ao apagar arquivo:", err)
		return
	}
	fmt.Println("Arquivo apagado:", filePath)
}

func createDirectory(dirPath string, currentUser *User) {
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		fmt.Println("Erro ao criar diretório:", err)
		return
	}
	fmt.Println("Diretório criado:", dirPath)
}

func deleteDirectory(dirPath string, currentUser *User) {
	err := os.Remove(dirPath)
	if err != nil {
		fmt.Println("Erro ao apagar diretório:", err)
		return
	}
	fmt.Println("Diretório apagado:", dirPath)
}

func deleteDirectoryForce(dirPath string, currentUser *User) {
	err := os.RemoveAll(dirPath)
	if err != nil {
		fmt.Println("Erro ao apagar diretório forçadamente:", err)
		return
	}
	fmt.Println("Diretório apagado (forçado):", dirPath)
}
