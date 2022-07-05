package pdf

import (
	"fmt"
	"os"
	"path"

	"github.com/jung-kurt/gofpdf"
	"training.go/gencert/cert"
)

type Saver struct {
	OutputDir string
}

func New(outputdir string) (*Saver, error) {
	var p *Saver
	err := os.MkdirAll(outputdir, os.ModePerm)
	if err != nil {
		return p, err
	}

	p = &Saver{
		OutputDir: outputdir,
	}

	return p, nil
}

func (p *Saver) Save(cert cert.Cert) error {
	pdf := gofpdf.New(gofpdf.OrientationLandscape, "mm", "A4", "")

	pdf.SetTitle(cert.LabelTitle, false)
	pdf.AddPage()

	// Background
	background(pdf)

	// Header
	header(pdf, &cert)

	// save file
	filename := fmt.Sprintf("%v.pdf", cert.LabelTitle)
	path := path.Join(p.OutputDir, filename)
	err := pdf.OutputFileAndClose(path)
	if err != nil {
		return err
	}
	fmt.Printf("Saved certificate to '%v' \n", path)

	return nil
}

// generation du backround
func background(pdf *gofpdf.Fpdf) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}
	pageWidth, pageHeight := pdf.GetPageSize()
	filename := "img/background.png"
	pdf.ImageOptions(filename,
		0,
		0,
		pageWidth,
		pageHeight,
		false,
		opts,
		0,
		"")
}

func header(pdf *gofpdf.Fpdf, c *cert.Cert) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}

	margin := 30.0
	x := 0.0
	imageWidth := 30.0
	filename := "img/gopher.png"
	pdf.ImageOptions(filename,
		x+margin,
		20,
		imageWidth,
		0,
		false,
		opts,
		0,
		"")

	pageWidth, _ := pdf.GetPageSize()
	x = pageWidth - imageWidth
	pdf.ImageOptions(filename,
		x-margin,
		20,
		imageWidth,
		0,
		false,
		opts,
		0,
		"")

	pdf.SetFont("Helvetica", "", 40)
	pdf.WriteAligned(0, 50, c.LabelCompletion, "C")
}
