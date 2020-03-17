package watermark

import (
	"bytes"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

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

// Program program name
var Program = "AddWaterMark"

// Version version number
var Version = "0.0.1"

// BuildTime build time
var BuildTime = ""

// SourceImage source file or directory
var SourceImage = ""

// WaterMarkText watermark text
var WaterMarkText = ""

// OutputDirectory new image output directory
var OutputDirectory = ""

// NewImageSuffix new image suffix
var NewImageSuffix = ""

// R RGBA red
var R uint8

// G RGBA Green
var G uint8

// B RGBA Blue
var B uint8

// A RGBA Alpha
var A uint8

// RandomAngle text random angle
var RandomAngle float64

// FontSize text font size
var FontSize float64

/*
WaterMarkType watermark type
0: default
1: top-left
2: top-right
3: bottom-right
4: bottom-left
*/
var WaterMarkType int

// WorkerThreadNumber worker thread number
var WorkerThreadNumber int

var startTime time.Time

func init() {
	rand.Seed(time.Now().UnixNano())

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&SourceImage, "source", "s", "", "image file or directory")
	rootCmd.PersistentFlags().StringVarP(&WaterMarkText, "text", "t", "minieye", "watermark text")
	rootCmd.PersistentFlags().StringVarP(&OutputDirectory, "output", "o", "", "output directory")
	rootCmd.PersistentFlags().StringVarP(&NewImageSuffix, "suffix", "e", "_marked", "new image suffix")

	rootCmd.PersistentFlags().Uint8VarP(&R, "Red", "r", 255, "text color")
	rootCmd.PersistentFlags().Uint8VarP(&G, "Green", "g", 255, "text color")
	rootCmd.PersistentFlags().Uint8VarP(&B, "Blue", "b", 255, "text color")
	rootCmd.PersistentFlags().Uint8VarP(&A, "Alpha", "a", 20, "text color")

	rootCmd.PersistentFlags().Float64VarP(&FontSize, "font", "f", 42, "font size")

	rootCmd.PersistentFlags().IntVarP(&WaterMarkType, "WaterMarkType", "w", 0, "watermark type")

	rootCmd.PersistentFlags().IntVarP(&WorkerThreadNumber, "workers", "k", 0, "worker thread number")

	rootCmd.Flags().BoolP("version", "v", false, "Show AddWaterMark version.")

	rootCmd.AddCommand(versionCmd)
}

func initConfig() {
	initVersionFlags()
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

	startTime = time.Now()
	log.Info("startTime = %v", startTime)

	// RandomAngle = math.Pi * rand.Float64() / 2
	// log.Info("RandomAngle = %v", RandomAngle)
	RandomAngle = 0.6

	if len(OutputDirectory) > 0 {
		os.MkdirAll(OutputDirectory, 0755)
	}

	if WorkerThreadNumber <= 0 || WorkerThreadNumber > runtime.NumCPU() {
		WorkerThreadNumber = runtime.NumCPU()
	}
	log.Info("WorkerThreadNumber = %v", WorkerThreadNumber)

	if utils.IsDir(SourceImage) {
		processDirectory()
	} else {
		processFile()
	}
}

func processDirectory() {

	wg := sync.WaitGroup{}
	in := make(chan string, 1)
	for i := 0; i < WorkerThreadNumber; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, in chan string) {
			defer wg.Done()
			for {
				select {
				case filePathName, ok := <-in:
					if !ok {
						return
					}

					fileExtension := path.Ext(filePathName)
					if !isValidImageExt(fileExtension) {
						continue
					}

					img, err := WaterMark(filePathName, WaterMarkText)
					if err != nil {
						log.Error("Add Water Mark failed: err = %v", err)
						continue
					}

					var dstFileBaseName string
					var dstPath string
					if len(OutputDirectory) == 0 {
						dstFileBaseName = strings.TrimSuffix(filePathName, fileExtension) + NewImageSuffix
						dstPath = dstFileBaseName + fileExtension
					} else {
						relativeSubPathName := strings.TrimPrefix(filePathName, SourceImage)
						newFilePathName := path.Join(OutputDirectory, relativeSubPathName)
						dstFileBaseName = strings.TrimSuffix(newFilePathName, fileExtension) + NewImageSuffix
						dstPath = dstFileBaseName + fileExtension
						newFileDir := path.Dir(dstPath)
						os.MkdirAll(newFileDir, 0755)
					}

					err = SaveMarkedImage(img, fileExtension, dstPath)
					if err != nil {
						log.Error("Save Marked Image failed: err = %v", err)
						continue
					}
				}
			}
		}(&wg, in)
	}

	count := 0
	err := filepath.Walk(SourceImage, func(filePathName string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		count++
		if count%100 == 0 {
			curTime := time.Now()
			log.Info("processed image number [%v], curTime = %v, elapsed = %v",
				count, curTime, curTime.Sub(startTime))
		}

		in <- filePathName

		return nil
	})
	if err != nil {
		log.Error("%v", err)
		return
	}

	close(in)
	wg.Wait()

	endTime := time.Now()
	log.Info("endTime = %v, elapsed = %v", endTime, endTime.Sub(startTime))
}

func processFile() {
	img, err := WaterMark(SourceImage, WaterMarkText)
	if err != nil {
		log.Error("Add Water Mark failed: err = %v", err)
		os.Exit(0)
	}

	fileExtension := path.Ext(SourceImage)
	if !isValidImageExt(fileExtension) {
		log.Error("Invalid file extension= %v", err)
		os.Exit(0)
	}

	dstFileBaseName := strings.TrimSuffix(SourceImage, fileExtension) + "_marked"
	dstPath := dstFileBaseName + fileExtension

	err = SaveMarkedImage(img, fileExtension, dstPath)
	if err != nil {
		log.Error("Save Marked Image failed: err = %v", err)
		os.Exit(0)
	}
}

func isValidImageExt(ext string) bool {
	ext = strings.ToLower(ext)
	if ext == ".jpg" || ext == ".jpeg" || ext == ".png" {
		return true
	}
	return false
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
	fontStyle, err := vg.MakeFont("Courier", vg.Length(FontSize))
	if err != nil {
		return nil, err
	}

	// set the color of markText
	c.SetColor(color.RGBA{R, G, B, A})

	if WaterMarkType == 0 {
		// repeat the markText
		markTextWidth := fontStyle.Width(markText)
		unitText := markText
		for markTextWidth <= diagonal {
			markText += strings.Repeat(" ", 3) + unitText
			markTextWidth = fontStyle.Width(markText)
		}

		// set a random angle between 0 and π/2
		c.Rotate(RandomAngle)

		// set the lineHeight and add the markText
		lineHeight := fontStyle.Extents().Height * 1
		for offset := -2 * diagonal; offset < 2*diagonal; offset += lineHeight {
			c.FillString(fontStyle, vg.Point{X: 0, Y: offset}, markText)
		}
	} else if WaterMarkType == 1 {
		// upper-left
		// set the lineHeight and add the markText
		lineHeight := fontStyle.Extents().Height * 1
		c.FillString(fontStyle, vg.Point{
			X: diagonal/2 - w/2 + 20,
			Y: diagonal/2 + h/2 - lineHeight,
		}, markText)
	} else if WaterMarkType == 2 {
		// upper-right
		// set the lineHeight and add the markText
		markTextWidth := fontStyle.Width(markText)
		lineHeight := fontStyle.Extents().Height * 1
		c.FillString(fontStyle, vg.Point{
			X: diagonal/2 + w/2 - markTextWidth - 20,
			Y: diagonal/2 + h/2 - lineHeight,
		}, markText)
	} else if WaterMarkType == 3 {
		// bottom-right
		// set the lineHeight and add the markText
		markTextWidth := fontStyle.Width(markText)
		//lineHeight := fontStyle.Extents().Height * 1
		c.FillString(fontStyle, vg.Point{
			X: diagonal/2 + w/2 - markTextWidth - 20,
			Y: diagonal/2 - h/2 + 20,
		}, markText)
	} else if WaterMarkType == 4 {
		// bottom-left
		// set the lineHeight and add the markText
		//lineHeight := fontStyle.Extents().Height * 1
		c.FillString(fontStyle, vg.Point{
			X: diagonal/2 - w/2 + 20,
			Y: diagonal/2 - h/2 + 20,
		}, markText)
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
