package mknote

import (
	"bytes"
	"errors"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/tiff"
)

var (
	// Ricoh is an exif.Parser for canon makernote data.
	RicohTheta = &ricohTheta{}
)

type ricohTheta struct{}

// Parse decodes all Ricoh makernote data found in x and adds it to x.
func (_ *ricohTheta) Parse(x *exif.Exif) error {
	m, err := x.Get(exif.MakerNote)
	if err != nil {
		return nil
	} else if bytes.Compare(m.Val[:5], []byte("Ricoh")) != 0 {
		return nil
	}
	r := bytes.NewReader(m.Val[8:])
	mkNotesDir, _, err := tiff.DecodeDir(r, x.Tiff.Order)
	if err != nil {
		return err
	}
	x.LoadTags(mkNotesDir, makerNoteRicohFields, false)

	// SubIFD1 0x2001 header:"[Ricoh Camera Info]"+1 = 20byte
	// TODO: implement

	// SubIFD2 0x4001 header:none 0byte
	err = loadSubDir(x, SubIFD2, 0, makerNoteThetaFields)
	if err != nil {
		return err
	}
	return nil
}

func loadSubDir(x *exif.Exif, ptr exif.FieldName, headerByte int64, fieldMap map[uint16]exif.FieldName) error {
	r := bytes.NewReader(x.Raw)

	tag, err := x.Get(ptr)
	if err != nil {
		return nil
	}
	offset := tag.Int(0)

	_, err = r.Seek(offset+headerByte, 0)
	if err != nil {
		return errors.New("exif: seek to sub-IFD failed: " + err.Error())
	}
	subDir, _, err := tiff.DecodeDir(r, x.Tiff.Order)

	if err != nil {
		return errors.New("exif: sub-IFD decode failed: " + err.Error())
	}

	x.LoadTags(subDir, fieldMap, false)
	return nil
}
