package main

import (
	"fmt"
	"log"
	"os"

	"github.com/abtiwary/goraytracer/ppm"
	"github.com/abtiwary/goraytracer/vec3"
)

func WriteTestOutput() {
	width := 256
	height := 256

	ppmImg := ppm.NewPPmImg(width, height)

	// write a gradient
	var color *vec3.Vec3
	for j := height - 1; j >= 0; j-- {
		for i := 0; i < width; i++ {
			color = vec3.NewVec3(float64(i)/float64(width-1), float64(j)/float64(height-1), 0.25)
			color.MultBy(255.)
			ppmImg.WriteImageData(i, j, color.X(), color.Y(), color.Z())
		}
	}

	// Todo read the output path from a flag
	outfile := "/Users/abhishek.tiwary/temp/myimg2.ppm"
	_, err := os.Stat(outfile)
	if err == nil {
		err := os.Remove(outfile)
		if err != nil {
			log.Fatal(err)
		}
	}

	var f *os.File
	f, err = os.Create(outfile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = ppmImg.PPMImgWriter(f)
	if err != nil {
		fmt.Println("error writing image to file!", err)
	}
	fmt.Println("done")
}

func RayColor(r *vec3.Ray) *vec3.Vec3 {
	unitVec := vec3.UnitVector(&r.Direction)
	t := 0.5 * (unitVec.Y() + 1.0)

	tmpVec1 := vec3.NewVec3(1.0, 1.0, 1.0)
	tmpVec1.MultBy(1.0 - t)

	tmpVec2 := vec3.NewVec3(0.5, 0.7, 1.0)
	tmpVec2.MultBy(t)

	return vec3.VectorAdd(tmpVec1, tmpVec2)
}

func main() {

}
