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

var inputfile = flag.String("f", "", "input file(jpg)")
var outputjson = flag.String("o", "", "output file(json)")
var verbose = flag.String("v", "", "verbose")

func main() {
	flag.Parse()

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
	//fmt.Println(x.String())
}
