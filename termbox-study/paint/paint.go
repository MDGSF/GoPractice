package main

import termbox "github.com/nsf/termbox-go"

var curCol = 0
var curRune = 0
var backbuf []termbox.Cell
var bbw, bbh int

var runes = []rune{' ', '░', '▒', '▓', '█'}
var colors = []termbox.Attribute{
	termbox.ColorBlack,
	termbox.ColorRed,
	termbox.ColorGreen,
	termbox.ColorYellow,
	termbox.ColorBlue,
	termbox.ColorMagenta,
	termbox.ColorCyan,
	termbox.ColorWhite,
}

func update_and_redraw_all(mx, my int) {

}

func reallocBackBuffer(w, h int) {

}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	reallocBackBuffer(termbox.Size())
	update_and_redraw_all(-1, -1)

mainloop:
	for {
		mx, my := -1, -1
	}
}
