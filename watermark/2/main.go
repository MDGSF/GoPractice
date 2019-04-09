package main

import (
	"bufio"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"

	"github.com/golang/freetype"
	"golang.org/x/image/font"
)

func main() {
	fontBytes, _ := ioutil.ReadFile("WeatherSunday.otf")
	myfont, _ := freetype.ParseFont(fontBytes)

	rgba := image.NewRGBA(image.Rect(0, 0, 640, 480))
	draw.Draw(rgba, rgba.Bounds(), image.White, image.ZP, draw.Src)

	c := freetype.NewContext()
	c.SetFont(myfont)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(image.Black)
	c.SetHinting(font.HintingFull)

	ruler := color.RGBA{0xdd, 0xdd, 0xdd, 0xff}
	for i := 0; i < 200; i++ {
		rgba.Set(10, 10+i, ruler)
		rgba.Set(10+i, 10, ruler)
	}

	pt := freetype.Pt(10, 10+int(c.PointToFixed(12)>>6))
	s := "This is a beautiful day."
	c.DrawString(s, pt)

	newFile, _ := os.Create("newImg.png")
	defer newFile.Close()

	b := bufio.NewWriter(newFile)
	png.Encode(b, rgba)
	b.Flush()
}
