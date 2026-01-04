// handler pode informar os cabeçalhos e os dados de formulário que recebe - tornando o servidor útil para inspecionar e depurar requisições HTTP.

// package main define que este é um programa executável
package main

// import declara os pacotes necessários para o programa
import (
	"fmt"      // pacote para formatação de texto e entrada/saída
	"log"      // pacote para registro de logs e erros
	"net/http" // pacote para funcionalidades de servidor e cliente HTTP
	"sync"     // pacote para primitivas de sincronização (mutex, waitgroups, etc.)
)

// var declara variáveis no escopo global do pacote
// mu é um mutex (mutual exclusion) usado para proteger o acesso concorrente
var mu sync.Mutex

// count armazena o número de requisições ao endpoint /count
// Nota: este contador NÃO é incrementado no handler principal deste servidor
var count int

// main é a função principal que inicializa e configura o servidor
func main() {
	// http.HandleFunc registra o handler para o caminho raiz "/"
	// Este handler mostrará todas as informações detalhadas da requisição HTTP
	http.HandleFunc("/", handler)
	
	// http.HandleFunc registra o handler para o caminho "/count"
	// Este handler mostrará o valor do contador (sempre 0 neste exemplo)
	http.HandleFunc("/count", counter)
	
	// log.Fatal inicia o servidor HTTP no endereço localhost:8000
	// Se houver erro ao iniciar o servidor, o programa termina e mostra o erro
	// Esta linha bloqueia a execução mantendo o servidor rodando indefinidamente
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler é a função que processa requisições HTTP e exibe informações detalhadas
// Esta função é útil para depuração, mostrando todos os componentes da requisição
// Parâmetros:
//   w (http.ResponseWriter) - usado para escrever a resposta HTTP ao cliente
//   r (*http.Request) - ponteiro para a estrutura que contém todos os dados da requisição
func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf escreve a primeira linha mostrando método HTTP, URL completa e protocolo
	// %s formata strings sem aspas
	// r.Method contém o método HTTP (GET, POST, PUT, DELETE, etc.)
	// r.URL contém a URL completa requisitada (caminho + query parameters)
	// r.Proto contém a versão do protocolo HTTP usado (ex: HTTP/1.1, HTTP/2)
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	
	// for range itera sobre todos os cabeçalhos HTTP da requisição
	// r.Header é um map onde a chave (k) é o nome do cabeçalho
	// e o valor (v) é um slice de strings (pois cabeçalhos podem ter múltiplos valores)
	// Exemplos de cabeçalhos: User-Agent, Accept, Content-Type, Authorization, etc.
	for k, v := range r.Header {
		// Imprime cada cabeçalho no formato: Header["nome"] = ["valor1", "valor2", ...]
		// %q formata strings com aspas duplas e escapa caracteres especiais
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	
	// fmt.Fprintf imprime o host para onde a requisição foi enviada
	// r.Host contém o host da requisição (ex: "localhost:8000", "example.com")
	// Útil quando o servidor responde a múltiplos hosts/domínios
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	
	// fmt.Fprintf imprime o endereço remoto do cliente que fez a requisição
	// r.RemoteAddr contém o IP e porta do cliente (ex: "127.0.0.1:54321")
	// Útil para logging, controle de acesso e análise de tráfego
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	
	// r.ParseForm() analisa e extrai os dados de formulário da requisição
	// Processa tanto query parameters da URL (ex: ?name=João&age=25)
	// quanto dados do corpo da requisição (POST/PUT com application/x-www-form-urlencoded)
	// Retorna um erro se houver problema ao fazer o parsing
	if err := r.ParseForm(); err != nil {
		// Se houver erro ao fazer parsing do formulário, registra no log
		// log.Print escreve a mensagem de erro no log padrão sem encerrar o programa
		// A requisição continua sendo processada mesmo com erro
		log.Print(err)
	}
	
	// for range itera sobre todos os dados de formulário extraídos
	// r.Form é um map onde a chave (k) é o nome do campo
	// e o valor (v) é um slice de strings (campos podem ter múltiplos valores)
	// Contém tanto query parameters da URL quanto dados do corpo da requisição
	for k, v := range r.Form {
		// Imprime cada campo do formulário no formato: Form["campo"] = ["valor1", "valor2", ...]
		// %q formata com aspas duplas permitindo ver espaços e caracteres especiais
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

// counter é a função que processa requisições HTTP na rota "/count"
// Responde com o valor atual do contador (sempre 0 porque não é incrementado)
// Parâmetros:
//   w (http.ResponseWriter) - usado para escrever a resposta HTTP ao cliente
//   r (*http.Request) - contém os dados da requisição recebida
func counter(w http.ResponseWriter, r *http.Request) {
	// mu.Lock() adquire o bloqueio do mutex antes de acessar a variável count
	// Protege contra race conditions caso múltiplas requisições acessem simultaneamente
	mu.Lock()
	
	// fmt.Fprintf escreve o valor do contador na resposta HTTP
	// %d formata o número inteiro em notação decimal
	// Neste servidor, count sempre será 0 porque não é incrementado em nenhum lugar
	fmt.Fprintf(w, "Count %d\n", count)
	
	// mu.Unlock() libera o bloqueio do mutex após acessar count
	// Permite que outras goroutines possam acessar a variável
	mu.Unlock()
}