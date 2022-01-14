package button

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

// Button represents a button.
type Button struct {
	Width       int
	Height      int
	text        string
	textOffsetX int
	textOffsetY int
	textColour  *color.NRGBA
	font        font.Face
	BgColour    *color.NRGBA
	X           int
	Y           int
	Image       *ebiten.Image
}

// Draw is used to draw a button to the screen.
func (b *Button) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(b.X), float64(b.Y))
	screen.DrawImage(b.Image, opts)
	text.Draw(screen, b.text, b.font, b.X+b.textOffsetX, b.Y+b.textOffsetY, b.textColour)
}

// SetBgColour is used to change the background colour of a button.
func (b *Button) SetBgColour(colour *color.NRGBA) {
	b.BgColour = colour
	b.Image.Fill(colour)
}

// SetText is used to change the text on a button.
func (b *Button) SetText(text string) {
	b.text = text
}

// SetTextColour is used to change the colour of text on a button.
func (b *Button) SetTextColour(colour *color.NRGBA) {
	b.textColour = colour
}

// SetX is used to change the X value of a button.
func (b *Button) SetX(x int) {
	b.X = x
}

// SetY is used to change the X value of a button.
func (b *Button) SetY(y int) {
	b.Y = y
}

// IsMouseOver returns true if the mouse cursor is over the button image.
func (b *Button) IsMouseOver() bool {
	mouseX, mouseY := ebiten.CursorPosition()
	// check if mouse X is over the image
	xIsOver := mouseX >= b.X && mouseX <= (b.X+b.Width)
	// check if mouse Y is over the image
	yIsOver := mouseY >= b.Y && mouseY <= (b.Y+b.Height)

	return xIsOver && yIsOver
}

// NewButton is a contructor that creates a new button.
func NewButton(width int, height int, text string, textOffsetX int, textOffsetY int, textColour *color.NRGBA, font font.Face, bgColour *color.NRGBA, x int, y int) Button {
	img := ebiten.NewImage(width, height)
	img.Fill(bgColour)
	return Button{
		Width:       width,
		Height:      height,
		text:        text,
		textOffsetX: textOffsetX,
		textOffsetY: textOffsetY,
		textColour:  textColour,
		font:        font,
		BgColour:    bgColour,
		X:           x,
		Y:           y,
		Image:       img,
	}
}
