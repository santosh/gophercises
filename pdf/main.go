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
	pdf.MoveTo(0, 0)
	pdf.SetFont("arial", "B", 16)
	_, lineHt := pdf.GetFontSize()
	pdf.SetTextColor(255, 0, 0)
	pdf.Text(0, lineHt, "Hello, world!")
	pdf.MoveTo(0, lineHt*2.0)

	pdf.SetFont("times", "", 18)
	pdf.SetTextColor(100, 100, 100)
	_, lineHt = pdf.GetFontSize()
	pdf.MultiCell(0, lineHt*1.5, "Here is some text. If it is too long it will be word wrapped automatically. If there is a new line if will be\nwrapped as well (unlike other ways of writing text in gopdf).", gofpdf.BorderNone, gofpdf.AlignRight, false)

	// Basic Shapes
	pdf.SetFillColor(0, 255, 0)
	pdf.SetDrawColor(0, 0, 255)
	pdf.Rect(10, 100, 100, 100, "FD")
	pdf.SetFillColor(100, 200, 200)
	pdf.Polygon([]gofpdf.PointType{
		{110, 250},
		{160, 300},
		{110, 350},
		{60, 300},
	}, "F")

	err := pdf.OutputFileAndClose("p1.pdf")
	if err != nil {
		panic(err)
	}
}
