# Seção 2.1 - Nomes

## Introdução

Os nomes são identificadores usados para variáveis, constantes, funções, tipos e outros elementos em programas Go. A escolha de nomes apropriados é fundamental para a legibilidade e manutenibilidade do código.

## Regras para Nomes

### Sintaxe Básica

Um nome em Go deve:

- Começar com uma **letra** (Unicode) ou **underscore** (`_`)
- Ser seguido por qualquer número de letras, dígitos ou underscores
- Ser sensível a maiúsculas/minúsculas (`nome`, `Nome` e `NOME` são diferentes)

```go
// Nomes válidos
x
i
nome
_temp
índice
π
すし  // caracteres Unicode são permitidos
```

```go
// Nomes inválidos
2nome      // não pode começar com dígito
nome-completo  // hífen não é permitido
nome completo  // espaços não são permitidos
```

### Palavras-chave Reservadas

Go possui **25 palavras-chave** que não podem ser usadas como identificadores:

```
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

### Identificadores Pré-declarados

Além das palavras-chave, Go possui identificadores pré-declarados que podem ser redefinidos (mas não é recomendado):

**Constantes:**

```
true    false    iota    nil
```

**Tipos:**

```
int     int8     int16    int32    int64
uint    uint8    uint16   uint32   uint64   uintptr
float32 float64  complex64 complex128
bool    byte     rune     string   error
```

**Funções:**

```
make    len      cap      new      append   copy
close   delete   complex  real     imag
panic   recover
```

## Visibilidade e Exportação

A visibilidade de um identificador é determinada pela **primeira letra** do seu nome:

### Exportados (Públicos)

Identificadores que começam com **letra maiúscula** são exportados e podem ser acessados de outros pacotes:

```go
package math

// Exportado - acessível de outros pacotes
func Sin(x float64) float64 {
    return sin(x)
}

// Exportado
const Pi = 3.14159
```

### Não Exportados (Privados)

Identificadores que começam com **letra minúscula** são privados ao pacote:

```go
package math

// Não exportado - apenas acessível dentro do pacote math
func sin(x float64) float64 {
    // implementação interna
}

// Não exportado
const epsilon = 1e-10
```

## Convenções de Nomenclatura

### CamelCase

Go usa **camelCase** (primeira letra minúscula) para nomes não exportados e **PascalCase** (primeira letra maiúscula) para nomes exportados:

```go
// Não exportado
var nomeCompleto string
var endereçoIP string

// Exportado
var NomeCompleto string
var EndereçoIP string
```

### Siglas e Acrônimos

Siglas e acrônimos devem ser **todas maiúsculas** ou **todas minúsculas**:

```go
// Bom
var userID int
var UserID int
var httpServer *Server
var HTTPServer *Server

// Ruim
var userId int      // evite misturar
var UserId int
var HttpServer *Server
```

### Nomes Curtos vs Longos

**Nomes curtos** são preferidos quando:

- O escopo é pequeno
- O significado é óbvio pelo contexto

```go
// Bom para loops curtos
for i := 0; i < 10; i++ {
    sum += i
}

// Bom para parâmetros óbvios
func Write(w io.Writer, buf []byte) (n int, err error)
```

**Nomes longos** são preferidos quando:

- O escopo é grande
- O nome precisa ser descritivo

```go
// Bom para variáveis de escopo amplo
var maxConnectionsPerClient int

// Bom para funções exportadas
func ParseHTTPResponseHeader(data []byte) (*Header, error)
```

### Convenções Específicas

#### Receptores de Métodos

Use nomes curtos (1-2 letras) consistentes:

```go
type Buffer struct { /* ... */ }

// Bom - 'b' é consistente em todos os métodos
func (b *Buffer) Write(p []byte) (n int, err error) { /* ... */ }
func (b *Buffer) Read(p []byte) (n int, err error) { /* ... */ }

// Evite
func (buf *Buffer) Write(p []byte) (n int, err error) { /* ... */ }
func (buffer *Buffer) Read(p []byte) (n int, err error) { /* ... */ }
```

#### Interfaces

Interfaces de um método geralmente terminam em **-er**:

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type Stringer interface {
    String() string
}
```

#### Getters

Go não usa o prefixo "Get" para métodos getter:

```go
// Bom
func (p *Person) Name() string { return p.name }

// Ruim
func (p *Person) GetName() string { return p.name }
```

Setters podem usar o prefixo "Set":

```go
func (p *Person) SetName(name string) { p.name = name }
```

## Boas Práticas

1. **Seja consistente**: Use o mesmo estilo em todo o código
2. **Seja descritivo**: Nomes devem revelar a intenção
3. **Evite abreviações obscuras**: Prefira clareza sobre brevidade
4. **Use o contexto**: Em um pacote `http`, `Server` é melhor que `HTTPServer`
5. **Evite redundância**: Em `user.UserName`, apenas `user.Name` é suficiente
6. **Evite underscores**: Use camelCase em vez de snake_case

## Exemplos Comparativos

### ❌ Ruim

```go
var n int                    // muito genérico
var number_of_users int      // use camelCase
var GetUserAge func() int    // não use Get para funções simples
const MAX_VALUE = 100        // use camelCase, não SCREAMING_CASE
```

### ✅ Bom

```go
var count int                // claro no contexto
var numberOfUsers int        // camelCase apropriado
var UserAge func() int       // exportado sem "Get"
const MaxValue = 100         // camelCase para constantes
```

## Resumo

- Nomes devem começar com letra ou underscore
- Primeira letra determina visibilidade (maiúscula = exportado)
- Use camelCase/PascalCase, não snake_case
- Prefira nomes descritivos, mas use contexto para decidir o comprimento
- Seja consistente com convenções da linguagem e da comunidade Go

## Recursos Adicionais

- [Effective Go - Names](https://golang.org/doc/effective_go.html#names)
- [Go Code Review Comments - Naming](https://github.com/golang/go/wiki/CodeReviewComments#naming)
- [Standard Library Style](https://golang.org/pkg/)
