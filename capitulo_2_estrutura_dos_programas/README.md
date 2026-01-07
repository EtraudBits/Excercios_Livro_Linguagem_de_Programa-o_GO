# Capítulo 2 - Estrutura dos Programas

Este capítulo aborda os conceitos fundamentais da estrutura de programas em Go, explorando como organizar e estruturar código de maneira eficaz.

## Tópicos Principais

### 2.1 Nomes

- Regras para nomenclatura de identificadores em Go
- Convenções de nomenclatura (camelCase, PascalCase)
- Palavras-chave reservadas
- Visibilidade de identificadores (exportados vs não exportados)
- Boas práticas para nomeação

### 2.2 Declarações

- Declaração de variáveis
- Declaração de constantes
- Declaração de tipos
- Declaração de funções
- Escopo das declarações

### 2.3 Variáveis

- Declaração com `var`
- Declaração curta com `:=`
- Valores zero (zero values)
- Ponteiros
- A função `new`
- Tempo de vida das variáveis

### 2.4 Atribuições

- Atribuição simples
- Atribuição de tuplas (múltiplos valores)
- Atribuibilidade de tipos
- Operadores de atribuição compostos (`+=`, `-=`, etc.)

### 2.5 Declarações de Tipo

- Definição de novos tipos com `type`
- Tipos nomeados vs tipos anônimos
- Conversão entre tipos
- Aliás de tipos

### 2.6 Pacotes e Arquivos

- Organização em pacotes
- Cláusula `package`
- Importação de pacotes
- O pacote `main`
- Estrutura de diretórios
- Inicialização de pacotes

### 2.7 Escopo

- Escopo de bloco
- Escopo léxico
- Escopo de funções
- Escopo de pacote
- Escopo de universo
- Sombreamento de variáveis (shadowing)

## Conceitos Importantes

### Zero Values

Go inicializa automaticamente variáveis com valores zero:

- `0` para tipos numéricos
- `false` para booleanos
- `""` (string vazia) para strings
- `nil` para ponteiros, slices, maps, channels, funções e interfaces

### Visibilidade

- Identificadores que começam com letra maiúscula são **exportados** (públicos)
- Identificadores que começam com letra minúscula são **não exportados** (privados ao pacote)

### Declaração vs Atribuição

- **Declaração**: introduz um novo identificador
- **Atribuição**: atualiza o valor de uma variável existente

## Boas Práticas

1. Use nomes descritivos e significativos
2. Prefira declarações curtas (`:=`) dentro de funções
3. Agrupe declarações relacionadas
4. Evite sombreamento de variáveis quando possível
5. Organize código em pacotes coesos
6. Mantenha escopos pequenos e localizados

## Exercícios

Os exercícios deste capítulo exploram:

- Manipulação de variáveis e escopos
- Organização de código em pacotes
- Boas práticas de nomenclatura
- Conversão entre tipos
- Uso efetivo de ponteiros

## Recursos Adicionais

- [Documentação Oficial Go - Package](https://golang.org/doc/code.html)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go by Example](https://gobyexample.com/)
