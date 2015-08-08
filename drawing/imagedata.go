package drawing

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/gopherjs/gopherjs/js"
)

// See https://developer.mozilla.org/en-US/docs/Web/API/Canvas_API/Tutorial/Pixel_manipulation_with_canvas
type imageData struct {
	*js.Object
	bounds image.Rectangle
	data   *js.Object // Uint8ClampedArray
}

var _ draw.Image = imageData{}

func newImageData(obj *js.Object) (id imageData) {
	id.Object = obj
	id.bounds = image.Rect(0, 0, id.Get("width").Int(), id.Get("height").Int())
	id.data = obj.Get("data")
	return
}

func (id imageData) Width() int {
	return id.bounds.Max.X
}

func (id imageData) Height() int {
	return id.bounds.Max.Y
}

// imageData implements image.Image

func (imageData) ColorModel() color.Model {
	return color.RGBAModel
}

func (id imageData) Bounds() image.Rectangle {
	return id.bounds
}

func (id imageData) At(x, y int) color.Color {
	pixelStart := id.getIndexOfPixel(x, y)

	// TODO optimize by getting a slice and getting values from that (see Uint8ClampedArray docs)
	return color.RGBA{
		uint8(id.data.Index(pixelStart).Int()),
		uint8(id.data.Index(pixelStart + 1).Int()),
		uint8(id.data.Index(pixelStart + 2).Int()),
		uint8(id.data.Index(pixelStart + 3).Int()),
	}
}

// imageData implements draw.Image

func (id imageData) Set(x, y int, c color.Color) {
	pixelStart := id.getIndexOfPixel(x, y)
	rgba := color.RGBAModel.Convert(c).(color.RGBA)

	// TODO optimize by getting a slice and setting values on that (see Uint8ClampedArray docs)
	id.data.SetIndex(pixelStart, rgba.R)
	id.data.SetIndex(pixelStart+1, rgba.G)
	id.data.SetIndex(pixelStart+2, rgba.B)
	id.data.SetIndex(pixelStart+3, rgba.A)
}

/*
Each pixel is represented by four one-byte values
(red, green, blue, and alpha, in that order; that is, "RGBA" format).
Each color component is represented by an integer between 0 and 255.
Each component is assigned a consecutive index within the array, with the
top left pixel's red component being at index 0 within the array. Pixels
then proceed from left to right, then downward, throughout the array.

For example, to read the blue component's value from the pixel at
column 200, row 50 in the image, you would do the following:

	blueComponent = imageData.data[((50*(imageData.width*4)) + (200*4)) + 2];
*/
func (id imageData) getIndexOfPixel(x, y int) int {
	return (x + y*id.Width()) * 4
}
