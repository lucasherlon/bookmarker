package main

import (
	"fmt"
	"os"
	"os/exec"
	"bufio"
	"math/rand"
    "time"

)

func abrirLink() {
	var lista []string

	file, err := os.Open("file.txt")
	if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        lista = append(lista, line)
    }
    if err := scanner.Err(); err != nil {
        fmt.Println("Error:", err)
    }
	tam := len(lista)
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(tam)
	link := lista[num]

	exec.Command("rundll32", "url.dll,FileProtocolHandler", link).Start()
	file.Close()

}

func cadastrarLink() {
	var link string

	file, err := os.OpenFile("file.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

	fmt.Println("Digite o link que deseja cadastrar: ")
	fmt.Scanln(&link)
	file.WriteString("\n"+link)
	file.Close()
}

func sair() {
	fmt.Println("Saindo...")
	time.Sleep(time.Second)
}

func main() {
		var resposta int
		fmt.Println("**Bookmarker**\n O que você quer fazer?")
		fmt.Println("1 - Abrir um link aleatório\n2 - Cadastrar um link\n3 - Sair")

		fmt.Scanln(&resposta)

		switch resposta {
			case 1:
				abrirLink()
				main()
			case 2:
				cadastrarLink()
				main()
			case 3:
				sair()
			default:
				fmt.Println("Entrada inválida! Tente novamente.")
				main()
		}

}
