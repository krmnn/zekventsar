package main

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	log.Println("zekventsar v0.1")

	// defaultClip := zekventsar.NewClip()
	// defaultClip.Randomize()

	// defaultSequencer := zekventsar.NewSequencer(180)
	// defaultSequencer.Play(defaultClip)

	// time.Sleep(5 * time.Second)

	// defaultSequencer.Stop()

	// time.Sleep(5 * time.Second)
	mainApp := app.New()
	mainWindow := mainApp.NewWindow("zekventsar v0.1")

	mainWindow.SetContent(widget.NewLabel("Hello World!"))

	mainWindow.Show()

	// TODO: update content

	mainApp.Run()
	tidyUp()
}

func tidyUp() {
	fmt.Println("Exited")
}
