# Seção 1.6 - Buscando URLs de Modo Concorrente

## 📋 Descrição

Este programa demonstra um dos recursos mais poderosos da linguagem Go: **concorrência com goroutines e canais**. Ele busca múltiplas URLs simultaneamente, mostrando o tempo que cada requisição leva e o tamanho do conteúdo obtido.

## 🎯 Para Que Serve

O programa ilustra como executar tarefas de I/O (entrada/saída) de forma concorrente, permitindo que múltiplas requisições HTTP sejam feitas ao mesmo tempo, em vez de uma após a outra. Isso resulta em:

- **Redução significativa do tempo total de execução** quando comparado à execução sequencial
- **Melhor aproveitamento do tempo de espera** de rede
- **Demonstração prática de goroutines e canais**, conceitos fundamentais em Go

### Comparação: Sequencial vs Concorrente

**Modo Sequencial:**

- URL1 (2s) → URL2 (3s) → URL3 (1s) = **6 segundos total**

**Modo Concorrente (este programa):**

- URL1 (2s) ┐
- URL2 (3s) ├→ Todas ao mesmo tempo = **~3 segundos total** (tempo da mais lenta)
- URL3 (1s) ┘

## 🔧 Como Usar

### Executar o programa

```bash
# Compilar
go build buscando_url_de_modo_concorrente.go

# Executar com múltiplas URLs
./buscando_url_de_modo_concorrente https://golang.org https://google.com https://github.com
```

### Exemplo de saída

```
1.23s    5432 https://google.com
1.45s   12845 https://github.com
1.78s   23451 https://golang.org
1.78s elapsed
```

A saída mostra:

- Tempo de cada requisição
- Número de bytes baixados
- URL acessada
- Tempo total de execução

## 💡 Explicação do Código

### Estrutura Principal

```go
func main() {
    start := time.Now()           // Marca início
    ch := make(chan string)        // Cria canal de comunicação

    for _, url := range os.Args[1:] {
        go fetch(url, ch)          // Lança goroutines concorrentes
    }

    for range os.Args[1:] {
        fmt.Println(<-ch)          // Recebe resultados
    }

    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
```

### Conceitos Fundamentais

#### 1. **Goroutines** (`go fetch(url, ch)`)

- São threads leves gerenciadas pelo runtime do Go
- Permitem execução concorrente de funções
- Muito mais eficientes que threads do sistema operacional
- Uma goroutine é criada para cada URL

#### 2. **Canais** (`chan string`)

- Mecanismo de comunicação entre goroutines
- Garantem sincronização segura
- `ch <- valor`: envia valor para o canal
- `valor := <-ch`: recebe valor do canal
- **Bloqueiam** até que haja uma operação complementar (send/receive)

#### 3. **Sincronização**

```go
for range os.Args[1:] {
    fmt.Println(<-ch)  // Aguarda uma resposta por URL
}
```

- O loop aguarda exatamente uma resposta para cada URL lançada
- Garante que o programa não termine antes de todas as goroutines concluírem

### Função fetch()

```go
func fetch(url string, ch chan<- string) {
    // chan<- string: canal apenas para envio

    resp, err := http.Get(url)    // Faz requisição HTTP
    if err != nil {
        ch <- fmt.Sprintf("erro ao buscar %s: %v", url, err)
        return
    }

    nbytes, err := io.Copy(io.Discard, resp.Body)  // Conta bytes
    resp.Body.Close()                               // Libera recursos

    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
```

**Pontos-chave:**

- `io.Copy(io.Discard, resp.Body)`: descarta o conteúdo mas conta os bytes
- `resp.Body.Close()`: **essencial** para evitar vazamento de recursos
- `chan<- string`: indica que a função só pode **enviar** para o canal (type safety)

## 🌍 Onde Encontrar no Dia a Dia

### 1. **Microsserviços e APIs**

```
Cenário: Um serviço precisa buscar dados de múltiplas APIs externas
Solução: Fazer requisições concorrentes em vez de sequenciais
Exemplo: Dashboard que agrega dados de 5 serviços diferentes
```

### 2. **Web Scrapers e Crawlers**

```
Cenário: Extrair dados de centenas de páginas web
Solução: Processar múltiplas páginas simultaneamente
Exemplo: Monitoramento de preços em vários e-commerces
```

### 3. **Processamento de Dados**

```
Cenário: Processar múltiplos arquivos ou chunks de dados
Solução: Processar cada chunk em uma goroutine separada
Exemplo: Análise paralela de logs, conversão de imagens
```

### 4. **Servidores Web**

```
Cenário: Atender múltiplos clientes simultaneamente
Solução: Go cria automaticamente uma goroutine por conexão
Exemplo: Servidores HTTP, WebSocket, gRPC
```

### 5. **Integração com Múltiplos Bancos de Dados**

```
Cenário: Consultar múltiplos bancos ou shards ao mesmo tempo
Solução: Executar queries em paralelo e agregar resultados
Exemplo: Busca distribuída, aggregação de métricas
```

### 6. **Notificações e Alertas**

```
Cenário: Enviar notificações por múltiplos canais (email, SMS, push)
Solução: Enviar todas as notificações concorrentemente
Exemplo: Sistema de alertas que notifica por 3 canais diferentes
```

### 7. **Testes de Carga**

```
Cenário: Simular múltiplos usuários acessando um serviço
Solução: Criar goroutines para simular requisições concorrentes
Exemplo: Benchmark de performance de APIs
```

## 📚 Conceitos Avançados

### Fan-out / Fan-in Pattern

Este código implementa o padrão básico de **fan-out**:

- Uma goroutine principal distribui trabalho para múltiplas goroutines workers
- Cada worker processa sua tarefa independentemente
- Os resultados são coletados através do canal

### Controle de Concorrência

Para limitar o número de requisições simultâneas, pode-se usar:

- **Buffered channels**
- **Worker pools**
- **Semáforos** (via channels ou sync.Semaphore)

Exemplo de limitação:

```go
semaphore := make(chan struct{}, 5) // Máximo 5 requisições simultâneas

for _, url := range urls {
    semaphore <- struct{}{}  // Adquire
    go func(url string) {
        defer func() { <-semaphore }()  // Libera
        fetch(url, ch)
    }(url)
}
```

## ⚠️ Boas Práticas

1. **Sempre feche connections e response bodies**

   ```go
   defer resp.Body.Close()
   ```

2. **Trate erros adequadamente**

   - Não ignore erros de rede
   - Considere timeouts e retries

3. **Use context para cancelamento**

   ```go
   ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
   defer cancel()
   ```

4. **Limite concorrência em produção**
   - Evite criar milhares de goroutines simultaneamente
   - Use worker pools para controlar recursos

## 🔗 Recursos Adicionais

- [Effective Go - Concurrency](https://golang.org/doc/effective_go#concurrency)
- [Go Blog - Share Memory By Communicating](https://go.dev/blog/codelab-share)
- [Go by Example - Goroutines](https://gobyexample.com/goroutines)
- [Go by Example - Channels](https://gobyexample.com/channels)

## 📝 Exercícios Sugeridos

1. Modifique o programa para salvar o conteúdo das URLs em arquivos
2. Adicione um timeout para requisições que demoram muito
3. Implemente retry automático em caso de falha
4. Limite o número de requisições simultâneas a 5
5. Adicione suporte para POST requests com dados customizados
