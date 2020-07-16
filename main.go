package main

import (
	"fmt"
	"image"
	"image/color"
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
	img, err := readImageFromFilePath("tst.png")
	if err != nil {
		log.Fatalf("Error loading image: %e\n", err)
	}
	log.Println("Image Loaded")

	grayImg := createGrayImageFromImage(img)

	ditheredImage := image.NewGray(grayImg.Bounds())
	for y := ditheredImage.Bounds().Min.Y; y < ditheredImage.Bounds().Max.Y; y++ {
		for x := ditheredImage.Bounds().Min.X; x < ditheredImage.Bounds().Max.X; x++ {
			oldPixel := grayImg.At(x, y)
			r, _, _, _ := oldPixel.RGBA()
			newPixel := 0
			if float64(r)/65535.0 > 0.5 {
				newPixel = 65535
			}
			newColor := color.Gray{uint8(newPixel)}
			ditheredImage.Set(x, y, newColor)
		}
	}

	err = writeImageToFile("tst_gray.png", ditheredImage)
	if err != nil {
		log.Fatalf("Error writing image: %e\n", err)
	}
}
