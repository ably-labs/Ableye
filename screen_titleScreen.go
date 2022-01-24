package main

import (
	"fmt"
	"github.com/ably-labs/Ableye/button"
	colour "github.com/ably-labs/Ableye/colours"
	"github.com/ably-labs/Ableye/config"
	font "github.com/ably-labs/Ableye/fonts"
	"github.com/ably-labs/Ableye/text"

	"github.com/hajimehoshi/ebiten/v2"
)

// The elements of the title screen.
var (
	title          text.Text
	tagline        text.Text
	sdkVersion     text.Text
	realtimeButton button.Button
)

func initialiseTitleScreen() {
	title = text.NewText(titleText, colour.White, font.MplusLargeFont, (screenWidth/2)-115, screenHeight/2)
	tagline = text.NewText(taglineText, colour.White, font.MplusNormalFont, (screenWidth/2)-200, (screenHeight/2)+75)
	sdkVersion = text.NewText(fmt.Sprintf("SDK version : ably-go %s", config.Cfg.AblyGoSDKVersion), colour.White, font.MplusSmallFont, (screenWidth/2)-100, (screenHeight/2)+125)
	realtimeButton = button.NewButton(300, 100, realtimeText, 25, 55, colour.White, font.MplusNormalFont, colour.BrightRed, (screenWidth/2)-150, (screenHeight/2)+200)
}

func drawTitleScreen(screen *ebiten.Image) {
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
