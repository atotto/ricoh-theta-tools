package mknote

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/rwcarlsen/goexif/exif"
)

type expected struct {
	fieldName exif.FieldName
	value     string
}

var exifTests = []struct {
	fname string
	exif  []expected
}{
	{
		fname: "../fixture/Ver1.10.JPG",
		exif: []expected{
			{fieldName: "CompassEs", value: "[\"900/10\"]"},
			{fieldName: "ZenithEs", value: "[\"470/10\",\"-350/10\"]"},
		},
	},
}

func TestThetaPaeser(t *testing.T) {

	for _, tt := range exifTests {
		x := decode(t, tt.fname)
		for n, field := range tt.exif {
			tag, err := x.Get(field.fieldName)
			if err != nil {
				t.Fatal(err)
			}
			data, err := tag.MarshalJSON()
			if err != nil {
				t.Fatal(err)
			}
			actual := string(data)
			if field.value != actual {
				t.Fatalf("%s #%d field:%s got %s, want %s",
					filepath.Base(tt.fname), n, field.fieldName, actual, field.value)
			}
		}
	}
}

func decode(t *testing.T, fname string) *exif.Exif {
	f, err := os.Open(fname)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	exif.RegisterParsers(RicohTheta)

	x, err := exif.Decode(f)
	if err != nil {
		t.Fatal(err)
	}
	return x
}
