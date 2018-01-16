package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/branthz/utarrow/lib/file"
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
	return nil

}

func main(){
	readfileLine()
}

// test readline
func readfileLine() {
	f, err := os.Open("./abc")
	if err != nil {
		fmt.Println(err)
		return
	}
	r:=bufio.NewReader(f)
	for {
		line, err := file.ReadLine(r)
		if err != nil  {
			if err==io.EOF{
				break	
			}
			fmt.Println(err)
			return
		}
		fmt.Printf("%s\n", string(line))
	}
}
