package complex

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	_ "math"
	"math/cmplx"
	"math/rand"
	"os"
	"time"
)

func MakeImage(width, height int) (*image.RGBA, error) {

	r := image.Rect(0, 0, width, height)
	img := image.NewRGBA(r)

	return img, nil
}

func RandomFillImage(img *image.RGBA) (*image.RGBA, error) {
	for v := range img.Pix {
		x, y := v%img.Stride, v/img.Stride
		r, g, b := rand.Intn(255), rand.Intn(255), rand.Intn(255)
		c := &color.RGBA{uint8(r), uint8(g), uint8(b), 255}
		img.Set(x, y, c)
	}

	return img, nil
}

func FillWithFunction(img *image.RGBA) (*image.RGBA, error) {
	for v := range img.Pix {
		x, y := float64(v%img.Stride), float64(v/img.Stride)
		// r, g, b := math.Cos(y/50), math.Hypot(x, y), math.Log(x+y)
		// r, g, b := 0, 0, math.Log(x+y)
		// r, g, b := 0, 0, math.Log(x+y)
		z := complex(x, y)
		// r, g, b := cmplx.Abs(cmplx.Log(z)), cmplx.Phase(z), cmplx.Abs(cmplx.Sinh(z/500))
		r, g, b := cmplx.Phase(1/z), cmplx.Abs(1/z), 0

		c := &color.RGBA{uint8(255 * r), uint8(255 * g), uint8(255 * b), 255}
		img.Set(int(x), int(y), c)
	}

	return img, nil
}

func SaveImage(filename string, img *image.RGBA) error {
	rand.Seed(time.Now().UTC().UnixNano())
	filename = fmt.Sprintf("%v_%d.png", filename, rand.Int())

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		return err
	}
	return nil
}
