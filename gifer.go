package main

import (
	"flag"
	"image"
	"image/gif"
	"io/ioutil"
	"log"
	"path"
	"regexp"
)

func main() {
	gifName := flag.String("name", "gifer.gif", "The filename of your gif!")
	folderPath := flag.String("folderpath", ".", "Path to a folder of images")

	matched, err := regexp.MatchString(".gif", *gifName)
	if err != nil || matched == false {
		log.Println("Invalid gifname, the filename must end in .gif")
		log.Fatalln(err)
	}

	flag.Parse()

	files, err := ioutil.ReadDir(*folderPath)
	if err != nil {
		log.Println("Failed to find files in", folderPath)
		log.Fatalln(err)
	}

	fileNames := make([]string, len(files))
	for index, file := range files {
		fileNames[index] = path.Join(*folderPath, file.Name())
	}

	// gifer := &gif.GIF{}
}
