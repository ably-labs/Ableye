package fonts

import (
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"golang.org/x/image/font"
)

const (
	fontDpi = 72
)

var (
	MplusSmallFont  font.Face
	MplusNormalFont font.Face
	MplusLargeFont  font.Face
)

func init() {

	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	MplusSmallFont = truetype.NewFace(tt, &truetype.Options{
		Size:    16,
		DPI:     fontDpi,
		Hinting: font.HintingFull,
	})

	MplusNormalFont = truetype.NewFace(tt, &truetype.Options{
		Size:    24,
		DPI:     fontDpi,
		Hinting: font.HintingFull,
	})

	MplusLargeFont = truetype.NewFace(tt, &truetype.Options{
		Size:    72,
		DPI:     fontDpi,
		Hinting: font.HintingFull,
	})
}
