package main

import (
	"image"
	"image/color"
	"image/gif"
	"math"
	"math/rand/v2"
	"os"
)

var palette = []color.Color{color.White, color.RGBA{0, 0xff, 0xBB, 0xff}}

const (
	whiteIndex = iota
	blackIndex
)

func lissajous() {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)

	}
	f, _ := os.Create("anim3.gif")
	gif.EncodeAll(f, &anim)
}

func main() {
	// rand.Seed(time.Now().UTC().UnixNano())
	// rand.New(rand.NewSource(time.Now().UnixNano()))
	lissajous()
}
