package main

import (  
    "fmt"
    "os"
    "net/http"
)

func main() {

	exibeIntroducao()
	for {

		exibirMenu()
		comando := leComando()
		switch comando {
		case 1: 
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0) // libera espaço na memória???
		default:
			fmt.Println("Comando não encontrado")
			os.Exit(-1) // sai indicando que ocorreu um erro inesperado
		}
	}
}

func exibeIntroducao() {
    nome := "Mateus Brandeurski Ramos"
    versao := 0.2
    fmt.Println("Olá, sr(a).", nome)
    fmt.Println("Este programa está na versão", versao)
}

func exibirMenu () {
    
    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")
}

func leComando() int {
    var comandoLido int
    fmt.Scan(&comandoLido)
    fmt.Println("O Comando escolhido foi", comandoLido)

    return comandoLido
}

func iniciarMonitoramento () {

    fmt.Println("Monitorando...")

	sites := []string{"https://alura.com.br","https://web-production-c42c.up.railway.app/", "https://api.cryptoscamdb.org/"}

    for _, site := range(sites) {	

		resp, _ := http.Get(site)

		if resp.StatusCode == 200 {
			fmt.Println("Site:", site, "Carregado com Sucesso! Status Code:", resp.StatusCode)
		} else if resp.StatusCode == 404 {
			fmt.Println("Site:", site, "não encontrado ou fora do ar")
		} else {
			fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		} 
	}  
}

// func exibiNomes() {

// 	nomes := []string{"Douglas", "Mateus", "Enzo"}
// 	fmt.Println(nomes)
// }