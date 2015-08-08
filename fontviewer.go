package main

import (
	"image"
	"image/color"
	"image/draw"

	"code.google.com/p/freetype-go/freetype"
	"code.google.com/p/freetype-go/freetype/truetype"
	"github.com/gopherjs/gopherjs/js"
	. "github.com/zach-klippenstein/fontviewer/drawing"
	"github.com/zach-klippenstein/fontviewer/fonts"
)

func main() {
	println("Initializing DOM...")
	document := js.Global.Get("document")

	canvas := NewCanvas(document.Get("body"), 500, 200)
	canvas.Fill(color.RGBA{240, 240, 240, 255})

	println("Loading font...")
	fontData, err := fonts.LoadFont(fonts.RobotoRegular)
	if err != nil {
		panic(err)
	}
	font, err := freetype.ParseFont(fontData)
	if err != nil {
		panic(err)
	}
	println("Loaded font:", font)

	printRuneInfo(canvas, font, 'a')
	drawString(canvas, font, 24, "hello world ")
}

func printRuneInfo(canvas *Canvas, font *truetype.Font, r rune) {
	index := font.Index(r)
	scale := int32(50)
	hmetric := font.HMetric(scale, index)
	vmetric := font.VMetric(scale, index)

	println("Index:", index)
	println("FUnitsPerEm:", font.FUnitsPerEm())
	println("Scale:", scale)
	println("HMetric:", hmetric)
	println("VMetric:", vmetric)
}

func drawString(canvas *Canvas, font *truetype.Font, size float64, str string) {
	pos := freetype.Pt(10, 100)

	canvas.Draw(func(img draw.Image) {
		context := freetype.NewContext()
		context.SetFont(font)
		context.SetDst(img)
		context.SetFontSize(size)
		context.SetSrc(&image.Uniform{color.Black})
		context.SetClip(img.Bounds())
		context.SetHinting(freetype.NoHinting)

		// First draw without hinting.
		pos, _ = context.DrawString(str, pos)

		// Then draw with hinting, on the same line.
		context.SetHinting(freetype.FullHinting)
		context.DrawString(str, pos)
	})
}
