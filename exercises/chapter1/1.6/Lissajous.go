// Упражнение 1.6.6
// Измените программу lissajous так, чтобы она генерировала изображения разных цветов,
// добавляя в палитру palette больше значений, а затем выводя их,
// путем изменения третьего аргумента функции SetColorIndex
// некоторым нетривиальным способом.
// example: go run ./Lissajous.go > 04.png

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{
	color.RGBA{R: 0xFF, G: 0x00, B: 0x00, A: 0xFF}, // red
	color.RGBA{R: 0x00, G: 0xFF, B: 0x00, A: 0xFF}, // green
	color.RGBA{R: 0xFF, G: 0xFF, B: 0x00, A: 0xFF}, // yellow
	color.RGBA{R: 0x00, G: 0x00, B: 0xFF, A: 0xFF}, // blue

}

const (
	redIndex    = 0 // Первый цвет палитры (red)
	greenIndex  = 1 // Следующий цвет (green)
	blueIndex   = 2 // Следующий цвет (blue)
	yellowIndex = 3 // Следующий цвет (yellow)
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // Количество полных колебаний x
		res     = 0.001 // Угловой разрешение
		size    = 100   // Канва изображения охватывает [size..+size]
		nframes = 64    // Количество кадров анимации
		delay   = 8     // Задержка между кадрами (единица - 10-мс)
	)

	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0 // Относительная частота колебания y
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // Разность фаз
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // Примечание: игнорируем ошибки
}
