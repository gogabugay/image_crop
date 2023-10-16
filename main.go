package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

func main() {

	fmt.Println("Enter folder name")

	var dir_name string

	fmt.Scanln(&dir_name)

	dirs, err := os.ReadDir(dir_name)
	if err != nil {
		log.Fatal(err)
	}

	dirChange(dir_name)

	fmt.Println("Enter x1")

	var x1 int

	fmt.Scanln(&x1)

	fmt.Println("Enter y1")

	var y1 int

	fmt.Scanln(&y1)

	fmt.Println("Enter x2")

	var x2 int

	fmt.Scanln(&x2)

	fmt.Println("Enter y2")

	var y2 int

	fmt.Scanln(&y2)

	//fmt.Println(dirs)
	for _, dir := range dirs {
		if dir.IsDir() == true {
			dirChange(dir.Name())

			cur_dir, err_dir := os.Getwd()
			if err_dir != nil {
				log.Fatal(err_dir)
			}

			files, err_f := os.ReadDir(cur_dir)
			if err_f != nil {
				log.Fatal(err_f)
			}

			for _, file := range files {
				if file.Name() != ".DS_Store" {
					fmt.Println(file.Name())
					crop(file.Name(), x1, y1, x2, y2)
				}

			}

			os.RemoveAll(cur_dir)

			// И вот тут нада кропать

			dirChange("../")

		}

	}
}

func crop(f string, x1 int, y1 int, x2 int, y2 int) {
	originalImageFile, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	defer originalImageFile.Close()

	originalImage, err := png.Decode(originalImageFile)
	if err != nil {
		panic(err)
	}

	cropSize := image.Rect(x1, y1, x2, y2)

	croppedImage := originalImage.(SubImager).SubImage(cropSize)

	cur_dir, err_dir := os.Getwd()
	if err_dir != nil {
		log.Fatal(err_dir)
	}

	newDir(cur_dir + "-cropped")

	path := filepath.Join(cur_dir+"-cropped", f+"cropped.png")

	croppedImageFile, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	defer croppedImageFile.Close()
	if err := png.Encode(croppedImageFile, croppedImage); err != nil {
		panic(err)
	}
}

func newDir(name string) {
	err := os.Mkdir(name, 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
}

func dirChange(nextdir string) {
	err_chd := os.Chdir(nextdir)
	if err_chd != nil {
		log.Fatal(err_chd)
	}
}
