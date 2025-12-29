// Pacote principal - ponto de entrada do programa
package main

import (
	"image"       // Tipos básicos para trabalhar com imagens
	"image/color" // Definição de cores
	"image/gif"   // Codificação/decodificação de arquivos GIF
	"io"          // Interface para operações de I/O
	"math"        // Funções matemáticas (Sin, Pi, etc)
	"math/rand"   // Geração de números aleatórios
	"os"          // Interação com sistema operacional (criar arquivos)
)

// Paleta de cores: [0]=preto (fundo), [1]=verde, [2]=laranja
var palette = []color.Color{color.Black, color.RGBA{0x00, 0xff, 0x00, 0xff}, color.RGBA{0xff, 0x7f, 0x00, 0xff}}

// Constantes que representam os índices das cores na paleta
const (
	blackIndex  = 0 // Índice 0: preto (cor de fundo)
	greenIndex  = 1 // Índice 1: verde
	orangeIndex = 2 // Índice 2: laranja
)

func main() {
	// Cria um novo arquivo chamado "lissajous.gif" para salvar a animação
	f, err := os.Create("lissajous.gif")
	if err != nil {
		// Se houver erro ao criar o arquivo, para o programa
		panic(err)
	}
	// defer: garante que o arquivo será fechado ao final da função main
	defer f.Close()
	// Chama a função lissajous passando o arquivo como destino
	lissajous(f)
}

// Função que gera a animação de figuras de Lissajous
func lissajous(out io.Writer) {
	// Constantes que controlam a animação
	const (
		cycles  = 5     // Número de oscilações completas do oscilador x
		res     = 0.001 // Resolução angular (menor = mais suave)
		size    = 100   // Tamanho da imagem (canvas será 2*size+1 x 2*size+1)
		nframes = 64    // Número de frames na animação
		delay   = 8     // Delay entre frames em unidades de 10ms (80ms)
	)
	// Frequência relativa do oscilador y (aleatória entre 0 e 3)
	freq := rand.Float64() * 3.0
	// Estrutura GIF: LoopCount=0 significa loop infinito
	anim := gif.GIF{LoopCount: 0} //no livro encontra-se GIF{LoopCount: nframes} foi modificado para GIF{LoopCount: 0}
	// Diferença de fase entre os osciladores x e y
	phase := 0.0
	// Loop que cria cada frame da animação
	for i := 0; i < nframes; i++ {
		// Define o retângulo da imagem: de (0,0) até (2*size+1, 2*size+1) = 201x201 pixels
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		// Cria uma imagem paletizada (usa índices da paleta ao invés de RGB direto)
		img := image.NewPaletted(rect, palette)
		// Loop que desenha a curva de Lissajous para este frame
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			// Calcula coordenada x usando seno de t (oscilação simples)
			x := math.Sin(t)
			// Calcula coordenada y usando seno de t*freq+phase (oscilação com frequência diferente e fase)
			y := math.Sin(t*freq + phase)
			// Alterna entre verde e laranja a cada frame
			colorIndex := uint8(greenIndex) // Por padrão usa verde
			if i%2 == 0 {                   // Se o número do frame for par
				colorIndex = uint8(orangeIndex) // Usa laranja
			}
			// Define o pixel na posição calculada com a cor escolhida
			// size+int(x*size+0.5): converte coordenada [-1,1] para [0, 2*size]
			// +0.5 faz arredondamento correto ao converter float para int
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				colorIndex)
		}
		// Incrementa a fase para o próximo frame (faz a figura girar)
		phase += 0.1
		// Adiciona o delay deste frame à animação
		anim.Delay = append(anim.Delay, delay)
		// Adiciona a imagem deste frame à animação
		anim.Image = append(anim.Image, img)
	}
	// Codifica toda a animação GIF e escreve no destino (arquivo)
	gif.EncodeAll(out, &anim)
}
