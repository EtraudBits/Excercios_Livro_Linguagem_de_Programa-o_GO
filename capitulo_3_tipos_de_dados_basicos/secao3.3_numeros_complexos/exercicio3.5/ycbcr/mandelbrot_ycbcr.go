// Exercício 3.5 - Mandelbrot com YCbCr
// Este programa gera uma imagem colorida do conjunto de Mandelbrot usando o espaço de cores YCbCr.
// YCbCr é usado em sistemas de vídeo (JPEG, MPEG) porque separa luminância (brilho) de crominância (cor),
// permitindo compressão mais eficiente já que o olho humano é mais sensível a mudanças de brilho.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

// mandelbrot calcula se um número complexo z pertence ao conjunto de Mandelbrot
// e retorna uma cor YCbCr baseada no número de iterações necessárias para divergir.
//
// YCbCr é um espaço de cores que separa:
// - Y: Luminância (brilho) - informação em preto e branco
// - Cb: Crominância azul-amarelo (componente de diferença azul)
// - Cr: Crominância vermelho-verde (componente de diferença vermelho)
//
// Parâmetros:
//   z - número complexo a ser testado (representa um ponto no plano complexo)
//
// Retorna:
//   Uma cor YCbCr que representa o comportamento do ponto
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
		if cmplx.Abs(v) > 2 {
			// Gera cores usando o espaço YCbCr
			// Este espaço é usado em compressão de vídeo e imagens (JPEG, MPEG)
			return color.YCbCr{
				Y:  255 - n,        // Luminância: pontos mais distantes são mais escuros
				Cb: 128 + n/2,      // Componente azul: aumenta suavemente (128 é neutro)
				Cr: 128 + n,        // Componente vermelho: aumenta mais rapidamente
			}
			// Nota: 128 é o valor neutro para Cb e Cr (sem cor adicional)
			// Valores < 128 puxam para azul/verde, > 128 puxam para amarelo/vermelho
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

	// Cria uma nova imagem RGBA (embora usemos YCbCr, Go converte automaticamente)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	
	// Percorre cada pixel da imagem
	for py := 0; py < height; py++ {
		// Converte a coordenada y do pixel para coordenada no plano complexo
		y := float64(py)/height*(ymax-ymin) + ymin
		
		for px := 0; px < width; px++ {
			// Converte a coordenada x do pixel para coordenada no plano complexo
			x := float64(px)/width*(xmax-xmin) + xmin
			
			// Cria um número complexo com as coordenadas calculadas
			z := complex(x, y)
			
			// Define a cor do pixel usando YCbCr (será automaticamente convertido para RGBA)
			img.Set(px, py, mandelbrot(z))
		}
	}

	// Codifica a imagem em formato PNG e envia para a saída padrão
	// Para salvar em arquivo, redirecione: go run mandelbrot_ycbcr.go > imagem.png
	png.Encode(os.Stdout, img)
}
