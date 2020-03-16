package AddWaterMark

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"

	"io/ioutil"
	"os"
	"path"

	"math"
	"math/rand"
	"time"

	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"

	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/vgimg"

	"github.com/MDGSF/utils"
	"github.com/MDGSF/utils/log"
	"github.com/spf13/cobra"
)

var Program = "AddWaterMark"
var Version = "0.0.1"
var BuildTime = ""

var SourceImage = ""
var WaterMarkText = ""

var R uint8
var G uint8
var B uint8
var A uint8

var RandomAngle float64

var FontSize float64

func init() {
	rand.Seed(time.Now().UnixNano())

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&SourceImage, "source", "s", "", "image file or directory")
	rootCmd.PersistentFlags().StringVarP(&WaterMarkText, "text", "t", "minieye", "watermark text")

	rootCmd.PersistentFlags().Uint8VarP(&R, "Red", "r", 255, "text color")
	rootCmd.PersistentFlags().Uint8VarP(&G, "Green", "g", 255, "text color")
	rootCmd.PersistentFlags().Uint8VarP(&B, "Blue", "b", 255, "text color")
	rootCmd.PersistentFlags().Uint8VarP(&A, "Alpha", "a", 20, "text color")

	rootCmd.PersistentFlags().Float64VarP(&FontSize, "font", "f", 42, "font size")

	rootCmd.Flags().BoolP("version", "v", false, "Show AddWaterMark version.")

	rootCmd.AddCommand(versionCmd)
}

func initConfig() {
	initVersionFlags()

	// RandomAngle = math.Pi * rand.Float64() / 2
	// log.Info("RandomAngle = %v", RandomAngle)
	RandomAngle = 0.6
}

func initVersionFlags() {
	needShowVersion, err := rootCmd.Flags().GetBool("version")
	if err != nil {
		fmt.Println("no version flags")
	}

	if needShowVersion {
		ShowVerion()
		os.Exit(0)
	}
}

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v", "V", "Version"},
	Short:   "Show AddWaterMark version.",
	Run: func(cmd *cobra.Command, args []string) {
		ShowVerion()
	},
}

// ShowVerion 打印出版本信息
func ShowVerion() {
	fmt.Printf("%s %s (%s) [%s-%s] (%s)\n", Program, Version, BuildTime, runtime.GOOS, runtime.GOARCH, runtime.Version())
}

// Execute program entrance.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

var rootCmd = &cobra.Command{
	Use: "AddWaterMark",
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

func start() {
	if utils.IsDir(SourceImage) {
		processDirectory()
	} else {
		processFile()
	}
}

func processDirectory() {
	files, _ := ioutil.ReadDir(SourceImage)
	for _, oneFile := range files {
		img, err := WaterMark(path.Join(SourceImage, oneFile.Name()), WaterMarkText)
		if err != nil {
			log.Error("Add Water Mark failed: err = %v", err)
			continue
		}

		fileExtention := path.Ext(oneFile.Name())
		dstFileBaseName := strings.Split(oneFile.Name(), ".")[0] + "_marked"
		dstFileName := dstFileBaseName + fileExtention
		dstPath := path.Join(SourceImage, dstFileName)

		err = SaveMarkedImage(img, fileExtention, dstPath)
		if err != nil {
			log.Error("Save Marked Image failed: err = %v", err)
			continue
		}
	}
}

func processFile() {
	img, err := WaterMark(SourceImage, WaterMarkText)
	if err != nil {
		log.Error("Add Water Mark failed: err = %v", err)
		os.Exit(0)
	}

	fileExtention := path.Ext(SourceImage)
	dstFileBaseName := strings.Split(SourceImage, ".")[0] + "_marked"
	dstPath := dstFileBaseName + fileExtention

	err = SaveMarkedImage(img, fileExtention, dstPath)
	if err != nil {
		log.Error("Save Marked Image failed: err = %v", err)
		os.Exit(0)
	}
}

// SaveMarkedImage save marked image to dstPath
func SaveMarkedImage(img image.Image, ext string, dstPath string) error {
	ext = strings.ToLower(ext)
	buff := new(bytes.Buffer)
	switch ext {
	case ".jpg", ".jpeg":
		err := jpeg.Encode(buff, img, &jpeg.Options{Quality: 100})
		if err != nil {
			return err
		}
	case ".png":
		err := png.Encode(buff, img)
		if err != nil {
			return err
		}
	}

	f, err := os.Create(dstPath)
	if err != nil {
		return err
	}

	if _, err = buff.WriteTo(f); err != nil {
		return err
	}

	return nil
}

// WaterMark for adding a watermark on the image
func WaterMark(filepath, markText string) (image.Image, error) {

	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	// image's length to canvas's length
	bounds := img.Bounds()
	w := vg.Length(bounds.Max.X) * vg.Inch / vgimg.DefaultDPI
	h := vg.Length(bounds.Max.Y) * vg.Inch / vgimg.DefaultDPI
	diagonal := vg.Length(math.Sqrt(float64(w*w + h*h)))

	// create a canvas, which width and height are diagonal
	c := vgimg.New(diagonal, diagonal)

	// draw image on the center of canvas
	rect := vg.Rectangle{}
	rect.Min.X = diagonal/2 - w/2
	rect.Min.Y = diagonal/2 - h/2
	rect.Max.X = diagonal/2 + w/2
	rect.Max.Y = diagonal/2 + h/2
	c.DrawImage(rect, img)

	// make a fontStyle, which width is vg.Inch * 0.7
	fontStyle, _ := vg.MakeFont("Courier", vg.Length(FontSize))

	// repeat the markText
	markTextWidth := fontStyle.Width(markText)
	unitText := markText
	for markTextWidth <= diagonal {
		markText += strings.Repeat(" ", 3) + unitText
		markTextWidth = fontStyle.Width(markText)
	}

	// set the color of markText
	c.SetColor(color.RGBA{R, G, B, A})

	// set a random angle between 0 and π/2
	c.Rotate(RandomAngle)

	// set the lineHeight and add the markText
	lineHeight := fontStyle.Extents().Height * 1
	for offset := -2 * diagonal; offset < 2*diagonal; offset += lineHeight {
		c.FillString(fontStyle, vg.Point{X: 0, Y: offset}, markText)
	}

	// canvas writeto jpeg
	// canvas.img is private
	// so use a buffer to transfer
	jc := vgimg.PngCanvas{Canvas: c}
	buff := new(bytes.Buffer)
	jc.WriteTo(buff)
	img, _, err = image.Decode(buff)
	if err != nil {
		return nil, err
	}

	// get the center point of the image
	ctp := int(diagonal * vgimg.DefaultDPI / vg.Inch / 2)

	// cutout the marked image
	size := bounds.Size()
	bounds = image.Rect(ctp-size.X/2, ctp-size.Y/2, ctp+size.X/2, ctp+size.Y/2)
	rv := image.NewRGBA(bounds)
	draw.Draw(rv, bounds, img, bounds.Min, draw.Src)
	return rv, nil
}
