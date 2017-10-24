package complex

import (
	"fmt"
	"testing"
	"time"
)

// make an image of a given size
// fill it with a colour
// save it as a png

const bmo = `░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
░░░░▄██████████████████████▄░░░░
░░░░█░░░░░░░░░░░░░░░░░░░░░░█░░░░
░░░░█░▄██████████████████▄░█░░░░
░░░░█░█░░░░░░░░░░░░░░░░░░█░█░░░░
░░░░█░█░░░░░░░░░░░░░░░░░░█░█░░░░
░░░░█░█░░█░░░░░░░░░░░░█░░█░█░░░░
░░░░█░█░░░░░▄▄▄▄▄▄▄▄░░░░░█░█░░░░
░░░░█░█░░░░░▀▄░░░░▄▀░░░░░█░█░░░░
░░░░█░█░░░░░░░▀▀▀▀░░░░░░░█░█░░░░
░░░░█░█░░░░░░░░░░░░░░░░░░█░█░░░░
░█▌░█░▀██████████████████▀░█░▐█░
░█░░█░░░░░░░░░░░░░░░░░░░░░░█░░█░
░█░░█░████████████░░░░░██░░█░░█░
░█░░█░░░░░░░░░░░░░░░░░░░░░░█░░█░
░█░░█░░░░░░░░░░░░░░░▄░░░░░░█░░█░
░▀█▄█░░░▐█▌░░░░░░░▄███▄░██░█▄█▀░
░░░▀█░░█████░░░░░░░░░░░░░░░█▀░░░
░░░░█░░░▐█▌░░░░░░░░░▄██▄░░░█░░░░
░░░░█░░░░░░░░░░░░░░▐████▌░░█░░░░
░░░░█░▄▄▄░▄▄▄░░░░░░░▀██▀░░░█░░░░
░░░░█░░░░░░░░░░░░░░░░░░░░░░█░░░░
░░░░▀██████████████████████▀░░░░
░░░░░░░░██░░░░░░░░░░░░██░░░░░░░░
░░░░░░░░██░░░░░░░░░░░░██░░░░░░░░
░░░░░░░░██░░░░░░░░░░░░██░░░░░░░░
░░░░░░░░██░░░░░░░░░░░░██░░░░░░░░
░░░░░░░▐██░░░░░░░░░░░░██▌░░░░░░░
░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
`
const (
	w = 1000
	h = 1000
)

func _Test_RandomFillImage(t *testing.T) {

	img, err := MakeImage(w, h)

	if err != nil {
		t.Fatal(err)
	}
	if img == nil {
		t.Fatal("No image returned")
	}
	if width := img.Rect.Max.X; width != w {
		t.Fatal("Width was incorrect, recieved ", width)
	}
	if height := img.Rect.Max.Y; height != h {
		t.Fatal("Width was incorrect, recieved ", height)
	}

	img, err = RandomFillImage(img)

	if err != nil {
		t.Fatal(err)
	}
	fn := "random_fill"
	if err := SaveImage(fn, img); err != nil {
		t.Fatal(err)
	}
}

func Test_FillWithFunction(t *testing.T) {
	img, err := MakeImage(w, h)

	if err != nil {
		t.Fatal(err)
	}
	if img == nil {
		t.Fatal("No image returned")
	}
	if width := img.Rect.Max.X; width != w {
		t.Fatal("Width was incorrect, recieved ", width)
	}
	if height := img.Rect.Max.Y; height != h {
		t.Fatal("Width was incorrect, recieved ", height)
	}

	start := time.Now()
	img, err = FillWithFunction(img)
	elapsed := time.Since(start)
	fmt.Printf("Filling image took %s\n", elapsed)

	if err != nil {
		t.Fatal(err)
	}

	fn := "function_fill"
	if err := SaveImage(fn, img); err != nil {
		t.Fatal(err)
	}

}
