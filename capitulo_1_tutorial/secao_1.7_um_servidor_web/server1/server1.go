// server1 é um servidor de "eco" simples que responde com a mesma mensagem que recebe.

// package main define que este é um programa executável (não uma biblioteca)
package main

// import declara os pacotes que serão usados neste programa
import (
	"fmt"      // pacote para formatação de texto e entrada/saída
	"log"      // pacote para registro de logs e erros
	"net/http" // pacote para funcionalidades de servidor e cliente HTTP
)

// main é a função principal que será executada quando o programa iniciar
func main() {
	// http.HandleFunc registra a função handler para processar requisições no caminho "/"
	// O padrão "/" corresponde a todas as URLs que começam com "/"
	// Sempre que uma requisição HTTP chegar, a função handler será chamada
	http.HandleFunc("/", handler)
	
	// log.Fatal executa a função fornecida e, se houver erro, registra o erro e encerra o programa
	// http.ListenAndServe inicia o servidor HTTP na porta 8000 do localhost
	// O primeiro parâmetro "localhost:8000" especifica o endereço e porta onde o servidor escutará
	// O segundo parâmetro nil indica que usaremos o multiplexador padrão do Go
	// Esta linha bloqueia a execução - o programa fica rodando e esperando requisições
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler é a função que processa cada requisição HTTP recebida pelo servidor
// Ela recebe dois parâmetros:
// w (http.ResponseWriter) - usado para escrever a resposta que será enviada ao cliente
// r (*http.Request) - contém todos os dados da requisição HTTP recebida
func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf escreve uma string formatada no ResponseWriter (w)
	// %q formata a string com aspas duplas e caracteres especiais escapados
	// r.URL.Path contém o caminho da URL requisitada (ex: "/", "/hello", "/api/users")
	// \n adiciona uma quebra de linha no final da resposta
	// O resultado é enviado como resposta HTTP ao cliente
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
