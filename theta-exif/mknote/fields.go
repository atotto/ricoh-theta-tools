package mknote

import (
	"github.com/rwcarlsen/goexif/exif"
)

// ref: http://www.ozhiker.com/electronics/pjmt/jpeg_info/ricoh_mn.html
// ref: http://d.hatena.ne.jp/xanxys/20131110/1384094832

const (
	// fields
	MakernoteDataType      exif.FieldName = "Makernote Data Type"
	Version                               = "Version"
	PrintImageMatchingInfo                = "Print Image Matching Info"

	SubIFD1 = "Ricoh SubIFD1 offset"
	SubIFD2 = "Ricoh SubIFD2 offset"

	// sub
	ZenithEs  = "ZenithEs"  // 0 <= ZenithEs[0] <= 360, -90 <= ZenithEs[1] <= 90
	CompassEs = "CompassEs" // 0 <= CompassEs <= 360

	Unknown0001 = "Theta Unknown Field 0001"
	Unknown0002 = "Theta Unknown Field 0002"
	Unknown0005 = "Theta Unknown Field 0005"
)

var makerNoteRicohFields = map[uint16]exif.FieldName{
	0x0001: MakernoteDataType,
	0x0002: Version,
	0x0E00: PrintImageMatchingInfo,

	0x2001: SubIFD1,
	0x4001: SubIFD2,
}

// for SUbIFD2
var makerNoteThetaFields = map[uint16]exif.FieldName{
	0x0001: Unknown0001,
	0x0002: Unknown0002,
	0x0003: ZenithEs,
	0x0004: CompassEs,
	0x0005: Unknown0005,
	//0x0101: , //ISO Speed Ratings?
	//0x0102: , //FNumber? ApertureValue?
	//0x0103: , //Exposure Time?
	//0x0104: , //Serial No?
	//0x0105: , //Serial No?
}
