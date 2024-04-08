package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

//Aula 2 - Arquivos e Diretórios
//Aula 3 - Tipos Básicos De Variáveis, etc

//Aula 7 - Requisições Web

type URL string
type Healthchek struct {
	Url          URL
	StatusCode   int
	ResponseBody []byte
}

func (obj *Healthchek) FormatResult() string {
	return fmt.Sprintf("Host: %s - StatusCode: %d \n", string(obj.Url), obj.StatusCode)
}

func main() {
	arquivo := "healthcheck"
	f, err := os.OpenFile(arquivo, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		panic(err.Error())
	}

	defer f.Close()

	//Verifica se o arquivo existe
	/*if _, err := os.Stat("hosts.txt"); os.IsNotExist(err) {
		fmt.Println("O Arquivo não existe!")
		return
	}*/

	//Abrir o arquivo com os host a serem verificados
	hosts, err := os.Open("hosts.txt")

	if err != nil {
		panic(err)
	}

	defer hosts.Close()

	scanner := bufio.NewScanner(hosts)

	for scanner.Scan() {
		url := scanner.Text()
		res, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		responseBytes, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		apiHealthCheck := &Healthchek{
			Url:          URL(url),
			StatusCode:   res.StatusCode,
			ResponseBody: responseBytes,
		}
		_, err = f.WriteString(apiHealthCheck.FormatResult())
		if err != nil {
			panic(err)
		}

		fmt.Printf("HOST: %s Status:%d \n", URL(url), res.StatusCode)

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Falha Ao Ler O Arquivo", err)
	}

}
