package main

import(
	"github.com/disintegration/imaging"
)

func main(){
	src, err := imaging.Open("/Users/brant/Desktop/123.png")
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	dstImage128 := imaging.Resize(src, 128, 128, imaging.Lanczos)
}
