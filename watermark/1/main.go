package main

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

func main() {
	imgFile, _ := os.Open("test.jpg")
	defer imgFile.Close()
	img, _ := jpeg.Decode(imgFile)

	wmFile, _ := os.Open("penguin.png")
	defer wmFile.Close()
	wmImg, _ := png.Decode(wmFile)

	x := img.Bounds().Dx() - wmImg.Bounds().Dx() - 10
	y := img.Bounds().Dy() - wmImg.Bounds().Dy() - 10
	offset := image.Pt(x, y)

	b := img.Bounds()
	m := image.NewNRGBA(b)

	draw.Draw(m, b, img, image.ZP, draw.Src)
	draw.Draw(m, wmImg.Bounds().Add(offset), wmImg, image.ZP, draw.Over)

	newImg, _ := os.Create("new.jpg")
	defer newImg.Close()
	jpeg.Encode(newImg, m, &jpeg.Options{100})
}
