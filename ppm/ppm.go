package ppm

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
)

type Pixel [3]uint8

type PPMImg struct {
	Header    *bytes.Buffer
	ImageData []Pixel

	Magic    string
	Width    int
	Height   int
	MaxColor int
}

func NewPPmImg(width int, height int) *PPMImg {
	maxcolor := 255
	magic := "P6\n"

	dims := fmt.Sprintf("%d %d\n", width, height)
	maxCol := fmt.Sprintf("%d\n", maxcolor)

	header := bytes.Buffer{}
	header.WriteString(magic)
	header.WriteString(dims)
	header.WriteString(maxCol)

	imgData := make([]Pixel, width*height)

	return &PPMImg{
		Header:    &header,
		ImageData: imgData,
		Magic:     "P6",
		Width:     width,
		Height:    height,
		MaxColor:  maxcolor,
	}
}

func (p *PPMImg) WriteImageData(x, y int, r, g, b float64) {
	pos := x + (y * p.Width)
	p.ImageData[pos][0] = uint8(r)
	p.ImageData[pos][1] = uint8(g)
	p.ImageData[pos][2] = uint8(b)
}

func (p *PPMImg) PixelAt(x, y int) (Pixel, error) {
	if x < 0 || x >= p.Width || y < 0 || y >= p.Height {
		return Pixel{}, errors.New("invalid coordinates")
	}
	pos := x + (y * p.Width)
	return p.ImageData[pos], nil
}

func (p *PPMImg) PPMImgWriter(w io.Writer) error {
	err := binary.Write(w, binary.BigEndian, p.Header.Bytes())
	if err != nil {
		return err
	}

	for i, d := range p.ImageData {
		errR := binary.Write(w, binary.BigEndian, d[0])
		if errR != nil {
			fmt.Fprintf(os.Stderr, fmt.Sprintf("error writing R value at pixel %d\n", i))
		}
		errG := binary.Write(w, binary.BigEndian, d[1])
		if errG != nil {
			fmt.Fprintf(os.Stderr, fmt.Sprintf("error writing G value at pixel %d\n", i))
		}
		errB := binary.Write(w, binary.BigEndian, d[2])
		if errB != nil {
			fmt.Fprintf(os.Stderr, fmt.Sprintf("error writing B value at pixel %d\n", i))
		}
	}

	return nil
}
