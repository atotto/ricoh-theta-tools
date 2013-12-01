package main

import (
	"github.com/atotto/ricoh-theta-tools/theta-exif/mknote"
	"github.com/rwcarlsen/goexif/exif"

	"io/ioutil"
	"os"

	"flag"
	"fmt"
	"log"
)

const usageMessage = "" +
	`'theta-exif' is a exif output tool of Ricoh Theta JPG image.

Usage theta-exif:

Display json format exif:
    theta-exif -f R0010001.JPG

Save to json file:
    theta-exif -f R0010001.JPG -o exif.json

Display text format exif:
    theta-exif -t -f R0010001.JPG

Print this usage:
    theta-exif -help
`

func usage() {
	fmt.Fprintln(os.Stderr, usageMessage)
	fmt.Fprintln(os.Stderr, "Flags:")
	flag.PrintDefaults()
	os.Exit(2)
}

var inputfile = flag.String("f", "", "input file(jpeg)")
var outputjson = flag.String("o", "", "output file(json)")
var textmode = flag.Bool("t", false, "output a text format instead of json")

//var verbose = flag.String("v", "", "verbose")

func main() {
	flag.Usage = usage
	flag.Parse()
	if *inputfile == "" {
		flag.Usage()
		os.Exit(2)
	}

	f, err := os.Open(*inputfile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	exif.RegisterParsers(mknote.RicohTheta)

	x, err := exif.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	if *textmode == true {
		fmt.Println(x.String())
		os.Exit(0)
	}

	b, err := x.MarshalJSON()
	if err != nil {
		log.Fatal(err)
	}

	if *outputjson != "" {
		err := ioutil.WriteFile(*outputjson, b, 0644)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println(string(b))
	}
}
