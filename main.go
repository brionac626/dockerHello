package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	picdir := os.Getenv("PIC_DIR")

	err := filepath.Walk(picdir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			fmt.Println(path)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
