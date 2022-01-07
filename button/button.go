package button

import(
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"

	"github.com/ably-labs/rosie-demo/colours"
	"github.com/ably-labs/rosie-demo/fonts"

)
// Button represents a button.
type Button struct {
	width       int
	height      int
	text        string
	textOffsetX int
	textOffsetY int
	bgColour    *color.NRGBA
	x           int
	y           int
	image       *ebiten.Image
}

// Draw is used to draw a button to the screen
func (b *Button) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(b.x), float64(b.y))
	screen.DrawImage(b.image, opts)
	text.Draw(screen, b.text, fonts.MplusNormalFont, b.x+b.textOffsetX, b.y+b.textOffsetY, colours.White)
}

// NewButton is a contructor that creates a new button
func NewButton(width int, height int, text string, textOffsetX int, textOffsetY int, bgColour *color.NRGBA, x int, y int) Button {
	img := ebiten.NewImage(width, height)
	img.Fill(bgColour)
	return Button{
		width:       width,
		height:      height,
		text:        text,
		textOffsetX: textOffsetX,
		textOffsetY: textOffsetY,
		bgColour:    bgColour,
		x:           x,
		y:           y,
		image:       img,
	}
}