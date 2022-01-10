package button

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

// Button represents a button.
type Button struct {
	width       int
	height      int
	text        string
	textOffsetX int
	textOffsetY int
	textColour  *color.NRGBA
	font        font.Face
	BgColour    *color.NRGBA
	x           int
	y           int
	Image       *ebiten.Image
}

// Draw is used to draw a button to the screen.
func (b *Button) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(b.x), float64(b.y))
	screen.DrawImage(b.Image, opts)
	text.Draw(screen, b.text, b.font, b.x+b.textOffsetX, b.y+b.textOffsetY, b.textColour)
}

// SetBgColour is used to change the background colour of a button.
func (b *Button) SetBgColour(colour *color.NRGBA) {
	b.BgColour = colour
	b.Image.Fill(colour)
}

// IsMouseOver returns true if the mouse cursor is over the button image.
func (b *Button) IsMouseOver() bool {

	mouseX, mouseY := ebiten.CursorPosition()
	// check if mouse X is over the image
	xIsOver := mouseX >= b.x && mouseX <= (b.x+b.width)
	// check if mouse Y is over the image
	yIsOver := mouseY >= b.y && mouseY <= (b.y+b.height)

	return xIsOver && yIsOver
}

// NewButton is a contructor that creates a new button.
func NewButton(width int, height int, text string, textOffsetX int, textOffsetY int, textColour *color.NRGBA, font font.Face, bgColour *color.NRGBA, x int, y int) Button {
	img := ebiten.NewImage(width, height)
	img.Fill(bgColour)
	return Button{
		width:       width,
		height:      height,
		text:        text,
		textOffsetX: textOffsetX,
		textOffsetY: textOffsetY,
		textColour:  textColour,
		font:        font,
		BgColour:    bgColour,
		x:           x,
		y:           y,
		Image:       img,
	}
}
