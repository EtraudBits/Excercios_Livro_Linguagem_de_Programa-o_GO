# Mandelbrot YCbCr - Espaço de Cores de Vídeo

## 📝 Sobre Este Código

Este programa gera o conjunto de Mandelbrot usando o espaço de cores **YCbCr**, amplamente utilizado em sistemas de vídeo, compressão de imagens (JPEG) e transmissão de TV.

## 🎨 O que é YCbCr?

**YCbCr** separa a informação de cor em três componentes:

- **Y** (Luma): Luminância ou brilho (0-255)
  - 0 = preto, 255 = branco
  - É a informação em "preto e branco"
- **Cb** (Chroma Blue): Diferença azul-amarelo (0-255)
  - 0 = máximo azul, 255 = máximo amarelo
  - 128 = neutro (sem cor adicional)
- **Cr** (Chroma Red): Diferença vermelho-verde (0-255)
  - 0 = máximo verde, 255 = máximo vermelho
  - 128 = neutro

## 🌈 Esquema de Cores Usado

```go
color.YCbCr{
    Y:  255 - n,        // Luminância diminui com iterações
    Cb: 128 + n/2,      // Levemente para o amarelo
    Cr: 128 + n,        // Progressivamente mais vermelho
}
```

### Interpretação

- **Pontos próximos ao conjunto**: Escuros e neutros
- **Pontos que escapam rápido**: Claros e avermelhados
- **Pontos intermediários**: Transição suave de cores

## 🚀 Como Usar

```bash
# Gerar a imagem
go run mandelbrot_ycbcr.go > mandelbrot_ycbcr.png

# Ver a imagem
xdg-open mandelbrot_ycbcr.png  # Linux
open mandelbrot_ycbcr.png       # macOS
start mandelbrot_ycbcr.png      # Windows
```

## 🎯 Por que YCbCr é Importante?

### 1. **Percepção Humana**

O olho humano é muito mais sensível a mudanças de **brilho** (Y) do que a mudanças de **cor** (Cb, Cr). YCbCr aproveita isso!

### 2. **Compressão Eficiente**

Como separamos luminância de cor, podemos:

- Manter alta resolução para Y (brilho)
- Reduzir resolução de Cb e Cr (cor)
- Resultado: arquivos menores sem perda visível de qualidade

### 3. **Compatibilidade**

- TV preto e branco pode exibir só o canal Y
- TV colorida usa Y + Cb + Cr
- Retrocompatibilidade perfeita!

## 🌍 Onde YCbCr é Usado na Vida Real?

### 📺 Televisão Digital

```
Sinal de TV → YCbCr → Transmissão
```

- TV aberta (digital terrestre)
- TV a cabo
- Satélite

### 📷 JPEG (Fotos na Internet)

```
Câmera RGB → Conversão YCbCr → Compressão → Arquivo JPEG
```

- Reduz tamanho em 10x ou mais
- Mantém qualidade visual
- Praticamente todas as fotos na web

### 🎬 Vídeo Streaming

```
Vídeo → YCbCr → H.264/H.265 → Netflix/YouTube
```

- **YouTube**: 4:2:0 YCbCr (cor em ¼ da resolução de luminância)
- **Netflix**: Mesmo esquema para streaming eficiente
- **Zoom/Teams**: Videoconferências usam YCbCr

### 📱 Câmeras Digitais

```
Sensor → Bayer Pattern → YCbCr → Processamento
```

- Câmeras de celular
- DSLRs profissionais
- Webcams

## 🔬 Subamostragem de Crominância

Um conceito chave do YCbCr é a **subamostragem**:

### 4:4:4 (Sem subamostragem)

```
Y  Y  Y  Y
Cb Cb Cb Cb
Cr Cr Cr Cr
```

Cada pixel tem informação completa.

### 4:2:2 (TV e vídeo profissional)

```
Y  Y  Y  Y
Cb    Cb
Cr    Cr
```

Cor a cada 2 pixels horizontalmente.

### 4:2:0 (JPEG, YouTube, Netflix)

```
Y  Y  Y  Y
Cb    Cb
Cr    Cr
```

Cor a cada 2x2 bloco de pixels.

**Economia**: 4:2:0 usa 50% menos dados que 4:4:4!

## 🧮 Conversão YCbCr ↔ RGB

### YCbCr para RGB

```
R = Y + 1.402 × (Cr - 128)
G = Y - 0.344 × (Cb - 128) - 0.714 × (Cr - 128)
B = Y + 1.772 × (Cb - 128)
```

### RGB para YCbCr

```
Y  =  0.299×R + 0.587×G + 0.114×B
Cb = -0.169×R - 0.331×G + 0.500×B + 128
Cr =  0.500×R - 0.419×G - 0.081×B + 128
```

**Nota**: Go faz essa conversão automaticamente quando você usa `img.Set()` com `color.YCbCr` em uma imagem RGBA!

## 🎨 Experimentos com Cores

### Tons de azul

```go
return color.YCbCr{
    Y:  200,        // Claro
    Cb: 200,        // Azul
    Cr: 128,        // Neutro
}
```

### Tons sépia (vintage)

```go
return color.YCbCr{
    Y:  255 - n,
    Cb: 128 - n/4,  // Levemente azul
    Cr: 128 + n/2,  // Bastante vermelho
}
```

### Preto e branco puro

```go
return color.YCbCr{
    Y:  255 - n,
    Cb: 128,        // Neutro
    Cr: 128,        // Neutro
}
```

## 📊 YCbCr vs RGBA

| Aspecto         | RGBA                            | YCbCr                                    |
| --------------- | ------------------------------- | ---------------------------------------- |
| **Componentes** | Vermelho, Verde, Azul, Alpha    | Luminância, Azul-Amarelo, Vermelho-Verde |
| **Separação**   | Não separa brilho e cor         | Separa explicitamente                    |
| **Compressão**  | Ruim (todos componentes iguais) | Excelente (Y prioritário)                |
| **Uso**         | Telas, computação gráfica       | Vídeo, fotos, transmissão                |
| **Percepção**   | Não otimizado                   | Otimizado para olho humano               |

## 💻 Conceitos de Go Utilizados

```go
// Tipo YCbCr
color.YCbCr{
    Y:  brightness,  // 0-255
    Cb: blue_yellow, // 0-255 (128 = neutro)
    Cr: red_green,   // 0-255 (128 = neutro)
}

// Go converte automaticamente para RGBA quando necessário
img.Set(px, py, ycbcrColor)  // Funciona perfeitamente!
```

## 🎓 O que Você Aprende

- ✅ Espaço de cores YCbCr e sua importância
- ✅ Como separar luminância de crominância
- ✅ Fundamentos de compressão de vídeo/imagem
- ✅ Como sistemas de TV e streaming funcionam
- ✅ Conversão automática entre espaços de cores em Go

## 📚 Padrões e Especificações

- **BT.601**: Padrão para TV SD (definição padrão)
- **BT.709**: Padrão para HDTV (alta definição)
- **BT.2020**: Padrão para Ultra HD / 4K / HDR
- **JPEG**: Usa YCbCr com subamostragem 4:2:0
- **MPEG-2/4**: Vídeo em YCbCr

## 🔗 Veja Também

- `../rgba/` - Versão usando RGBA tradicional
- Documentação Go: `image/color` package
- Wikipedia: [YCbCr](https://pt.wikipedia.org/wiki/YCbCr)
- JPEG Compression: Como funciona a compressão de imagens

## 💡 Curiosidade

Quando você tira uma foto com seu celular e ela fica com 2-3 MB ao invés de 30 MB, você tem o YCbCr (e JPEG) para agradecer! A separação de luminância e crominância é fundamental para isso.
