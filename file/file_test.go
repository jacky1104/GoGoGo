package GoGoGo

import (
	"testing"
	"github.com/jung-kurt/gofpdf"
	"fmt"
)

func TestGeneratePdf(t * testing.T){

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")
	err := pdf.OutputFileAndClose("hello.pdf")

	if err != nil{
		fmt.Println("pdf error")
	}
}