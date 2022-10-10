package main //pacote principal

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitortamentos = 3
const delay = 5

func main() {
	exibeIntroducao()
	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
			break
		case 2:
			fmt.Println("Exibindo logs..")
			imprimeLogs()
			break
		case 0:
			fmt.Println("Saindo do programa..")
			os.Exit(0)
			break
		default:
			fmt.Println("Não conheço este comando!")
			os.Exit(-1)
		}
	}

}

func devolveNomeEIdade() (string, int) {
	nome := "Guilherme"
	idade := 26
	return nome, idade
}

func exibeIntroducao() {
	// nome := "Guilherme"
	versao := 1.1

	fmt.Println("Olá") //virgula para concatenar
	fmt.Println("Este programa está na versão", versao)
	//fmt.Println("O tipo da variavel nome é", reflect.TypeOf(versao))
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	//fmt.Println("O endereco da minha variavel comendo é:", &comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	fmt.Println(" ")

	return comandoLido
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando..")
	// sites := []string{
	// 	"https://random-status-code.herokuapp.com",
	// 	"https://google.com.br",
	// 	"https://caelum.com.br"}

	//fmt.Println(sites)
	sites := leSitesDoArquivo()
	fmt.Println("---------------------------------------------------------------------")

	for i := 0; i < monitortamentos; i++ {
		for i, site := range sites {
			fmt.Println("testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println(" ")
		fmt.Println("---------------------------------------------------------------------")

	}

	fmt.Println(" ")

}

func testaSite(site string) {

	//site := "https://random-status-code.herokuapp.com"
	resp, err := http.Get(site)
	//fmt.Println(resp)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!  Status Code:", resp.StatusCode)
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas.  Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
	fmt.Println("---------------------------------------------------------------------")
}

// slice
func exibeNomes() {
	nomes := []string{
		"Douglas",
		"Daniel",
		"Guilherme",
		"Bernardo",
	}
	fmt.Println(nomes)
}

func leSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")
	//arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n') //aspas simples para declarar um byte
		linha = strings.TrimSpace(linha)      //tirar o pular linha

		sites = append(sites, linha) //add cada linha"site" dentro do slice

		if err == io.EOF {
			break
		}

	}
	arquivo.Close() //fechar arquivo apos o seu uso

	return sites
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) // abre e le esse arquivo para min OU crie esse arquivo para mim

	if err != nil {
		fmt.Println(err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "_online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}
