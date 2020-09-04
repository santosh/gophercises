package main

import (
	"math/rand"
	"os"

	svg "github.com/ajstarks/svgo"
)

func rn(n int) int { return rand.Intn(n) }

func main() {
	canvas := svg.New(os.Stdout)
	data := []int{100, 33, 73, 64}
	width := len(data)*60 + 10
	height := 100
	canvas.Start(width, height)
	for i, val := range data {
		canvas.Rect(i*60+10, height-val, 50, val, "fill:rgb(77, 200, 232)")
	}
	canvas.End()
}
