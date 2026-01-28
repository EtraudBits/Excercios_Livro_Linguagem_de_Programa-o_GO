# Programa Mandelbrot - Versão Básica

## 📝 Sobre Este Código

Este é o programa introdutório que demonstra como gerar o conjunto de Mandelbrot usando números complexos em Go. É a base para todos os exercícios posteriores da seção.

## 🎯 O que é o Conjunto de Mandelbrot?

Um **fractal** descoberto por Benoit Mandelbrot em 1980. É definido por uma regra simples:

```
z[0] = 0
z[n+1] = z[n]² + c
```

Para cada ponto `c` no plano complexo:

- Se a sequência permanece **limitada** (não vai para infinito) → `c` está **no conjunto** (preto)
- Se a sequência **escapa** para infinito → `c` está **fora do conjunto** (cinza)

## 🚀 Como Usar

```bash
# Gerar a imagem
go run mandelbrot.go > mandelbrot.png

# Ver a imagem
xdg-open mandelbrot.png  # Linux
open mandelbrot.png       # macOS
start mandelbrot.png      # Windows
```

## 🧮 A Matemática

### Critério de Escape

**Propriedade matemática**: Se em qualquer iteração `|z| > 2`, então z vai divergir para infinito.

Portanto, podemos parar de iterar assim que `|z| > 2` e declarar que o ponto está fora do conjunto.

### Exemplo Numérico

Testando o ponto `c = 0.5 + 0i`:

```
z[0] = 0
z[1] = 0² + (0.5+0i) = 0.5
z[2] = (0.5)² + 0.5 = 0.75
z[3] = (0.75)² + 0.5 = 1.0625
z[4] = (1.0625)² + 0.5 = 1.6289
z[5] = (1.6289)² + 0.5 = 3.1533
```

Em z[5], temos |z| ≈ 3.15 > 2, então sabemos que o ponto **escapa** (não está no conjunto).

### Coloração

```go
// Pontos no conjunto: preto
return color.Black

// Pontos fora: cinza baseado na velocidade
// Escapa rápido (n pequeno) → cinza claro
// Escapa devagar (n grande) → cinza escuro
return color.Gray{255 - contrast*n}
```

## 🎨 Parâmetros Ajustáveis

### Região do Plano Complexo

```go
const (
    xmin, ymin, xmax, ymax = -2, -2, +2, +2
)
```

Modifique para fazer zoom:

```go
// Zoom na "cauda" esquerda
xmin, ymin, xmax, ymax = -2.0, -1.0, -1.0, 1.0
```

### Resolução

```go
const (
    width, height = 1024, 1024  // Pixels
)
```

Resoluções sugeridas:

- `512, 512` - Rápido, boa qualidade
- `1024, 1024` - Padrão
- `2048, 2048` - Alta qualidade
- `4096, 4096` - 4K (muito lento)

### Número de Iterações

```go
const iterations = 200
```

- Menos iterações (50-100): Mais rápido, menos detalhes
- Mais iterações (500-1000): Mais lento, mais detalhes

### Contraste

```go
const contrast = 15
```

- Contraste baixo (5-10): Transições suaves
- Contraste alto (20-30): Bordas mais definidas

## 💻 Conceitos de Go Utilizados

### Números Complexos

```go
// Criar
z := complex(x, y)      // x + yi

// Operar
z = z*z + c             // Multiplicação e soma

// Módulo (distância da origem)
import "math/cmplx"
cmplx.Abs(z)            // √(real² + imag²)
```

### Geração de Imagens

```go
// Criar imagem
img := image.NewRGBA(image.Rect(0, 0, width, height))

// Definir pixel
img.Set(px, py, cor)

// Salvar PNG
import "image/png"
png.Encode(os.Stdout, img)
```

### Cores

```go
// Escala de cinza
color.Gray{valor}       // 0 = preto, 255 = branco

// Preto puro
color.Black
```

## 🎓 O que Você Aprende

- ✅ Uso básico de `complex128`
- ✅ Iteração numérica
- ✅ Mapeamento de coordenadas (pixel → complexo)
- ✅ Geração de imagens PNG
- ✅ Conceitos de fractais
- ✅ Teste de convergência/divergência

## 🔬 Experimentos Sugeridos

### 1. Explorar Regiões Interessantes

```go
// "Elephant Valley"
xmin, ymin, xmax, ymax = 0.275, 0.0, 0.285, 0.01

// "Seahorse Valley"
xmin, ymin, xmax, ymax = -0.75, -0.1, -0.74, 0.0

// Mini-Mandelbrot
xmin, ymin, xmax, ymax = -0.16, 1.035, -0.14, 1.045
```

### 2. Mudar Cores

```go
// Inverter (fundo branco)
return color.Gray{contrast*n}

// Alto contraste
return color.Gray{uint8(n * 5)}
```

### 3. Aumentar Detalhes

```go
const iterations = 500
const contrast = 5
```

### 4. Criar Variações

**Burning Ship**:

```go
v = complex(math.Abs(real(v)), math.Abs(imag(v)))
v = v*v + z
```

**Tricorn**:

```go
v = cmplx.Conj(v)*cmplx.Conj(v) + z
```

## 🌍 Aplicações na Vida Real

### 1. **Compressão de Imagens**

Algoritmos baseados em fractais para JPEG2000

### 2. **Antenas**

Design de antenas fractais (menores e mais eficientes) para celulares

### 3. **Gráficos de Computador**

Texturas procedurais em jogos e filmes

### 4. **Modelagem Natural**

Costas marítimas, montanhas, nuvens

### 5. **Criptografia**

Geradores de números pseudo-aleatórios

### 6. **Análise de Mercados**

Padrões fractais em finanças (Teoria de Dow)

## 📊 Performance

Para 1024×1024 pixels:

- **Tempo típico**: 150-250ms
- **Memória**: ~8MB
- **Tamanho do PNG**: ~100-300KB

## 🐛 Problemas Comuns

### Imagem toda preta

```go
// Verifique se está usando os limites corretos
// O conjunto está entre -2 e +2
```

### Imagem borrada

```go
// Aumente o número de iterações
const iterations = 500
```

### Muito lento

```go
// Reduza a resolução
width, height = 512, 512
```

## 🔗 Próximos Passos

Depois de dominar este programa:

1. **[exercicio3.5/](../exercicio3.5/)** - Adicione cores (RGBA e YCbCr)
2. **[exercicio3.6/](../exercicio3.6/)** - Implemente anti-aliasing
3. **[exercicio3.7/](../exercicio3.7/)** - Explore o Método de Newton
4. **[exercicio3.8/](../exercicio3.8/)** - Teste precisão numérica com zoom profundo

## 💡 Curiosidades

- O conjunto de Mandelbrot é **conectado** - há um caminho contínuo entre quaisquer dois pontos
- Sua **fronteira** tem dimensão fractal ≈ 2 (não é uma curva 1D nem área 2D!)
- Zoom infinito revela sempre novos detalhes (auto-similaridade)
- A NASA usou antenas fractais no rover Perseverance em Marte! 🚀

## 📚 Referências

- **Livro**: "A Linguagem de Programação Go" - Capítulo 3.3
- **Wikipedia**: [Conjunto de Mandelbrot](https://pt.wikipedia.org/wiki/Conjunto_de_Mandelbrot)
- **Vídeos**: Pesquise "Mandelbrot zoom" no YouTube para animações incríveis

Divirta-se explorando o infinito! ✨
