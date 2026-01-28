// Programa Básico - Conjunto de Mandelbrot em Escala de Cinza
//
// Este é o programa introdutório que demonstra como gerar o famoso conjunto de Mandelbrot
// usando números complexos em Go. É a base para todos os exercícios posteriores desta seção.
//
// O QUE É O CONJUNTO DE MANDELBROT?
// É um conjunto de pontos no plano complexo que permanecem "limitados" (não vão para infinito)
// quando iteramos a fórmula: z = z² + c
//
// DESCOBERTO POR:
// Benoit Mandelbrot em 1980, embora a matemática seja anterior (Pierre Fatou e Gaston Julia, 1917)
//
// POR QUE É ESPECIAL?
// - Estrutura infinitamente complexa a partir de uma fórmula simples
// - Auto-similar: zoom infinito revela sempre novos detalhes
// - Um dos exemplos mais famosos de fractal na matemática
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

// mandelbrot determina se um número complexo z pertence ao conjunto de Mandelbrot.
//
// ALGORITMO:
// 1. Começa com v = 0
// 2. Itera: v = v² + z
// 3. Se |v| > 2, o ponto "escapa" (não está no conjunto)
// 4. Se após muitas iterações |v| ≤ 2, provavelmente está no conjunto
//
// COLORAÇÃO:
// - Pontos no conjunto: preto
// - Pontos fora: cinza baseado na velocidade de escape
//   - Escapa rápido (poucas iterações): cinza claro
//   - Escapa devagar (muitas iterações): cinza escuro
//
// Parâmetros:
//   z - número complexo a ser testado (representa um ponto no plano complexo)
//
// Retorna:
//   color.Color - cinza se escapa, preto se permanece no conjunto
func mandelbrot(z complex128) color.Color {
	const iterations = 200  // Número máximo de iterações antes de considerar "no conjunto"
	const contrast = 15     // Contraste da escala de cinza (quanto maior, mais contraste)

	var v complex128  // Valor atual da iteração (começa em 0)

	// Itera aplicando a fórmula de Mandelbrot
	for n := uint8(0); n < iterations; n++ {
		// Fórmula principal: z[n+1] = z[n]² + c
		// onde 'z' (parâmetro) é o 'c' da fórmula, e 'v' é o z iterado
		v = v*v + z
		
		// Teste de escape: se |v| > 2, sabemos que vai divergir para infinito
		// Esta é uma propriedade matemática comprovada do conjunto de Mandelbrot
		if cmplx.Abs(v) > 2 {
			// O ponto escapou! Retorna um cinza baseado em quantas iterações levou
			// Menos iterações (n pequeno) → cinza mais claro
			// Mais iterações (n grande) → cinza mais escuro
			return color.Gray{255 - contrast*n}
		}
	}
	
	// Após todas as iterações, o ponto ainda não escapou
	// Consideramos que ele está no conjunto de Mandelbrot
	return color.Black
}

func main() {
	// Define a região do plano complexo a ser renderizada
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2  // Limites: de -2-2i a +2+2i
		width, height          = 1024, 1024       // Resolução da imagem (1024×1024 pixels)
	)

	// Cria uma nova imagem RGBA (Red, Green, Blue, Alpha)
	// Mesmo usando escala de cinza, Go usa RGBA internamente
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	
	// Loop duplo: percorre cada pixel da imagem
	for py := 0; py < height; py++ {
		// Converte a coordenada y do pixel para coordenada no plano complexo
		// Mapeia [0, height-1] para [ymin, ymax]
		// Exemplo: py=0 → y=ymin, py=height-1 → y≈ymax
		y := float64(py)/height*(ymax-ymin) + ymin
		
		for px := 0; px < width; px++ {
			// Converte a coordenada x do pixel para coordenada no plano complexo
			// Mapeia [0, width-1] para [xmin, xmax]
			x := float64(px)/width*(xmax-xmin) + xmin
			
			// Cria um número complexo z = x + yi
			// Este ponto no plano complexo será testado
			z := complex(x, y)
			
			// Define a cor do pixel baseado no teste de Mandelbrot
			img.Set(px, py, mandelbrot(z))
		}
	}

	// Codifica a imagem em formato PNG e envia para a saída padrão (stdout)
	// Para salvar em arquivo, redirecione a saída:
	// go run mandelbrot.go > mandelbrot.png
	png.Encode(os.Stdout, img)
}
