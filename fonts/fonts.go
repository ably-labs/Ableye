package fonts

import (
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"golang.org/x/image/font"
	"log"
)

const (
	fontDpi = 72
)

var (
	MplusNormalFont font.Face
	MplusSmallFont  font.Face
)

func init() {

	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	MplusNormalFont = truetype.NewFace(tt, &truetype.Options{
		Size:    24,
		DPI:     fontDpi,
		Hinting: font.HintingFull,
	})

	MplusSmallFont = truetype.NewFace(tt, &truetype.Options{
		Size:    16,
		DPI:     fontDpi,
		Hinting: font.HintingFull,
	})
}
