package main

import (
	"bytes"
	"flag"
	"github.com/disintegration/imaging"
	"image"
	"image/gif"
	"io/ioutil"
	"log"
	"os"
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

	gifer := &gif.GIF{}

	for _, filename := range fileNames {
		f, err := imaging.Open(filename)
		if err != nil {
			log.Println("Failed to open", filename)
			log.Fatalln(err)
		}

		buf := bytes.Buffer{}
		err = gif.Encode(&buf, f, nil)
		if err != nil {
			log.Println("Failed to decode", filename, "are you sure its an image?")
		}

		tmpImg, _ := gif.Decode(&buf)

		gifer.Image = append(gifer.Image, tmpImg.(*image.Paletted))
		gifer.Delay = append(gifer.Delay, 10)
	}

	f, err := os.OpenFile(*gifName, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	gif.EncodeAll(f, gifer)
}
