package main

import "github.com/phpdave11/gofpdf"

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")
	err := pdf.OutputFileAndClose("p1.pdf")
	if err != nil {
		panic(err)
	}
}
