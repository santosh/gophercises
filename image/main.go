package main

import (
	"math/rand"
	"os"

	svg "github.com/ajstarks/svgo"
)

func rn(n int) int { return rand.Intn(n) }

func main() {
	canvas := svg.New(os.Stdout)
	data := []struct {
		Month string
		Usage int
	}{
		{"Jan", 171},
		{"Feb", 180},
		{"Mar", 100},
		{"Apr", 87},
		{"May", 66},
		{"Jun", 40},
		{"Jul", 32},
		{"Aug", 55},
	}
	width := len(data)*60 + 10
	height := 300
	threshold := 120
	max := 0
	for _, item := range data {
		if item.Usage > max {
			max = item.Usage
		}
	}
	canvas.Start(width, height)
	for i, val := range data {
		percent := val.Usage * (height - 50) / max
		canvas.Rect(i*60+10, (height-50)-percent, 50, percent, "fill:rgb(77, 200, 232)")
		canvas.Text(i*60+10, height-24, val.Month, "font-size:14pt; fill:rgb(150, 150, 150); text-anchor: middle")
	}
	thresPercent := threshold * (height - 50) / max
	canvas.Line(0, height-thresPercent, width, height-thresPercent, "stroke: rgb(255, 100, 100); opacity: 0.8; stroke-width: 2px")
	canvas.Rect(0, 0, width, height-thresPercent, "fill: rgb(255, 100, 100); opacity: 0.3")
	canvas.Line(0, height-50, width, height-50, "stroke: rgb(150, 150, 150); stroke-width:2;")
	canvas.End()
}
