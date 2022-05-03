package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("vim-go")
}

func downloadFromUrl(url string) {
	resource, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resource.Body.Close()

	file, err := os.Create("image.jpg")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.ReadFrom(resource.Body)
}
