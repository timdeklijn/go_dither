package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"log"
	"os"
)

func readImageFromFilePath(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	image, imageType, err := image.Decode(f)
	log.Printf("Image Type: %v\n", imageType)
	return image, err
}

func createGrayImageFromImage(img image.Image) image.Image {
	grayImg := image.NewGray(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			grayImg.Set(x, y, img.At(x, y))
		}
	}
	return grayImg
}

func writeImageToFile(filePath string, img image.Image) error {
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = png.Encode(f, img)

	return err
}

func main() {
	fmt.Println("Hello World")
	img, err := readImageFromFilePath("elephant.jpg")
	if err != nil {
		log.Fatalf("Error loading image: %e\n", err)
	}
	log.Println("Image Loaded")

	grayImg := createGrayImageFromImage(img)

	copyGray := image.NewGray(grayImg.Bounds())
	for y := copyGray.Bounds().Min.Y; y < copyGray.Bounds().Max.Y; y++ {
		for x := copyGray.Bounds().Min.X; x < copyGray.Bounds().Max.X; x++ {
			pixel := grayImg.At(x, y)
			r, _, _, _ := pixel.RGBA()
			newColor := (float64(r) / 65535.0) * 255
			copyGray.Set(x, y, color.Gray{uint8(newColor)})
		}
	}

	for y := copyGray.Bounds().Min.Y; y < copyGray.Bounds().Max.Y; y++ {
		for x := copyGray.Bounds().Min.X; x < copyGray.Bounds().Max.X; x++ {

			oldPixel := copyGray.GrayAt(x, y).Y
			newPixel := 0
			if oldPixel > 122 {
				newPixel = 255
			}
			copyGray.Set(x, y, color.Gray{uint8(newPixel)})

			quantError := int(oldPixel) - newPixel

			tmpPix := float64(copyGray.GrayAt(x+1, y).Y)
			tmpColor := (tmpPix + float64(quantError*7/16))
			copyGray.Set(x+1, y, color.Gray{uint8(tmpColor)})

			tmpPix = float64(copyGray.GrayAt(x-1, y+1).Y)
			tmpColor = (tmpPix + float64(quantError*3/16))
			copyGray.Set(x-1, y+1, color.Gray{uint8(tmpColor)})

			tmpPix = float64(copyGray.GrayAt(x, y+1).Y)
			tmpColor = (tmpPix + float64(quantError*5/16))
			copyGray.Set(x, y+1, color.Gray{uint8(tmpColor)})

			tmpPix = float64(copyGray.GrayAt(x+1, y+1).Y)
			tmpColor = (tmpPix + float64(quantError*1/16))
			copyGray.Set(x+1, y+1, color.Gray{uint8(tmpColor)})
		}
	}

	err = writeImageToFile("tst_gray.png", copyGray)
	if err != nil {
		log.Fatalf("Error writing image: %e\n", err)
	}
}
