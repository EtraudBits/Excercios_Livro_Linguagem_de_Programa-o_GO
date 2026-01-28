# Seção 3.3 - Números Complexos em Go

## 📚 Visão Geral da Seção

Esta seção explora o uso de **números complexos** em Go através da renderização de fractais matemáticos, especificamente o conjunto de Mandelbrot e variações. Os exercícios progressivamente introduzem conceitos avançados de computação gráfica, precisão numérica e otimização.

## 📖 Conteúdo

### [progMandelbrot/](progMandelbrot/) - Programa Base

**Conceito**: Implementação básica do conjunto de Mandelbrot em escala de cinza  
**Aprende**:

- Usar `complex128` em Go
- Operações com números complexos
- Geração de imagens PNG
- Iteração numérica

### [exercicio3.5/](exercicio3.5/) - Mandelbrot Colorido

**Conceito**: Aplicar diferentes espaços de cores (RGBA e YCbCr)  
**Aprende**: Espaço de cores RGBA, YCbCr, diferença entre representações de cor

### [exercicio3.6/](exercicio3.6/) - Super Amostragem (Anti-Aliasing)

**Conceito**: Eliminar serrilhado usando múltiplas amostras por pixel  
**Aprende**: Anti-aliasing, trade-off qualidade vs. performance

### [exercicio3.7/](exercicio3.7/) - Fractal do Método de Newton

**Conceito**: Visualizar bacias de atração de raízes usando Método de Newton  
**Aprende**: Método Newton-Raphson, bacias de atração, fronteiras fractais

### [exercicio3.8/](exercicio3.8/) - Precisão Numérica e Zoom Profundo

**Conceito**: Comparar 4 tipos numéricos em diferentes níveis de zoom  
**Aprende**: Limites de precisão, `big.Float`, `big.Rat`, artefatos de renderização

## 🧮 Números Complexos em Go

```go
// Criar número complexo
z := complex(3.0, 4.0)  // 3 + 4i

// Componentes
real := real(z)    // 3.0
imag := imag(z)    // 4.0

// Operações
z1 + z2        // soma
z1 * z2        // multiplicação

// Funções do pacote cmplx
import "math/cmplx"
cmplx.Abs(z)    // módulo
cmplx.Phase(z)  // ângulo
```

## 🚀 Começando

```bash
# Programa básico
cd progMandelbrot
go run mandelbrot.go > mandelbrot.png

# Colorido
cd ../exercicio3.5/rgba
go run mandelbrot_colorido.go > colorido.png

# Com cada exercício, veja o README.md na pasta correspondente
# para documentação completa e exemplos
```

Cada pasta contém um **README.md detalhado** com:

- Explicações didáticas
- Exemplos de uso
- Aplicações na vida real
- Conceitos matemáticos
- Experimentos sugeridos

**Divirta-se explorando fractais! 🚀✨**

---

## 👨‍💻 Autor

<div align="center">

### 🐹 Gopher Developer

**Duarte Rodrigo Santos de Oliveira**

_Estudante Autodidata da linguagem Go_

[![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/duarte-backend-golang)

</div>
