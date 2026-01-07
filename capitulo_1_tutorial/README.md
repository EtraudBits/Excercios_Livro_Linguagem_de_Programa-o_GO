# Capítulo 1: Tutorial

## Visão Geral

Este capítulo apresenta uma introdução prática à linguagem Go através de exemplos completos e funcionais. O objetivo é demonstrar os conceitos fundamentais da linguagem através de programas reais, permitindo que você comece a escrever código Go útil imediatamente.

## Objetivos de Aprendizado

Ao completar este capítulo, você será capaz de:

- Compreender a estrutura básica de um programa Go
- Trabalhar com entrada e saída de dados
- Utilizar pacotes da biblioteca padrão
- Implementar concorrência básica com goroutines
- Criar servidores web simples
- Manipular argumentos de linha de comando

## Tópicos Abordados

### 1.1 Olá, Mundo

O exemplo clássico introdutório que demonstra a estrutura mínima de um programa Go, incluindo:

- Declaração de pacotes
- Importação de bibliotecas
- Função `main` como ponto de entrada
- Formatação básica de saída

### 1.2 Argumentos de Linha de Comando

Exploração de como programas Go podem receber e processar argumentos fornecidos pelo usuário ao executar o programa:

- Acesso aos argumentos via `os.Args`
- Iteração sobre slices
- Estruturas de repetição básicas

### 1.3 Encontrando Linhas Duplicadas

Programa prático que demonstra:

- Leitura de arquivos
- Manipulação de mapas (maps)
- Contagem e processamento de dados
- Entrada a partir de múltiplas fontes

### 1.4 GIFs Animados

Introdução à geração de conteúdo gráfico:

- Pacote `image` e suas variações
- Manipulação de cores e paletas
- Criação de animações programáticas
- Geração de arquivos GIF

**Exemplo nesta seção**: [secao_1.4_gif_animados](secao_1.4_gif_animados/)

### 1.5 Buscando um URL

Demonstração de operações de rede básicas:

- Requisições HTTP usando o pacote `net/http`
- Leitura de respostas da web
- Tratamento básico de erros
- Trabalho com streams de dados

**Exemplo nesta seção**: [secao_1.5_buscando_um_url](secao_1.5_buscando_um_url/)

### 1.6 Buscando URLs de Modo Concorrente

Introdução à programação concorrente em Go:

- Goroutines para execução paralela
- Channels para comunicação entre goroutines
- Medição de tempo de execução
- Comparação entre execução sequencial e concorrente

**Exemplo nesta seção**: [secao_1.6_buscando_url_de_modo_concorrente](secao_1.6_buscando_url_de_modo_concorrente/)

### 1.7 Um Servidor Web

Criação de servidores HTTP simples:

- Handler functions
- Roteamento de requisições
- Processamento de parâmetros de URL
- Geração dinâmica de conteúdo HTML

**Exemplos nesta seção**:

- [server1](secao_1.7_um_servidor_web/server1/)
- [server2](secao_1.7_um_servidor_web/server2/)
- [server3](secao_1.7_um_servidor_web/server3/)

## Conceitos-Chave Introduzidos

### Sintaxe Básica

- Declaração de variáveis e constantes
- Tipos de dados fundamentais
- Estruturas de controle (if, for, switch)
- Funções e seus retornos

### Pacotes

- Organização de código em pacotes
- Importação e uso de pacotes da biblioteca padrão
- Convenções de nomenclatura

### Concorrência

- Goroutines como unidade básica de concorrência
- Channels para comunicação segura entre goroutines
- Modelo de concorrência CSP (Communicating Sequential Processes)

### Tratamento de Erros

- Retorno explícito de erros
- Verificação e tratamento de erros
- Padrões idiomáticos para lidar com falhas

### I/O e Rede

- Operações de entrada e saída
- Requisições HTTP
- Leitura e escrita de arquivos
- Streams de dados

## Ferramentas e Comandos

### Comandos Go Essenciais

```bash
# Executar um programa Go
go run arquivo.go

# Compilar um programa
go build arquivo.go

# Formatar código automaticamente
go fmt arquivo.go

# Obter documentação de um pacote
go doc pacote
```

## Estrutura Típica de um Programa Go

```go
package main

import (
    "fmt"
    // outros imports
)

func main() {
    // código principal
    fmt.Println("Olá, Go!")
}
```

## Dicas para Estudo

1. **Execute os exemplos**: Não apenas leia o código, execute-o e experimente modificações
2. **Explore a documentação**: Use `go doc` para aprender sobre pacotes e funções
3. **Experimente variações**: Modifique os exemplos para testar diferentes abordagens
4. **Observe os erros**: Go tem mensagens de erro claras que ajudam no aprendizado
5. **Pratique concorrência**: Entender goroutines e channels desde o início é fundamental

## Recursos Adicionais

- [Documentação Oficial do Go](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go Playground](https://play.golang.org/) - para experimentar código online

## Próximos Passos

Após completar este capítulo tutorial, você estará preparado para:

- Explorar a estrutura detalhada dos programas (Capítulo 2)
- Aprofundar-se nos tipos de dados básicos
- Compreender funções em maior profundidade
- Trabalhar com estruturas de dados mais complexas

---

**Nota**: Esta documentação serve como guia de estudo para o Capítulo 1 do livro "A Linguagem de Programação Go". Para uma compreensão completa, recomenda-se a leitura do livro original e a prática com os exercícios fornecidos.
