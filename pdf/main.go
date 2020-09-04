package main

import (
	"fmt"

	"github.com/phpdave11/gofpdf"
)

func main() {
	pdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitPoint, gofpdf.PageSizeLetter, "")
	w, h := pdf.GetPageSize()
	fmt.Printf("width=%v, height=%v\n", w, h)
	pdf.AddPage()

	// Basic Text Stuff
	pdf.SetFont("arial", "B", 16)
	_, lineHt := pdf.GetFontSize()
	pdf.Text(0, lineHt, "Hello, world!")
	err := pdf.OutputFileAndClose("p1.pdf")
	if err != nil {
		panic(err)
	}
}
