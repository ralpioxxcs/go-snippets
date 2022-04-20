package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>

char *image;

const char* getString() {
	return "test";
}

const char** getImagePointer(size_t *imageSize, int value) {
	const int width = 1920;
	const int height = 1080;
	const int depth = 1; // 1byte
	const int channel = 4; // rgba

	*imageSize = width * height * depth * channel;
	image = malloc(sizeof(char)* (*imageSize));
	memset(image, value, *imageSize);

	return (const char**)&image;
}

void freeImage(char **imagePtr) {
	free(imagePtr[0]);
	//free(*imagePtr);
}

*/
import "C"
import (
	_ "bytes"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"strconv"
	"time"
	"unsafe"
)

func RetrieveBytes(width, height *int) []byte {
	var imageSize C.size_t
	imagePtr := C.getImagePointer(&imageSize, C.int(time.Now().Nanosecond()%255))

	//fmt.Println("size : ", imageSize)

	*width = 1920
	*height = 1080

	slice := C.GoBytes(unsafe.Pointer(*imagePtr), C.int(imageSize))
	defer C.freeImage(imagePtr)

	return slice
}

var (
	curCount int
)

func createImage() {
	var width, height int
	goImgBytes := RetrieveBytes(&width, &height)

	goImg := image.NewRGBA(
		image.Rectangle{image.Point{0, 0}, image.Point{width, height}},
	)

	var (
		channel = 4
		stride  = width * channel
	)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			goImg.Set(x, y, color.RGBA{
				goImgBytes[x*channel+(y*stride)+0],
				goImgBytes[x*channel+(y*stride)+1],
				goImgBytes[x*channel+(y*stride)+2],
				goImgBytes[x*channel+(y*stride)+3],
			})
		}
	}

	out, _ := os.Create(fmt.Sprintf("image_%v.jpeg", curCount))
	err := jpeg.Encode(out, goImg, nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	startTime := time.Now()

	count, _ := strconv.ParseInt(os.Args[1], 0, 64)
	for i := 0; i < int(count); i++ {
		curCount = i
		createImage()

		time.Sleep(500 * time.Millisecond)
		fmt.Println("tick .. ", time.Since(startTime))
	}

}
