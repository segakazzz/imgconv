package main

import (
	"flag"
	"fmt"
	"github.com/segakazzz/imgconv/imgconv"
	"log"
)


func main(){
	var (
		dir  = flag.String("d", ".", "Indicate directory to convert")
		in  = flag.String("i", "jpg", "Indicate input image file's extension")
		out = flag.String("o", "png", "Indicate output image file's extension")
		err error
	)

	flag.Parse()
	c, err := imgconv.NewConverter(*dir, *in, *out)
	if err != nil {
		log.Fatal(err)
	}
	err = c.Convert()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done!")
}

