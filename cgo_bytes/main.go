package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"strconv"
	"time"
	"unsafe"
)

/*
   #include <stdio.h>
   #include <stdlib.h>
   #include <stdbool.h>
   #include <string.h>

   extern void go_callback(char **imgData, size_t imgSize);

   static const int depth = 1; // 1byte
   static const int channel = 4; // rgba

   static const char* createEmptyImage(int width, int height, int value) {
   	size_t imgSize = width * height * depth * channel;
   	char* imgData = (char*)malloc(sizeof(char) * imgSize);
   	memset(imgData, value, imgSize);
   	return (const char*)imgData;
   }

   static const char** CreateImage(int width, int height, int value, size_t *imageSize) {
   	const char* imgData = createEmptyImage(width, height, value);
   	*imageSize = sizeof(imgData) / sizeof(*imgData);
   	return (const char**)&imgData;
   }

   static void CreateImageByCallback(int width, int height, int value) {
   	char* imgData = createEmptyImage(width, height, value);
   	size_t imgSize = width * height * depth * channel;
   	go_callback(&imgData, imgSize);
   }

   static void FreeImage(char **imagePtr) {
   	free(imagePtr[0]);
   }

*/
import "C"

var (
	curCount int
	imgBytes []byte
)

const (
	width  = 1920
	height = 1080
)

func createImage() {
	randomValue := time.Now().Nanosecond() % 255

	// var imageSize C.size_t
	// imagePtr := C.CreateImage(C.int(width), C.int(height), C.int(randomValue), &imageSize)
	// imgBytes = C.GoBytes(unsafe.Pointer(*imagePtr), C.int(imageSize))
	// defer C.FreeImage(imagePtr)

	C.CreateImageByCallback(C.int(width), C.int(height), C.int(randomValue))
	if len(imgBytes) <= 0 {
		fmt.Println("empty")
		return
	}

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
				imgBytes[x*channel+(y*stride)+0],
				imgBytes[x*channel+(y*stride)+1],
				imgBytes[x*channel+(y*stride)+2],
				imgBytes[x*channel+(y*stride)+3],
			})
		}
	}

	out, _ := os.Create(fmt.Sprintf("image_%v.jpeg", curCount))
	err := jpeg.Encode(out, goImg, nil)
	if err != nil {
		panic(err)
	}
}

//export go_callback
func go_callback(imgData **C.char, imgSize C.size_t) {
	imgBytes = C.GoBytes(unsafe.Pointer(*imgData), C.int(imgSize))
	defer C.FreeImage(imgData)
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("should specify count (e.g ./[out] 10)")
		return
	}

	startTime := time.Now()

	count, _ := strconv.ParseInt(os.Args[1], 0, 64)
	for i := 0; i < int(count); i++ {
		curCount = i
		createImage()

		time.Sleep(100 * time.Millisecond)
		fmt.Println("tick .. ", time.Since(startTime))
	}

}
