package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"math/big"
	"net/http"
	"sort"
)

const (
	width, height = 1600, 900
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
	startColor    = "FF0000"
	endColor      = "0000FF"
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

var c1r, c1g, c1b float64 = getColorComponents(startColor)
var c2r, c2g, c2b float64 = getColorComponents(endColor)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	writeSvg(w, r)
}

func writeSvg(w io.Writer, r *http.Request) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke:grey; fill:white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	zs := *new([]float64)
	rectangles := *new([]rectData)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			rectangles = append(rectangles, corners(i, j))
		}
	}

	zs = getZs(&rectangles)

	min, max := Minmax(zs)

	for i, rect := range rectangles {
		z := zs[i]
		zpct := (z - min) / (max - min)
		r, g, b := getShade(c1r, c1g, c1b, c2r, c2g, c2b, zpct)

		rtxt := fmt.Sprintf("%02x", int16(r))
		gtxt := fmt.Sprintf("%02x", int16(g))
		btxt := fmt.Sprintf("%02x", int16(b))
		color := rtxt + gtxt + btxt

		fmt.Fprintf(w, "<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' fill='#%v'/>\n",
			rect.a.x, rect.a.y, rect.b.x, rect.b.y, rect.c.x, rect.c.y, rect.d.x, rect.d.y, color)
	}

	fmt.Fprintln(w, "</svg>")
	sort.Float64s(zs)
	fmt.Println(zs[0], zs[len(zs)-1])
}

func getShade(c1r, c1g, c1b, c2r, c2g, c2b, hue float64) (float64, float64, float64) {
	return ((c2r - c1r) * hue) + c1r,
		((c2g - c1g) * hue) + c1g,
		((c2b - c1b) * hue) + c1b
}

func getColorComponents(s string) (float64, float64, float64) {
	return parseToFloat64(s[0:2]), parseToFloat64(s[2:4]), parseToFloat64(s[4:6])
}

func parseToFloat64(s string) float64 {
	var decoded big.Int
	decoded.SetString(s, 16)
	return float64(decoded.Int64())
}

func getZs(rectangles *[]rectData) []float64 {
	zs := *new([]float64)

	for _, rect := range *rectangles {
		a := Avg([]float64{rect.a.z, rect.b.z, rect.c.z, rect.d.z})
		zs = append(zs, a)
	}

	return zs
}

type rectData struct {
	a cornerData
	b cornerData
	c cornerData
	d cornerData
}

type cornerData struct {
	x float64
	y float64
	z float64
}

func corners(i, j int) rectData {
	return rectData{
		a: corner(i+1, j),
		b: corner(i, j),
		c: corner(i, j+1),
		d: corner(i+1, j+1)}
}

func corner(i, j int) cornerData {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return cornerData{x: sx, y: sy, z: z}
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
