// Exercício 3.6 - Super Amostragem (Supersampling / Anti-aliasing)
// Este programa implementa a técnica de super amostragem para reduzir o efeito de pixelation
// (serrilhado/aliasing) nas bordas do conjunto de Mandelbrot.
//
// ANTI-ALIASING: O problema
// Quando renderizamos uma imagem pixel por pixel, as bordas ficam serrilhadas (como escadas)
// porque cada pixel representa apenas UM ponto do plano complexo. Se esse ponto está dentro
// ou fora do conjunto, o pixel inteiro fica preto ou colorido, criando bordas abruptas.
//
// SUPER AMOSTRAGEM: A solução
// Em vez de testar apenas 1 ponto por pixel, testamos múltiplos pontos (subpixels) e
// calculamos a média das cores. Isso suaviza as transições e elimina o efeito de serrilhado.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

// mandelbrot calcula se um número complexo pertence ao conjunto de Mandelbrot
// Retorna uma cor em escala de cinza baseada na velocidade de escape
func mandelbrot(z complex128) color.Color {
	const iterations = 200  // Máximo de iterações
	const contrast = 15     // Contraste da imagem

	var v complex128  // Valor atual da iteração

	for n := uint8(0); n < iterations; n++ {
		// Aplica a fórmula de Mandelbrot: z = z² + c
		v = v*v + z
		
		// Se |v| > 2, o ponto vai divergir para infinito
		if cmplx.Abs(v) > 2 {
			// Retorna uma cor baseada na velocidade de escape
			return color.Gray{255 - contrast*n}
		}
	}
	
	// Ponto está no conjunto de Mandelbrot
	return color.Black
}

// superSample calcula a cor de um pixel usando super amostragem 2x2
// 
// TÉCNICA: Divide cada pixel em 4 subpixels (grade 2x2) e calcula a cor de cada um.
// A cor final do pixel é a MÉDIA das 4 cores obtidas.
//
// Parâmetros:
//   px, py - coordenadas do pixel na imagem
//   width, height - dimensões da imagem
//   xmin, ymin, xmax, ymax - limites do plano complexo
//
// Retorna:
//   A cor média dos 4 subpixels (suavizada)
//
// Exemplo visual de um pixel dividido:
//   +-------+-------+
//   | 0.25  | 0.75  |  <- offsets dx
//   | 0.25  | 0.25  |  <- offset dy = 0.25
//   +-------+-------+
//   | 0.25  | 0.75  |
//   | 0.75  | 0.75  |  <- offset dy = 0.75
//   +-------+-------+
//
// Resultado: 4 amostras por pixel ao invés de apenas 1!
func superSample(px, py int, width, height int, xmin, ymin, xmax, ymax float64) color.Color {
	// Acumuladores para os componentes RGBA (usamos uint32 para evitar overflow)
	var r, g, b, a uint32
	
	// Offsets para os 4 subpixels (posicionados nos quartis do pixel)
	// 0.25 = primeiro quartil, 0.75 = terceiro quartil
	// Isso distribui as amostras uniformemente pelo pixel
	offsets := []float64{0.25, 0.75}
	
	// Loop duplo: 2x2 = 4 subpixels
	for _, dy := range offsets {
		for _, dx := range offsets {
			// Calcula a posição exata deste subpixel no plano complexo
			// (float64(px)+dx) adiciona o offset ao pixel
			x := (float64(px)+dx)/float64(width)*(xmax-xmin) + xmin
			y := (float64(py)+dy)/float64(height)*(ymax-ymin) + ymin
			z := complex(x, y)
			
			// Calcula a cor deste subpixel específico
			c := mandelbrot(z)
			
			// Converte para RGBA e acumula os valores
			// RGBA() retorna valores de 0-65535 (16 bits por canal)
			r1, g1, b1, a1 := c.RGBA()
			r += r1
			g += g1
			b += b1
			a += a1
		}
	}
	
	// Retorna a MÉDIA das 4 amostras (divide por 4)
	// Isso suaviza as transições entre cores
	return color.RGBA64{
		R: uint16(r / 4),  // Média do vermelho
		G: uint16(g / 4),  // Média do verde
		B: uint16(b / 4),  // Média do azul
		A: uint16(a / 4),  // Média do alpha
	}
}

func main() {
	// Define a região do plano complexo e resolução da imagem
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2  // Área completa do Mandelbrot
		width, height          = 1024, 1024       // Resolução 1024x1024
	)

	// Cria a imagem de saída
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	
	// Percorre cada pixel da imagem
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			// Usa super amostragem ao invés de calcular apenas um ponto
			// Isso elimina o efeito de serrilhado (aliasing)
			img.Set(px, py, superSample(px, py, width, height, xmin, ymin, xmax, ymax))
		}
	}

	// Salva a imagem diretamente em um arquivo
	file, err := os.Create("mandelbrot_supersampled.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Codifica a imagem no formato PNG
	if err := png.Encode(file, img); err != nil {
		panic(err)
	}
}
