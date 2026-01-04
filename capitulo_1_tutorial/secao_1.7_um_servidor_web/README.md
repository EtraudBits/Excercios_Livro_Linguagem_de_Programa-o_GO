# Seção 1.7 - Um Servidor Web

Esta seção apresenta três exemplos progressivos de servidores web em Go, demonstrando desde um servidor básico até um servidor com recursos avançados de inspeção de requisições HTTP.

## Localização dos Programas

Os programas estão organizados em três diretórios:

- **server1/** - Servidor de eco básico
- **server2/** - Servidor com contador de requisições
- **server3/** - Servidor com inspeção completa de requisições

## Server 1: Servidor de Eco Básico

### Localização

`server1/server1.go`

### Para que serve

Este é o exemplo mais simples de um servidor web em Go. Ele responde a qualquer requisição HTTP retornando apenas o caminho (path) da URL requisitada.

### Funcionalidades

- Responde em `http://localhost:8000/`
- Ecoa o componente path da URL requisitada
- Demonstra o uso básico do pacote `net/http`

### Como executar

```bash
cd server1
go run server1.go
```

### Exemplos de uso

```bash
# Em outro terminal, faça requisições:
curl http://localhost:8000/
# Resposta: URL.Path = "/"

curl http://localhost:8000/hello
# Resposta: URL.Path = "/hello"

curl http://localhost:8000/api/users/123
# Resposta: URL.Path = "/api/users/123"
```

### Conceitos demonstrados

- `http.HandleFunc()` - Registra um handler para um padrão de URL
- `http.ListenAndServe()` - Inicia o servidor HTTP
- `http.ResponseWriter` - Interface para escrever a resposta HTTP
- `http.Request` - Estrutura que contém os dados da requisição

---

## Server 2: Servidor com Contador

### Localização

`server2/server2.go`

### Para que serve

Este servidor adiciona funcionalidade de contagem de requisições ao servidor básico. Ele mantém um contador global de quantas requisições foram feitas e fornece um endpoint dedicado para consultar esse contador.

### Funcionalidades

- Responde em `http://localhost:8000/` - ecoa o path da URL
- Responde em `http://localhost:8000/count` - mostra o número total de requisições
- Implementa sincronização com mutex para acesso thread-safe ao contador
- Incrementa o contador a cada requisição feita ao handler principal

### Como executar

```bash
cd server2
go run server2.go
```

### Exemplos de uso

```bash
# Faça algumas requisições ao servidor
curl http://localhost:8000/
# Resposta: URL.Path = "/"

curl http://localhost:8000/hello
# Resposta: URL.Path = "/hello"

curl http://localhost:8000/test
# Resposta: URL.Path = "/test"

# Consulte o contador de requisições
curl http://localhost:8000/count
# Resposta: Count 3
```

### Conceitos demonstrados

- **Múltiplos handlers** - Uso de mais de um `HandleFunc()` para diferentes rotas
- **Concorrência** - Uso de `sync.Mutex` para proteger variáveis compartilhadas
- **Variáveis globais** - Contador compartilhado entre requisições
- **Thread-safety** - Proteção contra race conditions usando `mu.Lock()` e `mu.Unlock()`

### Por que usar Mutex?

Go cria uma goroutine para cada requisição HTTP. Sem o mutex, múltiplas requisições simultâneas poderiam tentar modificar a variável `count` ao mesmo tempo, causando race conditions. O mutex garante que apenas uma goroutine por vez pode acessar o contador.

---

## Server 3: Servidor de Inspeção HTTP

### Localização

`server3/server3.go`

### Para que serve

Este é o servidor mais completo dos três exemplos. Ele funciona como uma ferramenta de depuração e inspeção de requisições HTTP, mostrando todos os detalhes de cada requisição recebida.

### Funcionalidades

- Responde em `http://localhost:8000/` - mostra informações detalhadas da requisição
- Responde em `http://localhost:8000/count` - mostra o número de requisições ao contador
- Exibe o método HTTP (GET, POST, etc.)
- Exibe a URL completa e o protocolo usado
- Exibe todos os cabeçalhos HTTP
- Exibe o host e o endereço remoto
- Exibe os dados de formulário (query parameters e POST data)

### Como executar

```bash
cd server3
go run server3.go
```

### Exemplos de uso

#### Requisição GET simples

```bash
curl http://localhost:8000/
```

Resposta:

```
GET / HTTP/1.1
Header["User-Agent"] = ["curl/7.x.x"]
Header["Accept"] = ["*/*"]
Host = "localhost:8000"
RemoteAddr = "127.0.0.1:xxxxx"
```

#### Requisição com query parameters

```bash
curl "http://localhost:8000/search?query=golang&page=1"
```

Resposta:

```
GET /search?query=golang&page=1 HTTP/1.1
Header["User-Agent"] = ["curl/7.x.x"]
Header["Accept"] = ["*/*"]
Host = "localhost:8000"
RemoteAddr = "127.0.0.1:xxxxx"
Form["query"] = ["golang"]
Form["page"] = ["1"]
```

#### Requisição POST com dados de formulário

```bash
curl -X POST -d "name=João&age=25" http://localhost:8000/submit
```

Resposta:

```
POST /submit HTTP/1.1
Header["User-Agent"] = ["curl/7.x.x"]
Header["Content-Type"] = ["application/x-www-form-urlencoded"]
Header["Content-Length"] = ["18"]
Host = "localhost:8000"
RemoteAddr = "127.0.0.1:xxxxx"
Form["name"] = ["João"]
Form["age"] = ["25"]
```

#### Consultar o contador

```bash
curl http://localhost:8000/count
```

Resposta:

```
Count 0
```

**Nota:** O contador no server3 não incrementa automaticamente porque o handler principal não o incrementa. Para que funcione como o server2, seria necessário adicionar o incremento do contador no handler.

### Conceitos demonstrados

- **Inspeção de requisições HTTP** - Acesso a todos os componentes da requisição
- **Cabeçalhos HTTP** - Iteração sobre `r.Header` usando range
- **Parsing de formulários** - Uso de `r.ParseForm()` para extrair dados
- **Método HTTP** - Acesso a `r.Method` (GET, POST, PUT, DELETE, etc.)
- **Protocolo** - Acesso a `r.Proto` (HTTP/1.1, HTTP/2, etc.)
- **Informações de conexão** - `r.Host` e `r.RemoteAddr`

### Quando usar este servidor

Este servidor é útil para:

- Depurar requisições HTTP de clientes
- Testar como diferentes ferramentas/bibliotecas enviam requisições
- Entender quais cabeçalhos e dados estão sendo enviados
- Aprender sobre o protocolo HTTP na prática

---

## Comparação dos Três Servidores

| Característica            | Server 1           | Server 2            | Server 3       |
| ------------------------- | ------------------ | ------------------- | -------------- |
| Complexidade              | Básico             | Intermediário       | Avançado       |
| Rotas                     | 1 (/)              | 2 (/, /count)       | 2 (/, /count)  |
| Contador de requisições   | ❌                 | ✅                  | ✅ (parcial)   |
| Sincronização             | ❌                 | ✅                  | ✅             |
| Exibe path                | ✅                 | ✅                  | ✅             |
| Exibe método HTTP         | ❌                 | ❌                  | ✅             |
| Exibe cabeçalhos          | ❌                 | ❌                  | ✅             |
| Exibe dados de formulário | ❌                 | ❌                  | ✅             |
| Uso principal             | Aprendizado básico | Contador de acessos | Depuração HTTP |

## Conceitos Importantes

### O Pacote net/http

Go possui um excelente suporte nativo para servidores HTTP através do pacote `net/http`. Este pacote fornece:

- Servidor HTTP completo
- Cliente HTTP
- Roteamento de requisições
- Suporte a HTTPS
- WebSockets (com extensões)

### Handlers

Um handler em Go é qualquer função que aceita `http.ResponseWriter` e `*http.Request`:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // Processar requisição e escrever resposta
}
```

### Goroutines Implícitas

O servidor HTTP do Go automaticamente cria uma goroutine para cada requisição recebida. Isso significa que seu servidor pode lidar com múltiplas requisições simultaneamente sem código adicional. Porém, você precisa proteger variáveis compartilhadas (como visto no contador).

## Dicas de Uso

1. **Para parar o servidor**: Pressione `Ctrl+C` no terminal
2. **Testando no navegador**: Você pode acessar `http://localhost:8000/` diretamente no navegador
3. **Mudando a porta**: Altere "localhost:8000" para "localhost:PORTA_DESEJADA" no código
4. **Logs**: Use `log.Printf()` para adicionar logs de depuração

## Próximos Passos

Após estudar estes exemplos, você pode:

- Adicionar mais rotas e handlers
- Implementar um servidor REST API completo
- Adicionar template HTML com `html/template`
- Implementar autenticação e sessões
- Servir arquivos estáticos com `http.FileServer`
- Adicionar middleware para logging, autenticação, etc.
