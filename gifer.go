package main

import (
	"flag"
	"image"
	"image/gif"
	"io/ioutil"
	"log"
	"regexp"
)

func main() {
	gifName := flag.String("name", "gifer.gif", "The filename of your gif!")
	folderPath := flag.String("folderPath", ".", "Path to a folder of images")

	matched, err := regexp.MatchString(".gif", *gifName)
	if err != nil || matched == false {
		log.Println("Invalid gifname, the filename must end in .gif")
		log.Fatalln(err)
	}

	flag.Parse()

	// gifer := &gif.GIF{}
}
