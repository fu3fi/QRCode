package main

import (
	"log"

	gopdf "github.com/signintech/gopdf"
	qrcode "github.com/skip2/go-qrcode"
)

func createQRCode(data string) []byte {
	png, _ := qrcode.Encode(data, qrcode.Medium, 256)
	return png
}

type Pdf struct {
	cursor *gopdf.GoPdf
}

func NewPdf(pathToFont string) *Pdf {
	pdf := &Pdf{}
	pdf.cursor = &gopdf.GoPdf{}
	pdf.cursor.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	err := pdf.cursor.AddTTFFont("default", "LiberationSerif-Regular.ttf")
	if err != nil {
		log.Print(err.Error())
		return nil
	}
	pdf.cursor.SetFont("default", "", 14)
	return pdf
}

func (pdf *Pdf) addPage() {
	pdf.cursor.AddPage()
}

func (pdf *Pdf) addImg(img []byte, x float64, y float64) {
	instanceOfImgk, _ := gopdf.ImageHolderByBytes(img)
	pdf.cursor.ImageByHolder(instanceOfImgk, x, y, nil)
}

func (pdf *Pdf) addText(text string, x float64, y float64) {
	pdf.cursor.SetXY(x, y)
	err := pdf.cursor.Cell(nil, text)
	if err != nil {
		log.Print(err.Error())
		return
	}
}

func (pdf *Pdf) WritePdf(name string) {
	pdf.cursor.WritePdf(name)
}

func main() {

	ex1 := createQRCode("ex1")
	ex2 := createQRCode("ex2")
	ex3 := createQRCode("ex3")
	ex4 := createQRCode("ex4")

	pdf := NewPdf("LiberationSerif-Regular.ttf")

	pdf.addPage()

	pdf.addImg(ex1, 20, 0)
	pdf.addText("Text1", 20, 135)

	pdf.addImg(ex2, 155, 0)
	pdf.addText("Text2", 155, 135)

	pdf.addImg(ex3, 290, 0)
	pdf.addText("Text3", 290, 135)

	pdf.addImg(ex4, 425, 0)
	pdf.addText("Text4", 425, 135)

	pdf.WritePdf("ex.pdf")

}
