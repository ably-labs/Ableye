package text

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

// Text represents text which is drawn to the screen
type Text struct {
	text       string
	textColour *color.NRGBA
	font       font.Face
	X          int
	Y          int
}

// Draw is used to draw some text to the screen.
func (t *Text) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(t.X), float64(t.Y))
	text.Draw(screen, t.text, t.font, t.X, t.Y, t.textColour)
}

// SetText is used to change the text string.
func (t *Text) SetText(text string) {
	t.text = text
}

// SetX is used to change the X location of the text string.
func (t *Text) SetX(x int) {
	t.X = x
}

// SetYis used to change the Y location of the text string.
func (t *Text) SetY(y int) {
	t.Y = y
}

// NewText is a contructor that creates new text.
func NewText(text string, textColour *color.NRGBA, font font.Face, x int, y int) Text {
	return Text{
		text:       text,
		textColour: textColour,
		font:       font,
		X:          x,
		Y:          y,
	}
}
