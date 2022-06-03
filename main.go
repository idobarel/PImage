package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"strings"

	"github.com/cheggaaa/pb/v3"
)

func main() {
	var pix []color.Color
	xSize, _ := strconv.Atoi(os.Args[1])
	ySize, _ := strconv.Atoi(os.Args[2])
	length := xSize * ySize
	digits := ""
	var q int = 1
	var r int = 0
	var t int = 1
	var k int = 1
	var n int = 3
	var l int = 3
	counter := 0
	bar := pb.StartNew(length)
	for i := 0; len(digits) <= length; i++ {
		bar.Increment()
		if 4*q+r-t < n*t {
			digits += strconv.Itoa(n)
			counter++
			nr := 10 * (r - n*t)
			n = (int((10 * (3*q + r)) / t)) - 10*n
			q *= 10
			r = nr
		} else {
			nr := (2*q + r) * l
			nn := int((q*(7*k) + 2 + (r * l)) / (t * l))
			q *= k
			t *= l
			l += 2
			k += 1
			n = nn
			r = nr
		}
	}
	bar.Finish()
	fmt.Println("Generated, Converting To Binary...")
	digits = strings.ReplaceAll(digits, "-", "")
	binaryDigits := ""
	bar1 := pb.StartNew(len(digits))
	for i := 0; i < len(digits); i++ {
		bar1.Increment()
		j, _ := strconv.Atoi(string(digits[i]))
		s := strconv.FormatInt(int64(j), 2)
		for len(s) < 4 {
			s += "0"
		}
		binaryDigits += s
	}
	bar1.Finish()
	fmt.Println("Done, Converting To Black Or White.")
	bar2 := pb.StartNew(len(binaryDigits))
	for i := 0; i < len(binaryDigits); i++ {
		bar2.Increment()
		if binaryDigits[i] == '0' {
			pix = append(pix, color.Black)
		} else {
			pix = append(pix, color.White)
		}
	}
	bar2.Finish()
	fmt.Println("Creating The Image.")
	rect := image.Rect(0, 0, xSize, ySize)
	img := image.NewRGBA(rect)
	temp := 0
	for y := 0; y < ySize; y++ {
		for x := 0; x < xSize; x++ {
			img.Set(x, y, pix[temp])
			temp++
		}
	}
	f, _ := os.Create("img.png")
	defer f.Close()
	png.Encode(f, img)
	fmt.Println("Done!")
	fmt.Println("Image Saved As 'img.png'")
}
