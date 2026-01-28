// Exercício 3.5 - Mandelbrot Colorido (RGBA)
// Este programa gera uma imagem colorida do conjunto de Mandelbrot usando o espaço de cores RGBA.
// O conjunto de Mandelbrot é um fractal matemático que mostra quais números complexos
// permanecem limitados quando iterados pela fórmula z = z² + c.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

// mandelbrot calcula se um número complexo z pertence ao conjunto de Mandelbrot
// e retorna uma cor RGBA baseada no número de iterações necessárias para divergir.
//
// Parâmetros:
//   z - número complexo a ser testado (representa um ponto no plano complexo)
//
// Retorna:
//   Uma cor que representa o comportamento do ponto:
//   - Cores vibrantes: pontos que escapam rapidamente (fora do conjunto)
//   - Preto: pontos que permanecem no conjunto (não escapam após 200 iterações)
func mandelbrot(z complex128) color.Color {
	// Número máximo de iterações antes de considerar que o ponto está no conjunto
	const iterations = 200
	
	// v representa o valor atual da iteração (começa em 0)
	var v complex128

	// Itera aplicando a fórmula z = z² + c
	for n := uint8(0); n < iterations; n++ {
		// Fórmula de Mandelbrot: v_novo = v_antigo² + z
		v = v*v + z
		
		// Se |v| > 2, sabemos que o ponto vai divergir para infinito
		// (propriedade matemática do conjunto de Mandelbrot)
		if cmplx.Abs(v) > 2 {
			// Gera cores vibrantes baseadas no número de iterações
			// Quanto menos iterações, mais longe do conjunto (cores mais variadas)
			return color.RGBA{
				R: 255 - n,     // Vermelho: diminui com mais iterações
				G: n * 2,       // Verde: aumenta rapidamente
				B: n * 3,       // Azul: aumenta ainda mais rápido
				A: 255,         // Alpha: totalmente opaco
			}
		}
	}
	
	// Se após todas as iterações ainda não divergiu, considera parte do conjunto
	return color.Black
}

func main() {
	// Define a região do plano complexo a ser renderizada
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2  // Limites do plano complexo
		width, height          = 1024, 1024       // Resolução da imagem em pixels
	)

	// Cria uma nova imagem RGBA (Red, Green, Blue, Alpha)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	
	// Percorre cada pixel da imagem
	for py := 0; py < height; py++ {
		// Converte a coordenada y do pixel para coordenada no plano complexo
		// Mapeia [0, height] para [ymin, ymax]
		y := float64(py)/height*(ymax-ymin) + ymin
		
		for px := 0; px < width; px++ {
			// Converte a coordenada x do pixel para coordenada no plano complexo
			// Mapeia [0, width] para [xmin, xmax]
			x := float64(px)/width*(xmax-xmin) + xmin
			
			// Cria um número complexo com as coordenadas calculadas
			z := complex(x, y)
			
			// Define a cor do pixel baseado no comportamento do ponto no Mandelbrot
			img.Set(px, py, mandelbrot(z))
		}
	}

	// Codifica a imagem em formato PNG e envia para a saída padrão (stdout)
	// Para salvar em arquivo, redirecione: go run mandelbrot_colorido.go > imagem.png
	png.Encode(os.Stdout, img)
}
