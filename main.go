package main

import (
	"bytes"
	"image"
	_ "image/jpeg"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"

	"github.com/ably-labs/rosie-demo/button"
	colour "github.com/ably-labs/rosie-demo/colours"
	"github.com/ably-labs/rosie-demo/config"
)

var (
	gophersImage *ebiten.Image
)

type App struct {
	count int
}

//Update updates the logical state.
func (a *App) Update() error {

	a.count++

	return nil
}

//Draw renders the screen.
func (a *App) Draw(screen *ebiten.Image) {

	//w, h := gophersImage.Size()
	//op := &ebiten.DrawImageOptions{}
	//op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	//op.GeoM.Rotate(float64(a.count%360) * 2 * math.Pi / 360)
	//op.GeoM.Translate(screenWidth/2, screenHeight/2)
	//screen.DrawImage(gophersImage, op)

	//Draw debug elements if debug mode is on
	if config.Cfg.DebugMode {
		drawDebugText(screen)
	}

	// When the "left mouse button" is pressed...
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		ebitenutil.DebugPrint(screen, "You're pressing the 'LEFT' mouse button.")
	}
	// When the "right mouse button" is pressed...
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		ebitenutil.DebugPrint(screen, "\nYou're pressing the 'RIGHT' mouse button.")
	}
	// When the "middle mouse button" is pressed...
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonMiddle) {
		ebitenutil.DebugPrint(screen, "\n\nYou're pressing the 'MIDDLE' mouse button.")
	}
	// When the "up arrow key" is pressed..
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		ebitenutil.DebugPrint(screen, "\nYou're pressing the 'UP' button.")
	}
	// When the "down arrow key" is pressed..
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		ebitenutil.DebugPrint(screen, "\n\nYou're pressing the 'DOWN' button.")
	}
	// When the "left arrow key" is pressed..
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		ebitenutil.DebugPrint(screen, "\n\n\nYou're pressing the 'LEFT' button.")
	}
	// When the "right arrow key" is pressed..
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		ebitenutil.DebugPrint(screen, "\n\n\n\nYou're pressing the 'RIGHT' button.")
	}

	realtimeButton := button.NewButton(200, 100, "Ably Realtime", 25, 50, colour.Red, screenWidth/4, screenHeight/2)
	realtimeButton.Draw(screen)

	restButton := button.NewButton(200, 100, "Ably Rest", 35, 50, colour.Red, (screenWidth/4)+(screenWidth/3), screenHeight/2)
	restButton.Draw(screen)

}

//Layout returns the logical screen size, the screen is automatically scaled.
func (a *App) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	// Decode an image from the image file's byte slice.
	// Now the byte slice is generated with //go:generate for Go 1.15 or older.
	// If you use Go 1.16 or newer, it is strongly recommended to use //go:embed to embed the image file.
	// See https://pkg.go.dev/embed for more details.
	img, _, err := image.Decode(bytes.NewReader(images.Gophers_jpg))
	if err != nil {
		log.Fatal(err)
	}
	gophersImage = ebiten.NewImageFromImage(img)

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle(titleText)
	if err := ebiten.RunGame(&App{}); err != nil {
		log.Fatal(err)
	}
}
