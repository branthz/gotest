package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type dirpath struct {
	src string
}

func (d *dirpath) zipfile(path string, f os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if !f.Mode().IsRegular() || f.Size() == 0 {
		return nil
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()

}

func main() {
	err := filepath.Walk()
	if err != nil {
		fmt.Println(err)
		return
	}
}
