package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

func exercise06() {
	file, _ := os.Create("lissajous.gif")
	palette := []color.Color{color.Black}
	for i := 0; i < 255; i++ {
		palette = append(palette, color.RGBA{0x00, uint8(i), 0x00, 0xff})
	}

	lissajous2(file, palette)
}

func lissajous2(out io.Writer, palette []color.Color) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	index := uint8(1)
	inverse := false
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			if inverse {
				index -= 1
				if index <= 1 {
					inverse = false
				}
			} else {
				index += 1
				if index >= uint8(len(palette)-1) {
					inverse = true
				}
			}

			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), index)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
