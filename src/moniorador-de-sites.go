package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoramentos = 5
const delay = 1 

func main() {

	exibeIntroducao()
	for {
		exibirMenu()
		comando := leComando()
		switch comando {
		case 1:
			iniciarMonitoramento(leSitesDoArquivo())
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

func exibirMenu() {
	fmt.Print("\n")
    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")
	fmt.Print("\n")
}

func leComando() int {
    var comandoLido int
	fmt.Print("Digite o comando: ")
    fmt.Scan(&comandoLido)
    fmt.Println("O Comando escolhido foi", comandoLido)

    return comandoLido
}

func iniciarMonitoramento(sites []string) {

    fmt.Println("Monitorando...")

	// sites := 
	cont := 0
	
	for i:=0 ; i < monitoramentos ; i++ {
		cont++
		fmt.Println("Número da tentiva:",cont)
		for _, site := range(sites) {
			testaSite(site)
		}
		time.Sleep(delay * time.Second)  
	}
}

func testaSite(site string) {

	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "Carregado com Sucesso! Status Code:", resp.StatusCode)
	} else if resp.StatusCode == 404 {
		fmt.Println("Site:", site, "não encontrado ou fora do ar")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	} 

}

func leSitesDoArquivo() []string {
	var sites []string

	arquivo, _ := os.Open("sites-monitorados.txt")
	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
		
	}
	arquivo.Close() //liberar o arquivo para outros programas do SO poderem usar.
	return sites
}	

