# fontviewer

This is just some garbage code I wrote to see how hard it would be to render TTF fonts to an HTML5
Canvas in Go.

![Sample](.github-assets/sample.png)

## Libraries

* [gopherjs](http://www.gopherjs.org/)
* [freetype-go](https://code.google.com/p/freetype-go/) ([godoc](https://godoc.org/code.google.com/p/freetype-go/freetype#Context.SetHinting))

## Building

```
$ go get github.com/zach-klippenstein/gofontviewer
$ ./build.sh
$ open fontviewer.html # or whatever command you use to open a browser
```

## Adding Fonts

Comes with Google's [Roboto](https://www.google.com/fonts/specimen/Roboto).

To add a font, drop a TTF file in `fonts/` and add a `//go:generate …` line to [fonts/fonts.go](fonts/fonts.go).