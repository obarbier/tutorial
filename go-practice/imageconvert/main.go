package main

import (
	"fmt"
	"github.com/h2non/bimg"
	"os"
)

func main() {

	buffer, err := bimg.Read("BAP_0138.jpg")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	newImage, err := bimg.NewImage(buffer).Convert(bimg.PNG)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if bimg.NewImage(newImage).Type() == "png" {
		fmt.Fprintln(os.Stderr, "The image was converted into png")
	}

}
