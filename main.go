package main

import (
	"fmt"
	"os"
)

func main() {
	videodir := os.Getenv("VIDEO_DIR")
	fmt.Println(videodir)
}
