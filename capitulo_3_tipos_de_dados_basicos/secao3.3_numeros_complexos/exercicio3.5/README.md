# Exercício 3.5 - Conjunto de Mandelbrot Colorido

## 📋 Descrição

Este exercício implementa o famoso **conjunto de Mandelbrot** usando diferentes espaços de cores para criar visualizações vibrantes deste fractal matemático. O objetivo é comparar como diferentes representações de cores (RGBA e YCbCr) afetam a aparência final do fractal.

## 🎨 O que é o Conjunto de Mandelbrot?

O conjunto de Mandelbrot é um dos fractais mais famosos da matemática. Ele é definido pela seguinte regra simples:

1. Começando com um número complexo `c`
2. Aplicamos repetidamente a fórmula: `z = z² + c` (começando com z = 0)
3. Se o valor de `z` permanece limitado (não vai para infinito), então `c` está no conjunto de Mandelbrot

**Propriedade importante**: Se em algum momento |z| > 2, sabemos que o ponto vai divergir para infinito, então não está no conjunto.

## 🌈 Espaços de Cores Implementados

### 1. RGBA (Red, Green, Blue, Alpha)

- **Pasta**: `rgba/`
- **Arquivo**: `mandelbrot_colorido.go`
- **Descrição**: Usa o espaço de cores mais comum em computação gráfica
- **Componentes**:
  - **R** (Red): Vermelho (0-255)
  - **G** (Green): Verde (0-255)
  - **B** (Blue): Azul (0-255)
  - **A** (Alpha): Transparência (0-255)

**Como funciona**:

```go
color.RGBA{
    R: 255 - n,     // Vermelho diminui com mais iterações
    G: n * 2,       // Verde aumenta
    B: n * 3,       // Azul aumenta ainda mais rápido
    A: 255,         // Totalmente opaco
}
```

### 2. YCbCr (Luminância e Crominância)

- **Pasta**: `ycbcr/`
- **Arquivo**: `mandelbrot_ycbcr.go`
- **Descrição**: Espaço de cores usado em sistemas de vídeo (JPEG, MPEG, transmissão de TV)
- **Componentes**:
  - **Y**: Luminância (brilho) - informação em preto e branco
  - **Cb**: Crominância azul-amarelo (diferença azul)
  - **Cr**: Crominância vermelho-verde (diferença vermelho)

**Por que YCbCr é importante?**

- Separa brilho (Y) de cor (Cb, Cr)
- O olho humano é mais sensível a mudanças de brilho que de cor
- Permite compressão mais eficiente em vídeos e imagens JPEG
- Usado em TV digital, streaming, e câmeras digitais

## 🚀 Como Executar

### RGBA (Colorido tradicional)

```bash
cd rgba
go run mandelbrot_colorido.go > mandelbrot_rgba.png
```

### YCbCr (Espaço de vídeo)

```bash
cd ycbcr
go run mandelbrot_ycbcr.go > mandelbrot_ycbcr.png
```

## 🔍 Conceitos Matemáticos

### Números Complexos

Um número complexo é da forma `a + bi`, onde:

- `a` é a parte real
- `b` é a parte imaginária
- `i` é a unidade imaginária (√-1)

No plano complexo:

- Eixo horizontal = parte real
- Eixo vertical = parte imaginária

### A Fórmula de Mandelbrot

```
z[n+1] = z[n]² + c
```

Onde:

- `c` é o ponto que estamos testando (constante)
- `z[0]` = 0 (começamos sempre do zero)
- Iteramos até z escapar (|z| > 2) ou atingir o limite de iterações

### Por que |z| > 2 é o critério?

É uma propriedade matemática provada: se |z| ultrapassa 2 em qualquer iteração, o valor vai divergir para infinito. Se nunca ultrapassa, pode estar no conjunto.

## 🌍 Aplicações na Vida Real

### 1. **Computação Gráfica e Arte Digital**

- Criação de texturas procedurais para jogos
- Efeitos visuais em filmes e animações
- Arte generativa e design

### 2. **Compressão de Vídeo (YCbCr)**

- **YouTube**: Usa YCbCr para comprimir vídeos
- **Netflix**: Streaming eficiente separando luminância e cor
- **JPEG**: Fotos na internet usam YCbCr para menor tamanho
- **TV Digital**: Transmissão de sinal de televisão

### 3. **Antenas e Telecomunicações**

- Design de antenas fractais (baseadas em Mandelbrot)
- Antenas menores e mais eficientes para celulares
- Telecomunicações móveis

### 4. **Análise de Sistemas Dinâmicos**

- Previsão do tempo (sistemas caóticos)
- Modelagem de populações
- Mercado financeiro (comportamento caótico)

### 5. **Processamento de Imagens Médicas**

- Análise de texturas em imagens de raios-X
- Detecção de padrões em tomografias

## 📊 Comparação RGBA vs YCbCr

| Característica       | RGBA                         | YCbCr                                    |
| -------------------- | ---------------------------- | ---------------------------------------- |
| **Uso Principal**    | Computação gráfica           | Vídeo e transmissão                      |
| **Componentes**      | Vermelho, Verde, Azul, Alpha | Luminância, Azul-Amarelo, Vermelho-Verde |
| **Percepção Humana** | Não otimizado                | Otimizado (separa brilho de cor)         |
| **Compressão**       | Menos eficiente              | Mais eficiente                           |
| **Onde Usar**        | Telas, jogos, web            | TV, vídeo, JPEG, streaming               |

## 🎓 O que Você Aprende

1. **Números Complexos em Go**: Uso do tipo `complex128` e pacote `math/cmplx`
2. **Espaços de Cores**: Diferença entre RGBA e YCbCr
3. **Processamento de Imagens**: Como criar imagens pixel por pixel
4. **Fractais**: Conceitos básicos de geometria fractal
5. **Iteração Numérica**: Como aplicar fórmulas repetidamente
6. **Mapeamento de Coordenadas**: Conversão entre pixels e plano complexo

## 🔬 Experimentos Sugeridos

1. **Mude os limites**: Altere `xmin, ymin, xmax, ymax` para fazer zoom em regiões interessantes
2. **Aumente iterações**: Mude `const iterations = 200` para ver mais detalhes
3. **Crie novos esquemas de cores**: Experimente outras fórmulas para R, G, B
4. **Compare as saídas**: Abra as duas imagens (RGBA e YCbCr) e compare as diferenças visuais

## 📚 Referências

- **Livro**: "A Linguagem de Programação Go" - Seção 3.3 (Números Complexos)
- **Mandelbrot Set**: [Wikipedia](https://pt.wikipedia.org/wiki/Conjunto_de_Mandelbrot)
- **YCbCr Color Space**: Usado em JPEG, MPEG, e transmissão de TV digital
- **Fractals in Nature**: Padrões fractais aparecem em montanhas, costas marítimas, e plantas

## 💡 Curiosidades

- O conjunto de Mandelbrot foi descoberto por Benoit Mandelbrot em 1980
- Não importa o quanto você amplie, sempre verá novos detalhes (auto-similaridade)
- É infinitamente complexo, mas definido por uma fórmula extremamente simples
- Antenas fractais foram usadas no rover Perseverance da NASA em Marte!

---

## 👨‍💻 Autor

<div align="center">

### 🐹 Gopher Developer

**Duarte Rodrigo Santos de Oliveira**

*Estudante Autodidata da linguagem Go*

[![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/duarte-backend-golang)

</div>
