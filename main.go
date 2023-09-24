package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	appIntroduce()
	for {
		showMenu()
		command := readCommand()
		switch command {
		case 1:
			initMonitoring()
		case 2:
			showLogs()
		case 0:
			exitWithSucess()
		default:
			exitWithFailure()
		}
	}

}

func showLogs() {
	fmt.Println("Exibindo logs...")
	file, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro ao ler arquivo:", err)
	}

	fmt.Println(string(file))
}

func initMonitoring() {
	fmt.Println("Monitorando...")
	sites := readSitesFromFile()
	for i := 0; i < monitoramentos; i++ {
		// forma mais simples de percorrer um arrayq
		for index, element := range sites {
			fmt.Println("Posição:", index, "Site:", element)
			testSite(element)
		}
		time.Sleep(delay * time.Second)
	}

}

func testSite(site string) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro na requisição:", err)
		logRegister(site, false)
	}

	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		logRegister(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", response.StatusCode)
		logRegister(site, false)
	}
}

func exitWithFailure() {
	fmt.Println("Não conheço esse comando")
	os.Exit(-1)
}

func exitWithSucess() {
	fmt.Println("Saindo do programa...")
	os.Exit(0)
}

func appIntroduce() {
	name := "Rafael"
	version := 1.1
	fmt.Println("Olá sr.", name, "sua versão é", version)
}

func showMenu() {
	fmt.Println("Menu de opções:")
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	return command
}

func readSitesFromFile() []string {

	var sites []string
	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro ao abrir arquivo:", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		sites = append(sites, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ocorreu um erro ao ler arquivo:", err)
	}

	return sites
}

func logRegister(site string, online bool) {
	createdFile, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro ao criar arquivo:", err)
	}

	createdFile.WriteString(
		time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + fmt.Sprint(online) + "\n",
	)

	createdFile.Close()

}
