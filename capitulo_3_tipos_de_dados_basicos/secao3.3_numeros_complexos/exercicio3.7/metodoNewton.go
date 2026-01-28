// Exercício 3.7 - Fractal do Método de Newton
// Este programa gera um fractal usando o Método de Newton para encontrar raízes da equação z⁴ - 1 = 0.
//
// O QUE É O MÉTODO DE NEWTON?
// É um algoritmo iterativo para encontrar raízes (zeros) de funções. A ideia é:
// 1. Começar com um palpite inicial (z₀)
// 2. Melhorar o palpite usando: z[n+1] = z[n] - f(z[n])/f'(z[n])
// 3. Repetir até convergir para uma raiz
//
// POR QUE É UM FRACTAL?
// Diferentes pontos iniciais convergem para diferentes raízes, criando "bacias de atração".
// As fronteiras entre essas bacias formam padrões fractais infinitamente complexos!
//
// A EQUAÇÃO: z⁴ - 1 = 0
// Esta equação tem 4 raízes complexas:
// - 1 (real positivo)
// - -1 (real negativo)
// - i (imaginário positivo)
// - -i (imaginário negativo)
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

// As quatro raízes de z⁴ - 1 = 0
// Estas são as "soluções" ou "zeros" da função f(z) = z⁴ - 1
var roots = []complex128{
	complex(1, 0),   // 1: raiz real positiva
	complex(-1, 0),  // -1: raiz real negativa
	complex(0, 1),   // i: raiz imaginária positiva
	complex(0, -1),  // -i: raiz imaginária negativa
}

// Cores associadas a cada raiz (para visualização)
// Cada "bacia de atração" será pintada com uma cor diferente
var rootColors = []color.RGBA{
	{255, 0, 0, 255},    // Vermelho para a raiz 1
	{0, 255, 0, 255},    // Verde para a raiz -1
	{0, 0, 255, 255},    // Azul para a raiz i
	{255, 255, 0, 255},  // Amarelo para a raiz -i
}

// newton aplica o Método de Newton para encontrar qual raiz um ponto inicial converge.
//
// MÉTODO DE NEWTON: z_novo = z - f(z)/f'(z)
// Para f(z) = z⁴ - 1:
//   f(z) = z⁴ - 1
//   f'(z) = 4z³  (derivada)
//   Logo: z_novo = z - (z⁴ - 1)/(4z³)
//
// Parâmetros:
//   z - ponto inicial no plano complexo
//
// Retorna:
//   Uma cor que indica:
//   - Qual raiz o ponto converge (cor base)
//   - Quão rápido converge (tom mais claro = convergência mais rápida)
//   - Preto se não convergiu
func newton(z complex128) color.Color {
	const iterations = 50    // Máximo de iterações antes de desistir
	const tolerance = 1e-6   // Distância para considerar que chegou na raiz

	// Itera aplicando o Método de Newton
	for n := 0; n < iterations; n++ {
		// Fórmula de Newton para f(z) = z⁴ - 1:
		// z_novo = z - (z⁴ - 1)/(4z³)
		//
		// Passo a passo:
		// 1. Calcula z⁴: z*z*z*z
		// 2. Calcula f(z): z⁴ - 1
		// 3. Calcula f'(z): 4z³ = 4*z*z*z
		// 4. Divide: f(z)/f'(z)
		// 5. Atualiza: z = z - f(z)/f'(z)
		z = z - (z*z*z*z-1)/(4*z*z*z)

		// Verifica se chegamos perto de alguma raiz
		for i, root := range roots {
			// Calcula a distância entre z atual e a raiz
			if cmplx.Abs(z-root) < tolerance {
				// CONVERGIU para a raiz i!
				
				// Sombreia baseado no número de iterações:
				// - Menos iterações (n pequeno) → shade alto → cor mais clara
				// - Mais iterações (n grande) → shade baixo → cor mais escura
				// Isto cria um efeito de "profundidade" no fractal
				shade := 255 - uint8(n*255/iterations)
				
				// Aplica o sombreamento a cada componente RGB
				// Multiplicamos a cor base pelo fator de sombreamento
				r := uint8(int(rootColors[i].R) * int(shade) / 255)
				g := uint8(int(rootColors[i].G) * int(shade) / 255)
				b := uint8(int(rootColors[i].B) * int(shade) / 255)
				
				return color.RGBA{r, g, b, 255}
			}
		}
	}

	// Se após todas as iterações não convergiu para nenhuma raiz,
	// retorna preto (pontos "problemáticos" ou muito lentos)
	return color.Black
}

func main() {
	// Define a região do plano complexo a ser renderizada
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2  // Região quadrada centrada na origem
		width, height          = 1024, 1024       // Resolução da imagem
	)

	// Cria a imagem de saída
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	
	// Percorre cada pixel da imagem
	for py := 0; py < height; py++ {
		// Converte a coordenada y do pixel para o plano complexo
		y := float64(py)/height*(ymax-ymin) + ymin
		
		for px := 0; px < width; px++ {
			// Converte a coordenada x do pixel para o plano complexo
			x := float64(px)/width*(xmax-xmin) + xmin
			
			// Cria o número complexo z = x + yi
			// Este é o "ponto inicial" para o Método de Newton
			z := complex(x, y)
			
			// Aplica o Método de Newton e obtém a cor
			// A cor indica para qual raiz este ponto inicial converge
			img.Set(px, py, newton(z))
		}
	}

	// Salva a imagem em arquivo PNG
	file, err := os.Create("newton_fractal.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Codifica e salva
	if err := png.Encode(file, img); err != nil {
		panic(err)
	}
}
