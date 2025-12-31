# Seção 1.4 - GIF Animados

## 📚 Sobre o Exercício

Este exercício do livro "A Linguagem de Programação Go" demonstra como criar GIFs animados programaticamente, gerando figuras de Lissajous - padrões matemáticos harmônicos que criam formas elegantes e hipnotizantes.

## 🎯 Objetivo

O programa gera uma animação GIF de curvas de Lissajous, que são figuras produzidas pela composição de dois movimentos harmônicos perpendiculares. O resultado é salvo em um arquivo `lissajous.gif`.

## 🔧 Funcionalidades

- ✅ Gera figuras de Lissajous matematicamente
- ✅ Cria animações GIF com múltiplos frames
- ✅ Usa paleta de cores customizada (preto, verde, laranja)
- ✅ Alterna cores entre frames para efeito visual
- ✅ Loop infinito da animação
- ✅ Frequência aleatória para variedade

## 💻 Como Usar

```bash
# Executar o programa (gera lissajous.gif no diretório atual)
go run gif_animados.go

# Compilar e executar
go build gif_animados.go
./gif_animados

# Abrir o GIF gerado
xdg-open lissajous.gif  # Linux
open lissajous.gif      # macOS
start lissajous.gif     # Windows
```

## 📖 Conceitos Aprendidos

### 1. **Pacote `image/gif`**

- Criação e codificação de arquivos GIF
- `gif.GIF` estrutura que armazena frames e delays
- `gif.EncodeAll()` codifica toda a animação

### 2. **Pacote `image`**

- Trabalho com imagens e coordenadas
- `image.Rect()` define retângulos
- `image.NewPaletted()` cria imagens paletizadas

### 3. **Pacote `image/color`**

- Definição de cores RGB
- `color.RGBA` para cores com canal alfa
- Paletas de cores indexadas

### 4. **Pacote `math`**

- Funções trigonométricas: `math.Sin()`, `math.Pi`
- Cálculos para gerar curvas harmônicas

### 5. **Pacote `math/rand`**

- Geração de números aleatórios
- `rand.Float64()` retorna float entre 0.0 e 1.0

### 6. **Gerenciamento de Arquivos**

- `os.Create()` cria novos arquivos
- `defer` garante fechamento do arquivo
- Interface `io.Writer` para escrita genérica

### 7. **Imagens Paletizadas**

- Uso de índices ao invés de RGB direto
- Economia de espaço em arquivos
- `SetColorIndex()` define pixel por índice

### 8. **Curvas de Lissajous**

- Composição de movimentos harmônicos
- `x = sin(t)` e `y = sin(t*freq + phase)`
- Variação de fase cria efeito de rotação

## 🎨 Parâmetros Configuráveis

| Parâmetro | Valor     | Descrição                              |
| --------- | --------- | -------------------------------------- |
| `cycles`  | 5         | Número de oscilações completas         |
| `res`     | 0.001     | Resolução angular (menor = mais suave) |
| `size`    | 100       | Tamanho do canvas (201x201 pixels)     |
| `nframes` | 64        | Quantidade de frames na animação       |
| `delay`   | 8         | Delay entre frames (80ms)              |
| `freq`    | aleatório | Frequência relativa (0 a 3)            |

## 🌍 Casos de Uso no Mundo Real

Este tipo de código pode ser usado em:

### 1. **Visualização Científica**

- Demonstrar conceitos de física e matemática
- Osciladores harmônicos
- Análise de frequências

### 2. **Arte Generativa**

- Criar padrões únicos e aleatórios
- Design procedural
- NFTs e arte digital

### 3. **Animações Web**

- GIFs para sites e redes sociais
- Loading screens criativas
- Elementos visuais dinâmicos

### 4. **Educação**

- Material didático para matemática
- Demonstrações interativas
- Visualização de conceitos abstratos

### 5. **Data Visualization**

- Representar dados cíclicos
- Gráficos animados
- Relatórios dinâmicos

## 🔍 Exemplo de Estrutura

```
Frame 1 (verde)  →  Frame 2 (laranja)  →  Frame 3 (verde)  →  ...
      ↑                                                            ↓
      ←  ←  ←  ←  ←  Frame 64 (laranja)  ←  ←  ←  ←  ←  ←  ←  ←
                          (Loop infinito)
```

## 🎭 Detalhes Técnicos

### Fórmulas Matemáticas

- **Coordenada X**: $x = \sin(t)$
- **Coordenada Y**: $y = \sin(t \times freq + phase)$
- **Transformação para pixels**: $pixel = size + \lfloor x \times size + 0.5 \rfloor$

### Estrutura da Animação GIF

```go
gif.GIF{
    Image:     []image.Paletted  // Array de frames
    Delay:     []int             // Delays em 10ms
    LoopCount: 0                 // 0 = loop infinito
}
```

### Paleta de Cores

- **Índice 0**: Preto `#000000` (fundo)
- **Índice 1**: Verde `#00FF00`
- **Índice 2**: Laranja `#FF7F00`

## ⚙️ Modificações do Código Original

O código foi modificado em relação ao livro:

```go
// Livro original
anim := gif.GIF{LoopCount: nframes}  // Para após nframes

// Código modificado
anim := gif.GIF{LoopCount: 0}        // Loop infinito
```

## ⚠️ Limitações Atuais

- Nome do arquivo é fixo (`lissajous.gif`)
- Não aceita parâmetros via linha de comando
- Paleta limitada a 3 cores
- Frequência é totalmente aleatória
- Tamanho da imagem não é configurável em runtime

## 🚀 Possíveis Melhorias

1. Aceitar nome do arquivo como argumento
2. Parâmetros configuráveis via flags (cycles, size, nframes, etc.)
3. Paleta de cores customizável
4. Opção de exportar para outros formatos (PNG, APNG)
5. Controle de seed do random para reproduzibilidade
6. Adicionar mais cores e gradientes
7. Suporte a diferentes tipos de curvas (espirais, roses, etc.)
8. Modo interativo para preview antes de salvar
9. Gerar múltiplos GIFs em batch
10. Adicionar texto ou marca d'água

## 📝 Notas Importantes

- **`defer f.Close()`**: Garante que o arquivo será fechado mesmo se houver erro
- **Imagens paletizadas**: Mais eficientes para GIFs (menor tamanho)
- **`+0.5` no arredondamento**: Técnica para arredondamento correto de float para int
- **`io.Writer`**: Interface que permite escrever tanto em arquivos quanto em HTTP, buffers, etc.
- **Loop infinito**: `LoopCount: 0` faz o GIF repetir eternamente

## 🔬 Experimente

Tente modificar os valores e veja o resultado:

```go
// Curva mais lenta e detalhada
const cycles = 10
const res = 0.0001

// Animação mais rápida
const delay = 2

// Imagem maior
const size = 200

// Mais frames = animação mais suave
const nframes = 128
```

## 🔗 Referências

- [Documentação image/gif](https://pkg.go.dev/image/gif)
- [Documentação image](https://pkg.go.dev/image)
- [Curvas de Lissajous - Wikipedia](https://pt.wikipedia.org/wiki/Curva_de_Lissajous)
- Livro: "A Linguagem de Programação Go" - Capítulo 1, Seção 1.4

---

**Data de estudo**: 30 de dezembro de 2025  
**Status**: ✅ Concluído e funcional  
**Output**: `lissajous.gif`
