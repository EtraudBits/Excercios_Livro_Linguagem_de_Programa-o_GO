// Define que este arquivo pertence ao pacote main, que é o pacote executável
package main

// Importa os pacotes necessários para o programa
import (
	"fmt"      // Para formatação e impressão de strings
	"io"       // Para operações de entrada e saída (I/O)
	"net/http" // Para fazer requisições HTTP
	"os"       // Para acessar argumentos da linha de comando
	"time"     // Para medir tempo de execução
)

// Função principal que será executada ao iniciar o programa
func main() {
	// Registra o momento de início da execução do programa
	start := time.Now()
	// Cria um canal (channel) para comunicação entre goroutines, que enviará strings
	ch := make (chan string)
	// Itera sobre os argumentos da linha de comando (os.Args[1:] pula o nome do programa)
	for _, url := range os.Args[1:] {
		// Inicia uma goroutine (execução concorrente) para buscar cada URL
		go fetch(url, ch)
	}
	// Aguarda receber uma resposta de cada goroutine lançada
	for range os.Args[1:] {
		// Imprime o resultado recebido do canal
		fmt.Println(<-ch)
	}
	// Imprime o tempo total decorrido desde o início, formatado com 2 casas decimais
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

// Função que busca uma URL e envia o resultado através do canal
// ch chan<- string indica que o canal é apenas para envio (send-only)
func fetch(url string, ch chan<- string) {
	// Registra o momento de início desta requisição específica
	start := time.Now()
	// Faz uma requisição HTTP GET para a URL fornecida
	resp, err := http.Get(url)
	// Verifica se houve erro na requisição
	if err != nil {
		// Envia mensagem de erro pelo canal e retorna
		ch <- fmt.Sprintf("erro ao buscar %s: %v", url, err)
		return
	}
	// Copia o corpo da resposta para io.Discard (descarta o conteúdo) e conta os bytes
	nbytes, err := io.Copy(io.Discard, resp.Body)
	// Fecha o corpo da resposta HTTP para liberar recursos
	resp.Body.Close()
	// Verifica se houve erro ao ler o corpo da resposta
	if err != nil {
		// Envia mensagem de erro pelo canal e retorna
		ch <- fmt.Sprintf("erro ao ler %s : %v", url, err)
		return
	}
	// Calcula o tempo decorrido desde o início desta requisição em segundos
	secs := time.Since(start).Seconds()
	// Envia pelo canal uma string formatada com: tempo, número de bytes e URL
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}