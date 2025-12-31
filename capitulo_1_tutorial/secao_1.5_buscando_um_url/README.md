# Seção 1.5 - Buscando um URL

## 📚 Sobre o Exercício

Este exercício do livro "A Linguagem de Programação Go" demonstra como fazer requisições HTTP em Go, buscando e exibindo o conteúdo de URLs fornecidas como argumentos na linha de comando.

## 🎯 Objetivo

O programa faz requisições HTTP GET para uma ou mais URLs e imprime o conteúdo HTML retornado no terminal. É uma introdução prática ao uso da biblioteca padrão `net/http` do Go.

## 🔧 Funcionalidades

- ✅ Aceita múltiplas URLs como argumentos
- ✅ Faz requisições HTTP GET
- ✅ Trata erros de conexão e leitura adequadamente
- ✅ Exibe o conteúdo completo de cada página
- ✅ Fecha conexões corretamente para evitar vazamento de recursos

## 💻 Como Usar

```bash
# Buscar uma única URL
go run buscando_um_url.go http://gopl.io

# Buscar múltiplas URLs
go run buscando_um_url.go http://gopl.io http://golang.org

# Compilar e executar
go build buscando_um_url.go
./buscando_um_url http://example.com
```

## 📖 Conceitos Aprendidos

### 1. **Pacote `net/http`**

- Biblioteca padrão do Go para requisições HTTP
- `http.Get(url)` faz requisições GET de forma simples

### 2. **Pacote `io`**

- `io.ReadAll()` lê todo o conteúdo de um leitor (Reader)
- Substitui o antigo `ioutil.ReadAll()` (deprecado desde Go 1.16)

### 3. **Tratamento de Erros**

- Go utiliza retorno de múltiplos valores
- Sempre verificar `err != nil` após operações que podem falhar
- Uso de `continue` para pular iterações com erro

### 4. **Gerenciamento de Recursos**

- `resp.Body.Close()` libera a conexão HTTP
- Importante fechar para evitar vazamento de recursos
- Boa prática: fechar logo após ler o conteúdo

### 5. **Argumentos de Linha de Comando**

- `os.Args` contém todos os argumentos
- `os.Args[0]` é o nome do programa
- `os.Args[1:]` são os argumentos fornecidos pelo usuário

### 6. **Saída de Erros**

- `os.Stderr` para mensagens de erro
- `fmt.Fprintf()` para escrever em diferentes streams
- Separar erros da saída normal facilita redirecionamento

## 🌍 Casos de Uso no Mundo Real

Este padrão de código pode ser usado em:

### 1. **Web Scraping**

- Coletar dados de websites
- Monitorar mudanças em páginas
- Extrair informações específicas

### 2. **Monitoramento de Serviços**

- Verificar se sites estão disponíveis
- Monitorar tempo de resposta
- Alertas de downtime

### 3. **Integrações com APIs**

- Consumir APIs REST
- Buscar dados de serviços externos
- Sincronizar informações

### 4. **Testes de Carga**

- Verificar desempenho de servidores
- Simular múltiplos acessos
- Análise de capacidade

### 5. **Ferramentas CLI**

- Criar utilitários de linha de comando
- Automatizar downloads
- Processar conteúdo web em batch

## 🔍 Exemplo de Saída

```bash
$ go run buscando_um_url.go http://gopl.io
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN"
          "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
  <title>The Go Programming Language</title>
...
```

## ⚠️ Limitações Atuais

- Não possui timeout para requisições longas
- Não suporta HTTPS com certificados inválidos
- Não faz parsing do conteúdo HTML
- Não salva o conteúdo em arquivos
- Não exibe status code ou headers da resposta

## 🚀 Possíveis Melhorias

1. Adicionar timeout nas requisições
2. Exibir código de status HTTP
3. Mostrar headers da resposta
4. Salvar conteúdo em arquivos
5. Adicionar flag para controlar verbosidade
6. Implementar pool de conexões para múltiplas URLs
7. Adicionar suporte a requisições POST/PUT/DELETE
8. Medir e exibir tempo de resposta

## 📝 Notas Importantes

- **Sempre feche `resp.Body`**: Essencial para liberar recursos de rede
- **Verifique erros**: Go não tem exceções, sempre verifique retornos de erro
- **Use `io.ReadAll`**: Não use `ioutil.ReadAll` (deprecado)
- **Ordem importa**: Feche o body após ler seu conteúdo, não antes

## 🔗 Referências

- [Documentação net/http](https://pkg.go.dev/net/http)
- [Documentação io](https://pkg.go.dev/io)
- Livro: "A Linguagem de Programação Go" - Capítulo 1, Seção 1.5

---

**Data de estudo**: 30 de dezembro de 2025  
**Status**: ✅ Concluído e funcional
