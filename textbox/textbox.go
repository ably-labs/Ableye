package textbox

import (
	"image/color"

	colour "github.com/ably-labs/rosie-demo/colours"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

// TextBox represents a textbox.
type TextBox struct {
	Width                  int
	Height                 int
	text                   string
	maxTextLength          int
	textOffsetX            int
	textOffsetY            int
	textColour             *color.NRGBA
	font                   font.Face
	BgColour               *color.NRGBA
	X                      int
	Y                      int
	Image                  *ebiten.Image
	selectedBorderImage    *ebiten.Image
	unselectedBorderImage  *ebiten.Image
	unselectedBorderColour *color.NRGBA
	borderSize             int
	isSelected             bool
	runes                  []rune
	counter                int // counter is used to make the cursor blink
}

// Draw is used to draw a textbox to the screen.
func (t *TextBox) Draw(screen *ebiten.Image) {

	// Draw the selected or unselected border.
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(t.X), float64(t.Y))
	if t.isSelected {
		screen.DrawImage(t.selectedBorderImage, opts)
	} else {
		screen.DrawImage(t.unselectedBorderImage, opts)
	}

	// Draw the background.
	opts.GeoM.Translate(float64(t.borderSize), float64(t.borderSize))
	screen.DrawImage(t.Image, opts)

	// Draw the text.
	text.Draw(screen, t.text, t.font, t.X+t.textOffsetX, t.Y+t.textOffsetY, t.textColour)

	// Blink the cursor.
	if t.isSelected {
		txt := t.text
		if t.counter%60 < 30 {
			txt += "_"
		}
		text.Draw(screen, txt, t.font, t.X+t.textOffsetX, t.Y+t.textOffsetY, t.textColour)
	}
}

// SetBgColour is used to change the background colour of a text box.
func (t *TextBox) SetBgColour(colour *color.NRGBA) {
	t.BgColour = colour
	t.Image.Fill(colour)
}

// SetText is used to change the text on a text box.
func (t *TextBox) SetText(text string) {
	t.text = text
}

// GetText is used to get the text from a text box.
func (t *TextBox) GetText() string {
	return t.text
}

// SetTextColour is used to change the colour of text in a text box.
func (t *TextBox) SetTextColour(colour *color.NRGBA) {
	t.textColour = colour
}

// SetX is used to change the X value of a text box.
func (t *TextBox) SetX(x int) {
	t.X = x
}

// SetY is used to change the X value of a text box.
func (t *TextBox) SetY(y int) {
	t.Y = y
}

// SetFocus is used to give focus to a text box.
func (t *TextBox) SetFocus() {
	t.isSelected = true
}

// RemoveFocus is used to remove focus from a text box.
func (t *TextBox) RemoveFocus() {
	t.isSelected = false
}

// IsMouseOver returns true if the mouse cursor is over the text box.
func (t *TextBox) IsMouseOver() bool {
	mouseX, mouseY := ebiten.CursorPosition()
	// check if mouse X is over the image
	xIsOver := mouseX >= t.X && mouseX <= (t.X+t.Width)
	// check if mouse Y is over the image
	yIsOver := mouseY >= t.Y && mouseY <= (t.Y+t.Height)

	return xIsOver && yIsOver
}

// NewTextBox is a contructor that creates a new text box.
func NewTextBox(width int, height int, borderSize int, text string, maxTextLength int, textOffsetX int, textOffsetY int, textColour *color.NRGBA, font font.Face, bgColour *color.NRGBA, unselectedBorderColour *color.NRGBA, x int, y int) TextBox {

	selectedBorderImg := ebiten.NewImage(width, height)
	selectedBorderImg.Fill(colour.White)

	unselectedBorderImg := ebiten.NewImage(width, height)
	unselectedBorderImg.Fill(unselectedBorderColour)

	img := ebiten.NewImage(width-(borderSize*2), height-(borderSize*2))
	img.Fill(bgColour)

	return TextBox{
		Width:                  width,
		Height:                 height,
		text:                   text,
		maxTextLength:          maxTextLength,
		textOffsetX:            textOffsetX,
		textOffsetY:            textOffsetY,
		textColour:             textColour,
		font:                   font,
		BgColour:               bgColour,
		X:                      x,
		Y:                      y,
		Image:                  img,
		selectedBorderImage:    selectedBorderImg,
		unselectedBorderImage:  unselectedBorderImg,
		unselectedBorderColour: unselectedBorderColour,
		borderSize:             borderSize,
		isSelected:             false,
	}
}

//Update logic for a text box.
func (t *TextBox) Update() {

	// Can only write in the text box if it is selected
	if t.isSelected {

		// Control the maximum text length.
		if len(t.text) < t.maxTextLength {
			// get the input character
			t.runes = ebiten.AppendInputChars(t.runes[:0])

			// add it to the text
			t.text += string(t.runes)
		}

		// If the backspace key is pressed, remove one character.
		if repeatingKeyPressed(ebiten.KeyBackspace) {
			if len(t.text) >= 1 {
				t.text = t.text[:len(t.text)-1]
			}
		}

		t.counter++
	}
}

// repeatingKeyPressed return true when key is pressed considering the repeat state.
func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)
	duration := inpututil.KeyPressDuration(key)
	if duration == 1 {
		return true
	}
	if duration >= delay && (duration-delay)%interval == 0 {
		return true
	}
	return false
}
