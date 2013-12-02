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
    theta-exif R0010001.JPG

Save to json file:
    theta-exif R0010001.JPG -o exif.json

Display text format exif:
    theta-exif R0010001.JPG -t

Print this usage:
    theta-exif -help
`

func usage() {
	fmt.Fprintln(os.Stderr, usageMessage)
	fmt.Fprintln(os.Stderr, "Flags:")
	flag.PrintDefaults()
	os.Exit(2)
}

var output = flag.String("o", "", "output file(default json format)")
var textmode = flag.Bool("t", false, "output a text format instead of json")

//var verbose = flag.String("v", "", "verbose")

var inputfile = ""

func main() {
	flag.Usage = usage
	flag.Parse()

	// Usage information when no arguments.
	if flag.NFlag() == 0 && flag.NArg() == 0 {
		flag.Usage()
	}

	arg1 := os.Args[1]
	if arg1[0] != '-' {
		inputfile = os.Args[1]
		os.Args = os.Args[1:]
		flag.Parse()
	} else {
		inputfile = flag.Arg(0)
	}

	if inputfile == "" {
		flag.Usage()
	}

	f, err := os.Open(inputfile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	exif.RegisterParsers(mknote.RicohTheta)

	x, err := exif.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	var buf []byte

	if *textmode == true {
		buf = []byte(x.String())
	} else {
		buf, err = x.MarshalJSON()
		if err != nil {
			log.Fatal(err)
		}
	}

	if *output != "" {
		err := ioutil.WriteFile(*output, buf, 0644)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println(string(buf))
	}
}
