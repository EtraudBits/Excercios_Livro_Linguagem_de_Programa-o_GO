# Exercício 3.7 - Fractal do Método de Newton

## 📋 Descrição

Este exercício gera um fractal fascinante usando o **Método de Newton** para encontrar raízes da equação **z⁴ - 1 = 0** no plano complexo. O resultado mostra as "bacias de atração" de cada raiz, criando fronteiras fractais infinitamente complexas.

## 🎯 O que é o Método de Newton?

O **Método de Newton-Raphson** é um algoritmo iterativo para encontrar raízes (zeros) de funções.

### Fórmula Geral

```
z[n+1] = z[n] - f(z[n])/f'(z[n])
```

Onde:

- `f(z)` é a função que queremos encontrar as raízes
- `f'(z)` é a derivada da função
- Começamos com um palpite inicial `z[0]`
- Iteramos até convergir para uma raiz

### Intuição Geométrica

```
      f(z)
       │
       │    ╱
       │  ╱
       │╱← tangente
   ────┼──────────── z
       │    ↓
       │    raiz
```

Em cada iteração:

1. Calculamos a reta tangente à curva no ponto atual
2. Encontramos onde essa tangente cruza o eixo x
3. Esse cruzamento é nossa próxima aproximação

## 🧮 A Equação: z⁴ - 1 = 0

### As 4 Raízes Complexas

Queremos encontrar todos os números `z` tais que `z⁴ = 1`.

```
z⁴ = 1
```

**Raízes**:

1. **z = 1** (vermelho)
2. **z = -1** (verde)
3. **z = i** (azul)
4. **z = -i** (amarelo)

### Verificação

```
1⁴ = 1 ✓
(-1)⁴ = 1 ✓
i⁴ = (i²)² = (-1)² = 1 ✓
(-i)⁴ = ((-i)²)² = (-1)² = 1 ✓
```

### Visualização no Plano Complexo

```
        i (azul)
        │
        ●
        │
────●───┼───●──── Re
   -1   │   1
  (verde)│(vermelho)
        │
        ●
        │
       -i (amarelo)
```

As raízes formam um quadrado simétrico centrado na origem!

## 🎨 Como Funciona o Fractal

### 1. Bacias de Atração

Cada ponto do plano complexo é um **ponto inicial** diferente para o Método de Newton.

- Pontos iniciais próximos da raiz **1** tendem a convergir para 1 → **vermelho**
- Pontos iniciais próximos da raiz **-1** tendem a convergir para -1 → **verde**
- Pontos iniciais próximos da raiz **i** tendem a convergir para i → **azul**
- Pontos iniciais próximos da raiz **-i** tendem a convergir para -i → **amarelo**

### 2. Fronteiras Fractais

As **fronteiras** entre as bacias são fractais! Pontos muito próximos podem convergir para raízes diferentes, criando padrões infinitamente complexos.

```
Zoom nas fronteiras:
────────────         ───╱╲╱╲╱╲───        ╱─╲╱─╲╱─╲
vermelho│verde  →    ╱─╲││││╱─╲    →    ││││││││││
────────────         ───╲╱╲╱╲╱───        ╲─╱╲─╱╲─╱
```

Não importa o quanto você amplie, sempre verá novos detalhes!

### 3. Sombreamento por Velocidade

```
Cor mais clara = Convergência rápida (poucas iterações)
Cor mais escura = Convergência lenta (muitas iterações)
```

Isto adiciona profundidade e destaca as regiões de fronteira.

## 🚀 Como Usar

```bash
# Executar o programa
go run metodoNewton.go

# A imagem é salva automaticamente como:
# newton_fractal.png
```

## 📐 A Matemática Detalhada

### Aplicando Newton à Nossa Função

Para `f(z) = z⁴ - 1`:

1. **Derivada**: `f'(z) = 4z³`

2. **Fórmula de Newton**:

```
z[n+1] = z[n] - f(z[n])/f'(z[n])
z[n+1] = z[n] - (z[n]⁴ - 1)/(4z[n]³)
```

3. **Simplificação**:

```
z[n+1] = z[n] - (z⁴ - 1)/(4z³)
z[n+1] = (4z⁴ - z⁴ + 1)/(4z³)
z[n+1] = (3z⁴ + 1)/(4z³)
```

### Exemplo de Iteração

Começando com `z₀ = 0.5 + 0.5i`:

```
Iteração 0: z = 0.5 + 0.5i
Iteração 1: z = 0.7 + 0.3i     (aplicando fórmula)
Iteração 2: z = 0.85 + 0.15i
Iteração 3: z = 0.95 + 0.05i
Iteração 4: z = 0.99 + 0.01i
Iteração 5: z ≈ 1.0            (convergiu para 1!)
```

## 🌍 Aplicações na Vida Real

### 1. **Engenharia e Física**

**Análise de Circuitos Elétricos**:

```
Encontrar frequências de ressonância: Z(ω) = 0
Usar Newton para resolver equações não-lineares
```

**Otimização**:

```
Minimizar f(x) → Encontrar onde f'(x) = 0
Usar Newton na derivada
```

### 2. **Computação Gráfica**

**Ray Tracing**:

```
Encontrar interseções raio-superfície
Resolver: raio(t) = superfície(x,y,z)
```

**Iluminação**:

```
Calcular reflexos e refrações
Resolver equações de Fresnel
```

### 3. **Aprendizado de Máquina**

**Treinamento de Redes Neurais**:

```
Minimizar função de custo
Método de Newton de segunda ordem (quasi-Newton)
Algoritmos: BFGS, L-BFGS
```

### 4. **Finanças**

**Volatilidade Implícita (Black-Scholes)**:

```
Preço da opção = f(volatilidade)
Dado o preço, encontrar a volatilidade
Usar Newton: volatilidade é a raiz!
```

### 5. **Astronomia**

**Órbitas de Planetas**:

```
Equação de Kepler: M = E - e·sin(E)
Resolver para E usando Newton
Calcular posições planetárias
```

### 6. **Robótica**

**Cinemática Inversa**:

```
Dado: posição desejada do braço robótico
Encontrar: ângulos das juntas
Resolver sistema não-linear com Newton
```

## 🔬 Propriedades Fractais

### Auto-similaridade

```
Zoom 1x:    Zoom 100x:    Zoom 10000x:
  ●●●●         ●●●●           ●●●●
  ●○○●    →    ●○○●      →    ●○○●
  ●○○●         ●●●●           ●●●●
  ●●●●
```

Os mesmos padrões aparecem em todas as escalas!

### Dimensão Fractal

A fronteira entre as bacias tem **dimensão fractal** ≈ 2 (preenche o espaço, mas não é 3D).

### Conjunto de Julia

As fronteiras são relacionadas ao **Conjunto de Julia** para a função `z⁴ - 1`.

## 💻 Conceitos de Go Utilizados

### 1. **Arrays e Slices Globais**

```go
var roots = []complex128{...}     // Slice de números complexos
var rootColors = []color.RGBA{...} // Slice de cores
```

### 2. **Iteração sobre Arrays**

```go
for i, root := range roots {
    // i = índice, root = valor
}
```

### 3. **Aritmética Complexa**

```go
z*z*z*z        // z⁴
cmplx.Abs(z)   // |z| (módulo)
```

### 4. **Manipulação de Cores**

```go
shade := 255 - uint8(n*255/iterations)
r := uint8(int(rootColors[i].R) * int(shade) / 255)
```

## 🎓 O que Você Aprende

- ✅ Método de Newton-Raphson
- ✅ Raízes de polinômios complexos
- ✅ Bacias de atração
- ✅ Fractais e auto-similaridade
- ✅ Convergência de algoritmos iterativos
- ✅ Visualização de dados matemáticos
- ✅ Aritmética de números complexos

## 🎨 Experimentos Sugeridos

### 1. Outras Equações

**z³ - 1 = 0** (3 raízes, triângulo):

```go
z = z - (z*z*z - 1)/(3*z*z)
```

**z⁵ - 1 = 0** (5 raízes, pentágono):

```go
z = z - (z*z*z*z*z - 1)/(5*z*z*z*z)
```

**z² + 1 = 0** (2 raízes, i e -i):

```go
z = z - (z*z + 1)/(2*z)
```

### 2. Esquemas de Cores Diferentes

**Gradiente contínuo**:

```go
hue := float64(i) / float64(len(roots)) * 360
// Converter HSV para RGB
```

**Sem sombreamento** (cores sólidas):

```go
return rootColors[i]  // Sem multiplicar por shade
```

### 3. Zoom em Regiões Interessantes

**Fronteira entre vermelho e azul**:

```go
xmin, ymin, xmax, ymax = 0.5, -0.5, 1.5, 0.5
```

### 4. Mais Iterações

```go
const iterations = 100  // Mais detalhes nas fronteiras
```

### 5. Visualizar Não-Convergência

```go
if n == iterations {
    return color.White  // Pontos problemáticos em branco
}
```

## 📊 Comparação: Newton vs Mandelbrot

| Aspecto        | Mandelbrot              | Newton                     |
| -------------- | ----------------------- | -------------------------- |
| **Tipo**       | Conjunto de escape      | Bacias de atração          |
| **Pergunta**   | "Escapa para infinito?" | "Para qual raiz converge?" |
| **Cores**      | Velocidade de escape    | Qual raiz + velocidade     |
| **Estrutura**  | Um conjunto conectado   | Múltiplas regiões          |
| **Fronteiras** | Auto-similar            | Auto-similar               |

## 🏆 Vantagens do Método de Newton

### ✅ Convergência Quadrática

```
Erro[n+1] ≈ (Erro[n])²
```

Cada iteração **dobra** o número de dígitos corretos!

### ✅ Funciona para Qualquer Função

- Polinômios
- Transcendentais (sen, cos, exp)
- Sistemas de equações

### ✅ Base de Muitos Algoritmos

- Otimização não-linear
- Machine learning
- Simulação física

## ⚠️ Limitações

### ❌ Pode Divergir

Alguns pontos iniciais não convergem para nenhuma raiz.

### ❌ Precisa da Derivada

Nem sempre é fácil calcular `f'(z)`.

### ❌ Divisão por Zero

Se `f'(z) = 0`, o método falha.

### ❌ Convergência para Mínimo Local

Pode convergir para máximo/mínimo ao invés de raiz.

## 📚 Variações do Método

| Método           | Descrição                  | Uso                      |
| ---------------- | -------------------------- | ------------------------ |
| **Newton**       | Usa derivada exata         | Quando f'(z) é conhecida |
| **Secante**      | Aproxima derivada          | Quando f'(z) é difícil   |
| **Halley**       | Usa segunda derivada       | Convergência cúbica      |
| **Quasi-Newton** | Atualiza aproximação de f' | Otimização               |

## 🔗 Veja Também

- `../exercicio3.5/` - Mandelbrot colorido
- `../exercicio3.6/` - Super amostragem
- `../exercicio3.8/` - Zoom com precisão arbitrária

## 💡 Curiosidade Histórica

**Isaac Newton** desenvolveu este método por volta de **1669**, mas só foi publicado em 1711. **Joseph Raphson** publicou uma versão simplificada em 1690, por isso o nome "Newton-Raphson".

O método é tão poderoso que, mesmo depois de 350 anos, ainda é amplamente usado em:

- Calculadoras científicas (para calcular √x)
- GPUs (ray tracing)
- Inteligência Artificial (otimização)
- Engenharia (simulações)

## 🎭 Conexão com Arte

O fractal de Newton tem inspirado artistas digitais desde os anos 1980. As fronteiras complexas e simétricas criam padrões hipnóticos que aparecem em:

- Álbuns de música eletrônica
- Visualizações de planetários
- Arte generativa NFT
- Screensavers clássicos

---

## 👨‍💻 Autor

<div align="center">

### 🐹 Gopher Developer

**Duarte Rodrigo Santos de Oliveira**

*Estudante Autodidata da linguagem Go*

[![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/duarte-backend-golang)

</div>
