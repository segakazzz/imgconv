package main

import (
	"fmt"
	"github.com/segakazzz/imgconv/imgconv"
	"log"
	"os"
)

func main(){
	args := os.Args[1:]
	if len(args) == 0{
		log.Fatal("Usage: (command) (dirName)")
	}
	err := imgconv.Convert(args[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done!")
}

