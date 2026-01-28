# Mandelbrot RGBA - Cores Vibrantes

## 📝 Sobre Este Código

Este programa gera o conjunto de Mandelbrot usando o espaço de cores **RGBA** (Red, Green, Blue, Alpha), o formato mais comum em computação gráfica.

## 🎨 Esquema de Cores

A coloração é baseada no número de iterações (`n`) necessárias para o ponto escapar:

```go
color.RGBA{
    R: 255 - n,     // Vermelho: 255 → 0 (diminui)
    G: n * 2,       // Verde: 0 → 510 (aumenta rápido)
    B: n * 3,       // Azul: 0 → 765 (aumenta muito rápido)
    A: 255,         // Alpha: sempre opaco
}
```

### O que isso significa?

- **Pontos que escapam rápido** (poucas iterações): Tons avermelhados
- **Pontos que escapam devagar** (muitas iterações): Tons azulados
- **Pontos no conjunto** (não escapam): Preto

## 🚀 Como Usar

```bash
# Gerar a imagem
go run mandelbrot_colorido.go > mandelbrot_rgba.png

# Ver a imagem (Linux)
xdg-open mandelbrot_rgba.png

# Ver a imagem (macOS)
open mandelbrot_rgba.png

# Ver a imagem (Windows)
start mandelbrot_rgba.png
```

## 🔧 Parâmetros Configuráveis

No código, você pode ajustar:

```go
const (
    xmin, ymin, xmax, ymax = -2, -2, +2, +2  // Área do plano complexo
    width, height          = 1024, 1024       // Tamanho da imagem
)

const iterations = 200  // Máximo de iterações
```

## 🎯 Experimentos Sugeridos

### 1. Zoom em Regiões Interessantes

```go
// Exemplo: Zoom na "cauda" do Mandelbrot
xmin, ymin, xmax, ymax = -0.8, -0.2, -0.4, 0.2
```

### 2. Mais Detalhes

```go
const iterations = 500  // Mais iterações = mais cores
```

### 3. Esquemas de Cores Diferentes

**Cores quentes**:

```go
return color.RGBA{
    R: 255,
    G: 255 - n,
    B: 0,
    A: 255,
}
```

**Escala de cinza com contraste**:

```go
return color.RGBA{
    R: 255 - n*3,
    G: 255 - n*3,
    B: 255 - n*3,
    A: 255,
}
```

## 🌈 Por que RGBA?

**Vantagens**:

- Formato universal, funciona em todas as plataformas
- Fácil de entender (R=vermelho, G=verde, B=azul)
- Ideal para telas e web
- Suporte nativo em todas as bibliotecas gráficas

**Onde é usado**:

- Monitores e TVs (cada pixel tem RGB)
- Navegadores web
- Jogos de vídeo game
- Software de edição de imagem

## 📊 Estrutura do Código

1. **Função `mandelbrot(z complex128)`**: Testa se um ponto está no conjunto
2. **Loop duplo**: Percorre cada pixel da imagem
3. **Conversão de coordenadas**: Mapeia pixels para números complexos
4. **Geração de cores**: Baseada no número de iterações

## 🧮 A Matemática Por Trás

### Fórmula Iterativa

```
z[0] = 0
z[n+1] = z[n]² + c
```

Onde `c` é o ponto que estamos testando.

### Critério de Escape

Se `|z| > 2`, o ponto vai para infinito (não está no conjunto).

### Exemplo Prático

Testando o ponto `c = 0.5 + 0.5i`:

- Iteração 0: z = 0
- Iteração 1: z = 0² + (0.5+0.5i) = 0.5+0.5i
- Iteração 2: z = (0.5+0.5i)² + (0.5+0.5i) = 0.5+1i
- Iteração 3: z = (0.5+1i)² + (0.5+0.5i) = -0.25+1.5i
- Iteração 4: |z| ≈ 1.52 (ainda < 2, continua)
- ... eventualmente escapa

## 💻 Conceitos de Go Utilizados

```go
// Números complexos
z := complex(x, y)          // Cria número complexo
cmplx.Abs(v)                // Módulo (|v|)

// Imagens
img := image.NewRGBA(...)   // Cria imagem
img.Set(px, py, cor)        // Define pixel

// Cores
color.RGBA{R, G, B, A}      // Cor RGBA
color.Black                 // Preto

// I/O
png.Encode(os.Stdout, img)  // Grava PNG
```

## 🎓 O que Você Aprende

- ✅ Manipulação de números complexos
- ✅ Criação de imagens pixel a pixel
- ✅ Espaço de cores RGBA
- ✅ Iteração numérica
- ✅ Conversão entre sistemas de coordenadas
- ✅ Geração de fractais

## 🔗 Veja Também

- `../ycbcr/` - Versão usando espaço YCbCr
- `../../exercicio3.6/` - Super amostragem (anti-aliasing)
- `../../progMandelbrot/` - Versão básica em escala de cinza

---

## 👨‍💻 Autor

<div align="center">

### 🐹 Gopher Developer

**Duarte Rodrigo Santos de Oliveira**

_Estudante Autodidata da linguagem Go_

[![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/duarte-backend-golang)

</div>
