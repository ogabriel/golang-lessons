package main

import "fmt"
import "os"
import "net/http"

func main() {
	intro()

	for {
		exibeMenu()
		comando := lerComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}

func intro() {
	nome := "Gabriel"
	versao := 1.1

	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}
func lerComando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando é", comando)

	return comando
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	site := "https://random-status-code.herokuapp.com/"
	resposta, _ := http.Get(site)

	if resposta.StatusCode == 200 {
		fmt.Println("Site: ", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site: ", site, "não foi carregado com sucesso, status code: ", resposta.StatusCode)
	}
}
