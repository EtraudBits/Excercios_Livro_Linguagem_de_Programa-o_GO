// Define que este arquivo pertence ao pacote main (programa executável)
package main

// Importa os pacotes necessários
import (
	"fmt"      // Para formatação e impressão de texto
	"io"       // Para operações de entrada/saída (leitura de dados)
	"net/http" // Para fazer requisições HTTP
	"os"       // Para acessar argumentos da linha de comando e stderr
)

func main() {
	// Percorre cada URL passada como argumento na linha de comando
	// os.Args[1:] pega todos os argumentos exceto o nome do programa (que está em os.Args[0])
	// O _ (blank identifier) ignora o índice, pois só precisamos da URL
	for _, url := range os.Args[1:] {
		// Faz uma requisição HTTP GET para a URL
		// Retorna a resposta (resp) e um possível erro (err)
		resp, err := http.Get(url)
		
		// Verifica se houve erro na requisição
		if err != nil {
			// Imprime o erro no stderr (saída de erros) em vez do stdout
			fmt.Fprintf(os.Stderr, "Erro ao buscar %s: %v\n", url, err)
			// Continue pula para a próxima iteração do loop
			continue
		}
		
		// Lê todo o conteúdo do corpo da resposta HTTP
		// Converte os dados recebidos em um slice de bytes
		body, err := io.ReadAll(resp.Body)
		
		// Fecha o corpo da resposta para liberar recursos
		// É importante sempre fechar após usar
		resp.Body.Close()
		
		// Verifica se houve erro ao ler o corpo da resposta
		if err != nil {
			// Imprime o erro no stderr
			fmt.Fprintf(os.Stderr, "Erro ao ler o corpo de %s: %v\n", url, err)
			// Pula para a próxima URL
			continue
		}
		
		// Imprime o conteúdo (body) convertendo de bytes para string
		// %s formata como string
		fmt.Printf("%s\n", body)
	}
} 