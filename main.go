package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	picdir := os.Getenv("PIC_DIR")
	fmt.Println(filepath.Glob(picdir + "/*"))

}
