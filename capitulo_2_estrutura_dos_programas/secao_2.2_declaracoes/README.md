# Seção 2.2 - Declarações

## Introdução

As declarações em Go definem nomes que são vinculados a entidades do programa, como variáveis, constantes, tipos e funções. Este documento explora os quatro principais tipos de declarações em Go.

## Tipos de Declarações

Go possui quatro tipos principais de declarações:

1. **var** - para declarar variáveis
2. **const** - para declarar constantes
3. **type** - para declarar tipos
4. **func** - para declarar funções

## 1. Declaração de Variáveis (var)

A palavra-chave `var` declara uma variável de um tipo específico e opcionalmente a inicializa.

### Sintaxe Básica

```go
var nome tipo = expressão
```

### Exemplos

```go
// Declaração com inicialização explícita
var idade int = 25
var nome string = "João"
var ativo bool = true

// Declaração sem inicialização (recebe o valor zero do tipo)
var contador int          // 0
var mensagem string       // ""
var disponivel bool       // false

// Declaração múltipla
var x, y int = 1, 2
var a, b, c = true, 2.5, "texto"

// Bloco de declarações
var (
    nome      string = "Maria"
    idade     int    = 30
    salario   float64 = 5000.50
    ativo     bool   = true
)
```

### Valor Zero (Zero Value)

Quando uma variável é declarada sem inicialização explícita, ela recebe o "valor zero" do seu tipo:

- Números: `0`
- Booleanos: `false`
- Strings: `""` (string vazia)
- Ponteiros, slices, maps, channels, funções, interfaces: `nil`

```go
var i int       // 0
var f float64   // 0.0
var s string    // ""
var p *int      // nil
```

### Declaração Curta (:=)

Dentro de funções, pode-se usar a declaração curta `:=` que infere o tipo automaticamente:

```go
func exemplo() {
    nome := "Carlos"           // string
    idade := 28                // int
    salario := 3500.75         // float64
    ativo := true              // bool
    x, y := 10, 20            // múltiplas variáveis
}
```

**Importante**: A declaração curta `:=` só pode ser usada dentro de funções, não no nível de pacote.

## 2. Declaração de Constantes (const)

Constantes são valores que não podem ser modificados após sua declaração. São avaliadas em tempo de compilação.

### Sintaxe Básica

```go
const nome tipo = expressão
```

### Exemplos

```go
// Constantes simples
const Pi float64 = 3.14159265359
const Titulo string = "Sistema de Gestão"
const MaxUsuarios int = 100

// Constantes sem tipo (untyped)
const Verdade = true
const Numero = 42
const Texto = "Olá"

// Múltiplas constantes
const x, y = 1, 2
const a, b, c = true, 2.5, "três"

// Bloco de constantes
const (
    Domingo = 0
    Segunda = 1
    Terca   = 2
    Quarta  = 3
    Quinta  = 4
    Sexta   = 5
    Sabado  = 6
)
```

### iota - Gerador de Constantes

O identificador `iota` é usado para criar constantes enumeradas:

```go
const (
    Segunda = iota  // 0
    Terca           // 1
    Quarta          // 2
    Quinta          // 3
    Sexta           // 4
    Sabado          // 5
    Domingo         // 6
)

// Com valores customizados
const (
    _  = iota             // ignora o primeiro valor (0)
    KB = 1 << (10 * iota) // 1 << (10*1) = 1024
    MB                     // 1 << (10*2) = 1048576
    GB                     // 1 << (10*3) = 1073741824
    TB                     // 1 << (10*4)
)

// Resetando iota
const (
    A = iota  // 0
    B         // 1
)
const (
    C = iota  // 0 (iota reseta)
    D         // 1
)
```

## 3. Declaração de Tipos (type)

A palavra-chave `type` cria um novo tipo nomeado ou um alias para um tipo existente.

### Sintaxe Básica

```go
type NovoTipo TipoBase
```

### Exemplos

```go
// Tipo baseado em tipo primitivo
type Celsius float64
type Fahrenheit float64
type Idade int
type Nome string

// Tipo struct
type Pessoa struct {
    Nome      string
    Idade     int
    Email     string
    Ativo     bool
}

// Tipo slice
type ListaInteiros []int
type ListaPessoas []Pessoa

// Tipo map
type Catalogo map[string]float64

// Tipo função
type Operacao func(int, int) int
type Manipulador func(string) error

// Tipo interface
type Escritor interface {
    Escrever([]byte) (int, error)
}

type Geometria interface {
    Area() float64
    Perimetro() float64
}
```

### Alias de Tipos (Go 1.9+)

```go
// Alias - cria um nome alternativo para um tipo existente
type MeuInt = int

// É diferente de criar um novo tipo
type NovoInt int
```

## 4. Declaração de Funções (func)

Funções são declaradas com a palavra-chave `func` seguida do nome, parâmetros e tipo de retorno.

### Sintaxe Básica

```go
func nomeFuncao(parametros) tipoRetorno {
    // corpo da função
}
```

### Exemplos

```go
// Função sem parâmetros e sem retorno
func saudar() {
    fmt.Println("Olá!")
}

// Função com parâmetros
func soma(a int, b int) int {
    return a + b
}

// Parâmetros do mesmo tipo (forma curta)
func multiplicar(a, b int) int {
    return a * b
}

// Múltiplos valores de retorno
func dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("divisão por zero")
    }
    return a / b, nil
}

// Retorno nomeado
func retangulo(largura, altura float64) (area, perimetro float64) {
    area = largura * altura
    perimetro = 2 * (largura + altura)
    return // retorno implícito
}

// Função variádica
func somar(numeros ...int) int {
    total := 0
    for _, n := range numeros {
        total += n
    }
    return total
}

// Função como tipo
type Operacao func(int, int) int

func aplicar(op Operacao, a, b int) int {
    return op(a, b)
}

// Métodos (funções com receptor)
type Retangulo struct {
    Largura, Altura float64
}

func (r Retangulo) Area() float64 {
    return r.Largura * r.Altura
}
```

## Escopo das Declarações

O escopo determina onde um nome declarado pode ser usado no código.

### Níveis de Escopo

1. **Escopo de Universo**: nomes predeclarados (int, bool, true, false, etc.)
2. **Escopo de Pacote**: declarações no nível superior do pacote
3. **Escopo de Arquivo**: importações (apenas no arquivo atual)
4. **Escopo de Função**: parâmetros e variáveis locais
5. **Escopo de Bloco**: variáveis dentro de `{}`, `if`, `for`, etc.

### Exemplos de Escopo

```go
package main

import "fmt"

// Escopo de pacote
var global = "variável global"

const ConstantePacote = 100

type Usuario struct {
    Nome string
}

func exemplo() {
    // Escopo de função
    local := "variável local"

    for i := 0; i < 10; i++ {
        // Escopo de bloco (loop)
        temp := i * 2
        fmt.Println(temp)
    }
    // temp não existe aqui

    if x := 5; x > 0 {
        // Escopo de bloco (if)
        fmt.Println(x)
    }
    // x não existe aqui
}
```

### Visibilidade (Exportação)

Em Go, a visibilidade de um nome é determinada pela primeira letra:

- **Maiúscula**: exportado (visível fora do pacote)
- **Minúscula**: não exportado (visível apenas no pacote)

```go
package matematica

// Exportado - pode ser usado por outros pacotes
var Pi = 3.14159
const MaxValor = 100

type Circulo struct {
    Raio float64  // campo exportado
}

func Somar(a, b int) int {  // função exportada
    return a + b
}

// Não exportado - uso interno apenas
var versao = "1.0.0"
const minValor = 0

type configuracao struct {
    timeout int  // campo não exportado
}

func validar(valor int) bool {  // função não exportada
    return valor >= minValor
}
```

## Ordem de Declaração

Em Go, a ordem das declarações no nível de pacote não importa na maioria dos casos, exceto para:

1. Inicialização de variáveis que dependem umas das outras
2. A função `init()` é executada após todas as variáveis serem inicializadas

```go
package main

import "fmt"

// Declarações no nível de pacote
var b = a * 2  // OK: pode referenciar 'a' antes de sua declaração
var a = 10

func init() {
    fmt.Println("Inicialização do pacote")
    // Executado automaticamente antes de main()
}

func main() {
    fmt.Println(a, b) // 10 20
}
```

## Boas Práticas

1. **Declare variáveis próximas ao seu uso**: Torna o código mais legível
2. **Use declaração curta (:=) quando possível**: Mais conciso dentro de funções
3. **Declare constantes para valores fixos**: Evita "números mágicos"
4. **Use nomes descritivos**: Facilita a compreensão
5. **Agrupe declarações relacionadas**: Use blocos `var ()` ou `const ()`
6. **Prefira tipos nomeados para clareza**: `type Celsius float64` é mais claro que `float64`
7. **Use iota para enumerações**: Simplifica constantes sequenciais

## Exemplos Práticos

### Exemplo 1: Declarações Variadas

```go
package main

import "fmt"

// Constantes do programa
const (
    AppName    = "Sistema de Gestão"
    AppVersion = "1.0.0"
    MaxUsers   = 1000
)

// Tipos customizados
type User struct {
    ID       int
    Username string
    Email    string
    Active   bool
}

type UserList []User

// Variáveis globais
var (
    totalUsers   int
    activeUsers  int
    databaseURL  string
)

func main() {
    // Declaração local
    admin := User{
        ID:       1,
        Username: "admin",
        Email:    "admin@exemplo.com",
        Active:   true,
    }

    fmt.Printf("%s v%s\n", AppName, AppVersion)
    fmt.Printf("Usuário: %+v\n", admin)
}
```

### Exemplo 2: Função com Declarações

```go
package main

import (
    "fmt"
    "errors"
)

// Tipo para operações
type Operation func(int, int) int

// Constantes de operação
const (
    OpSum = iota
    OpSub
    OpMul
    OpDiv
)

func calculate(op Operation, a, b int) int {
    return op(a, b)
}

func getOperation(opType int) (Operation, error) {
    switch opType {
    case OpSum:
        return func(a, b int) int { return a + b }, nil
    case OpSub:
        return func(a, b int) int { return a - b }, nil
    case OpMul:
        return func(a, b int) int { return a * b }, nil
    case OpDiv:
        return func(a, b int) int { return a / b }, nil
    default:
        return nil, errors.New("operação inválida")
    }
}

func main() {
    op, err := getOperation(OpSum)
    if err != nil {
        fmt.Println("Erro:", err)
        return
    }

    result := calculate(op, 10, 5)
    fmt.Println("Resultado:", result) // 15
}
```

## Resumo

- **var**: Declara variáveis com tipo explícito ou inferido
- **const**: Declara constantes imutáveis em tempo de compilação
- **type**: Define novos tipos ou aliases
- **func**: Declara funções e métodos
- O escopo determina a visibilidade dos nomes
- A primeira letra determina se um nome é exportado (maiúscula) ou não (minúscula)
- Use declarações apropriadas para tornar o código claro e manutenível

## Recursos Adicionais

- [The Go Programming Language Specification - Declarations](https://golang.org/ref/spec#Declarations_and_scope)
- [Effective Go - Names](https://golang.org/doc/effective_go#names)
- [Go by Example - Variables](https://gobyexample.com/variables)
- [Go by Example - Constants](https://gobyexample.com/constants)
