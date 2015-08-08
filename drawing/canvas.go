package drawing

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"github.com/gopherjs/gopherjs/js"
)

type Canvas struct {
	elem, context *js.Object
	bounds        image.Rectangle
}

func NewCanvas(parent *js.Object, width, height int) *Canvas {
	document := js.Global.Get("document")
	canvas := document.Call("createElement", "canvas")
	parent.Call("appendChild", canvas)

	canvas.Set("width", width)
	canvas.Set("height", height)

	context := canvas.Call("getContext", "2d")

	return &Canvas{canvas, context, image.Rect(0, 0, width, height)}
}

func (c *Canvas) SetFillStyle(style string) {
	c.context.Set("fillStyle", style)
}

func (c *Canvas) SetFillColor(color color.RGBA) {
	c.SetFillStyle(fmt.Sprintf("rgb(%d,%d,%d)", color.R, color.G, color.B))
}

func (c *Canvas) FillRect(r image.Rectangle) {
	c.context.Call("fillRect", r.Min.X, r.Min.Y, r.Dx(), r.Dy())
}

func (c *Canvas) Fill(color color.RGBA) {
	c.SetFillColor(color)
	c.FillRect(c.bounds)
}

func (c *Canvas) Width() int {
	return c.bounds.Dx()
}

func (c *Canvas) Height() int {
	return c.bounds.Dy()
}

// GetImage returns a read-only copy of the contents of this Canvas.
func (c *Canvas) GetImage() image.Image {
	return c.getImageData()
}

func (c *Canvas) Draw(drawer func(draw.Image)) {
	img := c.getImageData()
	defer c.setImageData(img)
	drawer(img)
}

func (c *Canvas) getImageData() imageData {
	imageData := newImageData(c.context.Call("getImageData", 0, 0, c.Width(), c.Height()))

	if imageData.Width() != c.Width() || imageData.Height() != c.Height() {
		panic(fmt.Sprintf("(%d, %d) != (%d, %d)",
			imageData.Width(), imageData.Height(), c.Width(), c.Height()))
	}

	return imageData
}

func (c *Canvas) setImageData(img imageData) {
	c.context.Call("putImageData", img.Object, 0, 0)
}
