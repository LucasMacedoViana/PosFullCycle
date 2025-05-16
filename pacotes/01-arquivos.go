package pacotes

import (
	"bufio"
	"fmt"
	"os"
)

func Os() {
	// manipulaçao de arquivos
	createFile()
	readFile()
	removeFile()
}

func createFile() {
	// cria um arquivo
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}
	// escreve no arquivo
	tamanho1, err := f.WriteString("Olá, mundo!")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Tamanho: %d bytes\n", tamanho1)

	// escreve no arquivo
	tamanho2, err := f.Write([]byte("Olá, mundo!"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Tamanho: %d bytes\n", tamanho2)

	f.Close()
}

func readFile() {
	// lê um arquivo
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(arquivo))

	// leitura usando buffer
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 4)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Print(string(buffer[:n]))
	}

}

func removeFile() {
	// remove um arquivo
	err := os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
