package drawing

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"net/http"
)

func Main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	rect := image.Rect(0, 0, 300, 300)
	p := color.Palette{color.Black, color.White}
	img := image.NewPaletted(rect, p)
	for i := 0; i < 300; i++ {
		img.SetColorIndex(i, i, 1)
		img.SetColorIndex(i, 299-i, 1)
	}

	png.Encode(w, img)
}
