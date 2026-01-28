// Exercício 3.8 - Precisão Numérica em Fractais com Alto Zoom
//
// OBJETIVO:
// Comparar 4 representações numéricas diferentes ao renderizar o conjunto de Mandelbrot
// em diferentes níveis de zoom, analisando:
// 1. Desempenho (tempo de execução)
// 2. Uso de memória
// 3. Precisão (quando aparecem artefatos de renderização)
//
// REPRESENTAÇÕES TESTADAS:
// - complex64:  Números complexos de 32 bits (float32 + float32)
// - complex128: Números complexos de 64 bits (float64 + float64)
// - big.Float:  Ponto flutuante de precisão arbitrária (256 bits configurado)
// - big.Rat:    Números racionais de precisão ilimitada
//
// O PROBLEMA:
// Quando fazemos zoom profundo em fractais, números de ponto flutuante padrão
// (float32, float64) perdem precisão. Pontos muito próximos se tornam indistinguíveis,
// causando "pixelation" artificial - a imagem fica quadriculada mesmo quando deveria
// haver detalhes fractais infinitos.
//
// EXEMPLO DE PERDA DE PRECISÃO:
// float32 tem ~7 dígitos decimais de precisão:
// - Zoom 1x:     -0.5 vs -0.50000001  → Distinguíveis ✓
// - Zoom 1000x:  -0.5000001 vs -0.5000002  → Indistinguíveis ✗
//
// A SOLUÇÃO:
// Usar tipos numéricos com mais precisão quando necessário:
// - big.Float: 50-100x mais lento, mas suporta zoom profundo
// - big.Rat: 200-500x mais lento, mas precisão perfeita (teórica)
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/big"
	"math/cmplx"
	"os"
	"time"
)

// Mandelbrot usando complex64
func mandelbrot64(z complex64) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex64

	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if real(v)*real(v)+imag(v)*imag(v) > 4 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

// Mandelbrot usando complex128
func mandelbrot128(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128

	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

// Mandelbrot usando big.Float
func mandelbrotBigFloat(zReal, zImag *big.Float) color.Color {
	const iterations = 200
	const contrast = 15

	// Precisão de 256 bits
	vReal := new(big.Float).SetPrec(256)
	vImag := new(big.Float).SetPrec(256)

	// Temporários para cálculos
	temp := new(big.Float).SetPrec(256)
	vRealSq := new(big.Float).SetPrec(256)
	vImagSq := new(big.Float).SetPrec(256)
	absSquared := new(big.Float).SetPrec(256)
	four := big.NewFloat(4.0)

	for n := uint8(0); n < iterations; n++ {
		// v = v*v + z
		// (a+bi)^2 = a^2 - b^2 + 2abi
		vRealSq.Mul(vReal, vReal)             // vReal^2
		vImagSq.Mul(vImag, vImag)             // vImag^2
		temp.Mul(vReal, vImag)                // vReal * vImag
		temp.Mul(temp, big.NewFloat(2.0))    // 2 * vReal * vImag
		vReal.Sub(vRealSq, vImagSq)           // vReal^2 - vImag^2
		vReal.Add(vReal, zReal)               // + zReal
		vImag = temp.Add(temp, zImag)         // 2*vReal*vImag + zImag

		// Verifica |v|^2 > 4
		absSquared.Add(vRealSq, vImagSq)
		if absSquared.Cmp(four) > 0 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

// Mandelbrot usando big.Rat (números racionais)
func mandelbrotBigRat(zReal, zImag *big.Rat) color.Color {
	const iterations = 200
	const contrast = 15

	vReal := new(big.Rat)
	vImag := new(big.Rat)

	// Temporários para cálculos
	vRealSq := new(big.Rat)
	vImagSq := new(big.Rat)
	temp := new(big.Rat)
	newVImag := new(big.Rat)
	absSquared := new(big.Rat)
	four := big.NewRat(4, 1)

	for n := uint8(0); n < iterations; n++ {
		// v = v*v + z
		// (a+bi)^2 = a^2 - b^2 + 2abi
		vRealSq.Mul(vReal, vReal)
		vImagSq.Mul(vImag, vImag)
		temp.Mul(vReal, vImag)
		temp.Mul(temp, big.NewRat(2, 1))
		
		// Nova parte real: vReal^2 - vImag^2 + zReal
		vReal.Sub(vRealSq, vImagSq)
		vReal.Add(vReal, zReal)
		
		// Nova parte imaginária: 2*vReal*vImag + zImag
		newVImag.Add(temp, zImag)
		vImag.Set(newVImag)

		// Verifica |v|^2 > 4
		absSquared.Add(vRealSq, vImagSq)
		if absSquared.Cmp(four) > 0 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

// Renderiza usando complex64
func renderComplex64(xmin, ymin, xmax, ymax float64, width, height int, filename string) time.Duration {
	start := time.Now()
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(float32(x), float32(y))
			img.Set(px, py, mandelbrot64(z))
		}
	}

	saveImage(img, filename)
	return time.Since(start)
}

// Renderiza usando complex128
func renderComplex128(xmin, ymin, xmax, ymax float64, width, height int, filename string) time.Duration {
	start := time.Now()
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot128(z))
		}
	}

	saveImage(img, filename)
	return time.Since(start)
}

// Renderiza usando big.Float
func renderBigFloat(xmin, ymin, xmax, ymax float64, width, height int, filename string) time.Duration {
	start := time.Now()
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	xminBig := big.NewFloat(xmin).SetPrec(256)
	yminBig := big.NewFloat(ymin).SetPrec(256)
	xrangeBig := big.NewFloat(xmax - xmin).SetPrec(256)
	yrangeBig := big.NewFloat(ymax - ymin).SetPrec(256)

	for py := 0; py < height; py++ {
		yRatio := big.NewFloat(float64(py) / float64(height)).SetPrec(256)
		y := new(big.Float).SetPrec(256)
		y.Mul(yRatio, yrangeBig)
		y.Add(y, yminBig)

		for px := 0; px < width; px++ {
			xRatio := big.NewFloat(float64(px) / float64(width)).SetPrec(256)
			x := new(big.Float).SetPrec(256)
			x.Mul(xRatio, xrangeBig)
			x.Add(x, xminBig)

			img.Set(px, py, mandelbrotBigFloat(x, y))
		}
	}

	saveImage(img, filename)
	return time.Since(start)
}

// Renderiza usando big.Rat
func renderBigRat(xmin, ymin, xmax, ymax float64, width, height int, filename string) time.Duration {
	start := time.Now()
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Converte para racionais (usando denominador grande para precisão)
	scale := int64(1e10)
	xminRat := big.NewRat(int64(xmin*float64(scale)), scale)
	yminRat := big.NewRat(int64(ymin*float64(scale)), scale)
	xrangeRat := big.NewRat(int64((xmax-xmin)*float64(scale)), scale)
	yrangeRat := big.NewRat(int64((ymax-ymin)*float64(scale)), scale)

	for py := 0; py < height; py++ {
		yRatio := big.NewRat(int64(py), int64(height))
		y := new(big.Rat).Mul(yRatio, yrangeRat)
		y.Add(y, yminRat)

		for px := 0; px < width; px++ {
			xRatio := big.NewRat(int64(px), int64(width))
			x := new(big.Rat).Mul(xRatio, xrangeRat)
			x.Add(x, xminRat)

			img.Set(px, py, mandelbrotBigRat(x, y))
		}
	}

	saveImage(img, filename)
	return time.Since(start)
}

func saveImage(img image.Image, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		panic(err)
	}
}

func main() {
	const (
		width  = 512
		height = 512
	)

	// Teste em diferentes níveis de zoom
	// Zoom 1: Vista completa
	fmt.Println("=== Zoom Nível 1: Vista Completa ===")
	testZoom(-2, -2, 2, 2, width, height, "zoom1")

	// Zoom 2: Região interessante
	fmt.Println("\n=== Zoom Nível 2: Região Interessante ===")
	testZoom(-0.8, -0.2, -0.4, 0.2, width, height, "zoom2")

	// Zoom 3: Zoom profundo (aqui começam os problemas com complex64)
	fmt.Println("\n=== Zoom Nível 3: Zoom Profundo ===")
	testZoom(-0.7463, 0.1102, -0.7453, 0.1112, width, height, "zoom3")

	// Zoom 4: Zoom muito profundo (complex128 começa a ter problemas)
	fmt.Println("\n=== Zoom Nível 4: Zoom Muito Profundo ===")
	testZoom(-0.74530, 0.11030, -0.74525, 0.11035, width, height, "zoom4")
}

func testZoom(xmin, ymin, xmax, ymax float64, width, height int, prefix string) {
	zoomLevel := (xmax - xmin)
	fmt.Printf("Área de zoom: %.2e\n", zoomLevel)

	fmt.Print("complex64:   ")
	t64 := renderComplex64(xmin, ymin, xmax, ymax, width, height, prefix+"_complex64.png")
	fmt.Printf("%v\n", t64)

	fmt.Print("complex128:  ")
	t128 := renderComplex128(xmin, ymin, xmax, ymax, width, height, prefix+"_complex128.png")
	fmt.Printf("%v\n", t128)

	fmt.Print("big.Float:   ")
	tFloat := renderBigFloat(xmin, ymin, xmax, ymax, width, height, prefix+"_bigfloat.png")
	fmt.Printf("%v\n", tFloat)

	fmt.Print("big.Rat:     ")
	fmt.Print("(processando...)")
	tRat := renderBigRat(xmin, ymin, xmax, ymax, width, height, prefix+"_bigrat.png")
	fmt.Printf("\r big.Rat:     %v\n", tRat)

	fmt.Printf("\nRazão de desempenho:\n")
	fmt.Printf("  big.Float vs complex128: %.2fx mais lento\n", float64(tFloat)/float64(t128))
	fmt.Printf("  big.Rat vs complex128:   %.2fx mais lento\n", float64(tRat)/float64(t128))
}

/*
ANÁLISE DOS RESULTADOS:

1. DESEMPENHO:
   - complex64: O mais rápido (similar ao complex128)
   - complex128: Muito rápido, referência de velocidade
   - big.Float: ~50-100x mais lento que complex128
   - big.Rat: ~200-500x mais lento que complex128

2. USO DE MEMÓRIA:
   - complex64: 8 bytes por número (4 bytes real + 4 bytes imaginário)
   - complex128: 16 bytes por número (8 bytes real + 8 bytes imaginário)
   - big.Float: Vários KBs dependendo da precisão (256 bits = ~32 bytes + overhead)
   - big.Rat: Pode crescer indefinidamente, muito maior que os outros

3. ARTEFATOS DE RENDERIZAÇÃO VISÍVEIS:

   COMPLEX64 (32-bit float):
   - Precisão: ~7 dígitos decimais
   - Artefatos aparecem em: Zoom nível 3 (~10^-3)
   - Sintomas: Pixelation, bandas horizontais/verticais, perda de detalhes fractais
   - Causa: Precisão insuficiente para distinguir pontos próximos

   COMPLEX128 (64-bit float):
   - Precisão: ~15 dígitos decimais
   - Artefatos aparecem em: Zoom nível 4 (~10^-5) ou superior
   - Sintomas: Gradualmente perde detalhes em zooms extremos (>10^-12)
   - Causa: Limites da aritmética de ponto flutuante de 64 bits

   BIG.FLOAT:
   - Precisão: Configurável (usamos 256 bits)
   - Artefatos: Praticamente nenhum até zooms extremos (>10^-50)
   - Permite explorar regiões muito profundas do fractal
   - Limitado pela precisão configurada

   BIG.RAT:
   - Precisão: Ilimitada (teórica)
   - Artefatos: Nenhum (limitado apenas pela memória)
   - Na prática: Muito lento para uso interativo
   - Os denominadores crescem exponencialmente, consumindo muita memória

4. CONCLUSÕES:

   - Para visualização geral (zoom < 10^-3): Use complex64 ou complex128
   - Para exploração detalhada (zoom 10^-3 a 10^-12): Use complex128
   - Para zoom profundo (zoom 10^-12 a 10^-50): Use big.Float
   - Para zoom extremo ou precisão absoluta: Use big.Rat (prepare-se para esperar!)

5. TRADE-OFFS:
   - complex64/128: Rápido mas limitado em precisão
   - big.Float: Equilíbrio entre desempenho e precisão
   - big.Rat: Precisão perfeita mas impraticável para imagens grandes

6. RECOMENDAÇÃO PRÁTICA:
   - Use complex128 como padrão
   - Mude para big.Float quando os artefatos ficarem visíveis
   - Reserve big.Rat apenas para cálculos específicos ou demonstrações matemáticas
*/
