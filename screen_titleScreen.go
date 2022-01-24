package main

import (
	"fmt"
	"github.com/ably-labs/Ableye/button"
	colour "github.com/ably-labs/Ableye/colours"
	"github.com/ably-labs/Ableye/config"
	font "github.com/ably-labs/Ableye/fonts"
	"github.com/ably-labs/Ableye/text"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// The elements of the title screen.
var (
	title          text.Text
	tagline        text.Text
	sdkVersion     text.Text
	realtimeButton button.Button
	ablyLogo       *ebiten.Image
)

func initialiseTitleScreen() {
	title = text.NewText(titleText, colour.White, font.MplusLargeFont, (screenWidth/2)-115, screenHeight/2)
	tagline = text.NewText(taglineText, colour.White, font.MplusNormalFont, (screenWidth/2)-200, (screenHeight/2)+75)
	sdkVersion = text.NewText(fmt.Sprintf("SDK version : ably-go %s", config.Cfg.AblyGoSDKVersion), colour.White, font.MplusSmallFont, (screenWidth/2)-100, (screenHeight/2)+125)
	realtimeButton = button.NewButton(320, 100, realtimeText, 25, 55, colour.White, font.MplusNormalFont, colour.BrightRed, (screenWidth/2)-155, (screenHeight/2)+200)

	// initialise images from image files.
	var err error
	ablyLogo, _, err = ebitenutil.NewImageFromFile("./images/Ably.png")
	if err != nil {
		log.Println(err)
	}

}

func drawTitleScreen(screen *ebiten.Image) {

	//Draw images.
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(screenWidth/2-350), float64(screenHeight/2)-315)
	screen.DrawImage(ablyLogo, opts)

	//Draw elements.
	title.Draw(screen)
	tagline.Draw(screen)
	sdkVersion.Draw(screen)
	realtimeButton.Draw(screen)

}

func updateTitleScreen() {
	if realtimeButton.IsMouseOver() {
		realtimeButton.SetBgColour(colour.JazzyPink)
	} else {
		realtimeButton.SetBgColour(colour.BrightRed)
	}

	// Handle mouse click on realtime button
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && realtimeButton.IsMouseOver() {
		state = realtimeScreen
	}
}
