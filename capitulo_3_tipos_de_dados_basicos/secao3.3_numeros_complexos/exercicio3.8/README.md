# Exercício 3.8 - Precisão Numérica e Zoom Profundo em Fractais

## 📋 Descrição

Este exercício explora os **limites da precisão numérica** ao renderizar fractais com diferentes níveis de zoom. Compara 4 tipos numéricos diferentes em termos de desempenho, uso de memória e qualidade de imagem em zooms profundos.

## 🎯 O Problema da Precisão

### Por que Precisamos de Mais Precisão?

Fractais como o Mandelbrot têm **detalhes infinitos** em todas as escalas. Porém, números de ponto flutuante têm **precisão limitada**:

```
Zoom 1x (vista completa):
█████░░░░░  ← Detalhes visíveis
█████░░░░░
█████░░░░░

Zoom 1000x com float32:
█████████  ← Perdeu os detalhes!
█████████     (pixelation artificial)
█████████

Zoom 1000x com big.Float:
█████▓▒░░  ← Detalhes preservados
████▓▓▒░░     (fractal continua)
███▓▓▒▒░░
```

### O que Acontece?

Quando fazemos zoom, as coordenadas ficam muito próximas:

```python
# Sem zoom
x1 = -0.5
x2 = -0.6
diferença = 0.1 ✓ (distinguível)

# Zoom 10000x
x1 = -0.7453000001
x2 = -0.7453000002
diferença = 0.0000000001
```

Com float32 (7 dígitos), ambos viram `-0.7453` → **indistinguíveis!**

## 🔢 Os 4 Tipos Numéricos Testados

### 1. complex64 (32-bit)

```go
type complex64 = complex(float32, float32)
```

**Características**:

- **Precisão**: ~7 dígitos decimais
- **Tamanho**: 8 bytes (4 + 4)
- **Velocidade**: Muito rápida (nativa)
- **Limitação**: Zoom máximo ~10⁻³

**Quando usar**: Visualizações rápidas, baixa resolução

### 2. complex128 (64-bit)

```go
type complex128 = complex(float64, float64)
```

**Características**:

- **Precisão**: ~15 dígitos decimais
- **Tamanho**: 16 bytes (8 + 8)
- **Velocidade**: Muito rápida (nativa)
- **Limitação**: Zoom máximo ~10⁻¹²

**Quando usar**: Uso geral, 99% dos casos

### 3. big.Float (Precisão Arbitrária)

```go
import "math/big"
z := new(big.Float).SetPrec(256)  // 256 bits
```

**Características**:

- **Precisão**: Configurável (256 bits = ~77 dígitos)
- **Tamanho**: ~32 bytes + overhead
- **Velocidade**: ~50-100x mais lento
- **Limitação**: Zoom máximo ~10⁻⁵⁰

**Quando usar**: Zoom profundo, precisão crítica

### 4. big.Rat (Racional Ilimitado)

```go
import "math/big"
z := new(big.Rat)  // Numerador/Denominador
```

**Características**:

- **Precisão**: Ilimitada (teórica)
- **Tamanho**: Cresce com a precisão necessária
- **Velocidade**: ~200-500x mais lento
- **Limitação**: Apenas memória disponível

**Quando usar**: Demonstrações matemáticas, precisão perfeita

## 🚀 Como Usar

```bash
# Executar os testes
go run zoom.go

# O programa gera 16 imagens:
# - zoom1_complex64.png, zoom1_complex128.png, zoom1_bigfloat.png, zoom1_bigrat.png
# - zoom2_complex64.png, zoom2_complex128.png, zoom2_bigfloat.png, zoom2_bigrat.png
# - zoom3_complex64.png, zoom3_complex128.png, zoom3_bigfloat.png, zoom3_bigrat.png
# - zoom4_complex64.png, zoom4_complex128.png, zoom4_bigfloat.png, zoom4_bigrat.png

# E exibe estatísticas de desempenho no terminal
```

## 📊 Resultados Esperados

### Níveis de Zoom Testados

| Nível | Área    | Zoom   | O que Acontece                           |
| ----- | ------- | ------ | ---------------------------------------- |
| **1** | 4.0     | 1x     | Vista completa, todos funcionam bem      |
| **2** | 0.4     | 10x    | Região interessante, todos OK            |
| **3** | 0.001   | 4000x  | **complex64 falha** (pixelation visível) |
| **4** | 0.00005 | 80000x | **complex128 começa a falhar**           |

### Desempenho Típico (512×512 pixels)

```
=== Zoom Nível 1 (Vista Completa) ===
complex64:   120ms
complex128:  125ms
big.Float:   8.5s    (68x mais lento)
big.Rat:     45s     (360x mais lento)

=== Zoom Nível 4 (Zoom Profundo) ===
complex64:   110ms   (resultados inválidos!)
complex128:  120ms   (começando a falhar)
big.Float:   9.2s    (resultados perfeitos)
big.Rat:     50s     (resultados perfeitos, mas muito lento)
```

### Uso de Memória

| Tipo                | Memória por Número       | Para 512×512 Imagem |
| ------------------- | ------------------------ | ------------------- |
| complex64           | 8 bytes                  | ~2 MB               |
| complex128          | 16 bytes                 | ~4 MB               |
| big.Float (256-bit) | ~40 bytes                | ~10 MB              |
| big.Rat             | Variável (50-200+ bytes) | ~25-100 MB          |

## 🔍 Artefatos de Renderização

### Zoom Nível 1 - Todos Perfeitos

```
complex64:  ████▓▓▒▒░░
complex128: ████▓▓▒▒░░  ← Idênticos
big.Float:  ████▓▓▒▒░░
big.Rat:    ████▓▓▒▒░░
```

### Zoom Nível 3 - complex64 Falha

```
complex64:  ██████████  ← Pixelation!
complex128: ████▓▓▒▒░░  ← Ainda OK
big.Float:  ████▓▓▒▒░░
big.Rat:    ████▓▓▒▒░░
```

### Zoom Nível 4 - complex128 Falha

```
complex64:  ██████████  ← Completamente errado
complex128: ████▓▓████  ← Começando a pixelar
big.Float:  ████▓▓▒▒░░  ← Perfeito
big.Rat:    ████▓▓▒▒░░  ← Perfeito (mas demorou!)
```

## 🧮 A Matemática da Precisão

### Float32 (complex64)

```
Bits: 1 (sinal) + 8 (expoente) + 23 (mantissa) = 32 bits
Precisão: 2^23 ≈ 8 milhões de valores
Dígitos decimais: log₁₀(2^23) ≈ 6.9 ≈ 7 dígitos
```

**Exemplo**:

```
1.234567  ✓ Representável
1.2345678 ✗ Arredondado para 1.234568
```

### Float64 (complex128)

```
Bits: 1 (sinal) + 11 (expoente) + 52 (mantissa) = 64 bits
Precisão: 2^52 ≈ 4.5 quatrilhões de valores
Dígitos decimais: log₁₀(2^52) ≈ 15.6 ≈ 15 dígitos
```

### big.Float

```
Precisão configurável:
- 53 bits = float64
- 113 bits = quad precision
- 256 bits = ~77 dígitos decimais (usado neste código)
- 1024 bits = ~308 dígitos (se necessário)
```

### big.Rat

```
Representação: numerador/denominador
Exemplo: 1/3 = exatamente 1/3 (não 0.333333...)

Vantagem: Precisão perfeita para números racionais
Desvantagem: Denominadores crescem exponencialmente
  Iteração 1: 1/1000
  Iteração 10: 1/1000000000000 (trilhão!)
  Memória explode rapidamente
```

## 🌍 Aplicações na Vida Real

### 1. **Criptografia**

```go
// RSA usa big.Int para números enormes
p := big.NewInt(...)  // Primo gigante
q := big.NewInt(...)  // Outro primo
n := p * q            // Produto (chave pública)
```

Precisão perfeita é **crítica** - um erro destrói a segurança!

### 2. **Finanças**

```go
// Dinheiro requer precisão exata
valor := big.NewRat(123456789, 100)  // $1,234,567.89
```

Float causaria erros de arredondamento:

```
0.1 + 0.2 = 0.30000000000000004  // float64
1/10 + 2/10 = 3/10               // big.Rat (exato!)
```

### 3. **Astronomia**

Distâncias astronômicas exigem alta precisão:

```
Distância Terra-Lua: 384,400 km
Erro de 1 mm: requer 11 dígitos de precisão
float64 (15 dígitos) é suficiente ✓
```

### 4. **Computação Simbólica**

Sistemas como Mathematica/Maple usam precisão arbitrária internamente:

```mathematica
N[Pi, 1000]  (* 1000 dígitos de π *)
```

### 5. **Simulação Física**

**Clima/Meteorologia**:

- Sistemas caóticos amplificam pequenos erros
- Precisão maior = previsões mais longas
- Mas: custo computacional aumenta muito

**N-Corpos** (galáxias):

- Simular órbitas por milhões de anos
- Erros acumulam ao longo do tempo
- Precisão maior necessária

### 6. **Arte Generativa**

Explorar fractais em zoom profundo para criar arte:

```
Zoom 10^50: Requer big.Float
Zoom 10^100: Requer big.Rat ou precisão customizada
```

## 💻 Implementação em Go

### complex64/128 - Nativo

```go
z := complex(1.5, 2.3)  // Simples!
z = z*z + c
```

### big.Float - Manual

```go
// Criar
real := new(big.Float).SetPrec(256)
imag := new(big.Float).SetPrec(256)

// Multiplicação complexa: (a+bi)²
realSq := new(big.Float).Mul(real, real)
imagSq := new(big.Float).Mul(imag, imag)
// ... muito código manual
```

### big.Rat - Ainda Mais Manual

```go
// Cada operação cria novos objetos
real := new(big.Rat)
imag := new(big.Rat)
temp := new(big.Rat).Mul(real, imag)
// ... denominadores crescem rapidamente
```

## 🎓 O que Você Aprende

- ✅ Limites da precisão de ponto flutuante
- ✅ Trade-off entre velocidade e precisão
- ✅ Quando usar cada tipo numérico
- ✅ Como artefatos de renderização surgem
- ✅ Aritmética de precisão arbitrária
- ✅ Medição de desempenho em Go
- ✅ Gerenciamento de memória

## 🔬 Experimentos Sugeridos

### 1. Zoom Extremo

Teste até onde cada tipo aguenta:

```go
// Zoom 10^20
xmin, xmax = -0.74530000000000000001, -0.74529999999999999999
```

### 2. Comparação Visual

Gere imagens lado a lado:

```go
// Metade complex128, metade big.Float
if px < width/2 {
    use complex128
} else {
    use big.Float
}
```

### 3. Precisões Diferentes de big.Float

```go
precisions := []uint{53, 113, 256, 512, 1024}
for _, prec := range precisions {
    // Testar cada precisão
}
```

### 4. Análise de Erro

Compare pixel a pixel:

```go
diff := abs(color64 - color128)
maxDiff := max(diff)  // Onde está o maior erro?
```

### 5. Benchmark Detalhado

```go
func BenchmarkMandelbrot64(b *testing.B) {
    for i := 0; i < b.N; i++ {
        mandelbrot64(complex(0.5, 0.5))
    }
}
```

## 📚 Conceitos Avançados

### Epsilon de Máquina

O menor número que, somado a 1, dá resultado diferente de 1:

```go
// float32
epsilon32 := float32(1.1920929e-07)  // ~1.2e-7

// float64
epsilon64 := float64(2.220446049250313e-16)  // ~2.2e-16
```

### Erro de Cancelamento

```go
a := 1.0000001
b := 1.0000000
c := a - b  // Perde precisão!
// Apenas 1 dígito significativo resta
```

### Acúmulo de Erro

```go
sum := 0.0
for i := 0; i < 10000000; i++ {
    sum += 0.1  // Erro acumula!
}
// sum != 1000000.0 (esperado)
// sum ≈ 1000000.0000001
```

## 🏆 Recomendações Práticas

| Situação               | Tipo Recomendado                | Razão               |
| ---------------------- | ------------------------------- | ------------------- |
| Visualização geral     | complex128                      | Equilíbrio perfeito |
| Zoom até 10⁻³          | complex128                      | Suficiente          |
| Zoom 10⁻³ a 10⁻¹²      | complex128                      | Ainda funciona      |
| Zoom 10⁻¹² a 10⁻⁵⁰     | big.Float (256-bit)             | Necessário          |
| Zoom > 10⁻⁵⁰           | big.Float (512+ bit) ou big.Rat | Extremo             |
| Dinheiro/Contabilidade | big.Rat ou decimal              | Precisão exata      |
| Jogos/Gráficos RT      | complex128 ou complex64         | Velocidade          |
| Criptografia           | big.Int                         | Segurança           |

## ⚠️ Armadilhas Comuns

### 1. Achar que float64 é "suficiente"

```go
// Zoom profundo com float64
x1 := -0.7453000000000001
x2 := -0.7453000000000002
// Podem ser iguais! (apenas 15 dígitos)
```

### 2. Usar big.Rat desnecessariamente

```go
// Desperdiça memória e tempo
// Use apenas quando realmente precisa precisão ilimitada
```

### 3. Não medir desempenho

```go
// Sempre meça antes de otimizar!
start := time.Now()
// ... código
fmt.Printf("Levou %v\n", time.Since(start))
```

### 4. Esquecer de configurar precisão

```go
x := new(big.Float)  // Precisão padrão (53 bits)!
x.SetPrec(256)        // Agora sim: 256 bits
```

## 🔗 Veja Também

- `../exercicio3.5/` - Mandelbrot colorido
- `../exercicio3.6/` - Anti-aliasing
- `../exercicio3.7/` - Método de Newton
- Go docs: `math/big` package

## 💡 Citação Famosa

> "Floating point arithmetic is like moving piles of sand; every time you move them, you lose a little sand and you pick up a little dirt."
>
> — **Anônimo**

A moral: sempre tenha consciência das limitações de precisão dos números que você usa!
