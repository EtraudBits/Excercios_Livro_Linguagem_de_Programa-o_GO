// server2 é um servidor mínimo de "eco" e contador que responde com a mesma mensagem que recebe

// package main define que este é um programa executável
package main

// import declara os pacotes necessários para o programa
import (
	"fmt"      // pacote para formatação de texto e entrada/saída
	"log"      // pacote para registro de logs e erros
	"net/http" // pacote para funcionalidades de servidor HTTP
	"sync"     // pacote para primitivas de sincronização (mutex, waitgroups, etc.)
)

// var declara variáveis no escopo global do pacote
// mu é um mutex (mutual exclusion) usado para proteger o acesso concorrente à variável count
// Mutex garante que apenas uma goroutine por vez possa acessar a variável protegida
var mu sync.Mutex

// count é a variável global que armazena o número total de requisições recebidas
// Como é global, ela persiste entre diferentes chamadas de handler
var count int

// main é a função principal que inicializa e configura o servidor
func main() {
	// http.HandleFunc registra o handler para o caminho raiz "/"
	// Todas as requisições que chegarem em "/" serão processadas pela função handler
	http.HandleFunc("/", handler)
	
	// http.HandleFunc registra um segundo handler para o caminho "/count"
	// Requisições em "/count" serão processadas pela função counter
	// Isso demonstra como ter múltiplas rotas no mesmo servidor
	http.HandleFunc("/count", counter)
	
	// log.Fatal inicia o servidor HTTP e registra qualquer erro fatal
	// http.ListenAndServe bloqueia a execução e mantém o servidor rodando
	// O servidor escuta no endereço "localhost:8000"
	// O parâmetro nil significa que usaremos o multiplexador HTTP padrão do Go
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler é a função que processa requisições HTTP na rota raiz "/"
// Ela incrementa o contador e ecoa o path da URL requisitada
// Parâmetros:
//   w (http.ResponseWriter) - usado para escrever a resposta HTTP
//   r (*http.Request) - contém os dados da requisição recebida
func handler(w http.ResponseWriter, r *http.Request) {
	// mu.Lock() adquire o bloqueio do mutex antes de acessar a variável count
	// Isso impede que outras goroutines acessem count ao mesmo tempo (race condition)
	// É crucial porque o servidor cria uma goroutine para cada requisição
	mu.Lock()
	
	// count++ incrementa o contador de requisições em 1
	// O operador ++ é equivalente a: count = count + 1
	// Esta operação precisa estar protegida pelo mutex
	count++
	
	// mu.Unlock() libera o bloqueio do mutex
	// Outras goroutines podem agora acessar count
	// É importante sempre liberar o mutex após usá-lo para evitar deadlocks
	mu.Unlock()
	
	// fmt.Fprintf escreve a resposta formatada no ResponseWriter
	// %q formata a string com aspas duplas
	// r.URL.Path contém o caminho da URL (ex: "/", "/hello", etc.)
	// A resposta é enviada de volta ao cliente HTTP
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter é a função que processa requisições HTTP na rota "/count"
// Ela responde com o número total de requisições que o servidor recebeu
// Parâmetros:
//   w (http.ResponseWriter) - usado para escrever a resposta HTTP
//   r (*http.Request) - contém os dados da requisição recebida
func counter(w http.ResponseWriter, r *http.Request) {
	// mu.Lock() adquire o bloqueio do mutex antes de ler a variável count
	// Mesmo leituras precisam ser protegidas para garantir consistência dos dados
	// Isso previne ler um valor enquanto ele está sendo modificado
	mu.Lock()
	
	// fmt.Fprintf lê o valor de count e escreve na resposta HTTP
	// %d formata o número inteiro em notação decimal
	// A resposta mostra quantas requisições foram feitas ao handler principal
	fmt.Fprintf(w, "Count %d\n", count)
	
	// mu.Unlock() libera o bloqueio do mutex após a leitura
	// Outras goroutines podem agora acessar count novamente
	mu.Unlock()	
}