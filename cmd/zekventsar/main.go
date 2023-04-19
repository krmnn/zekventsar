package main

import (
	"fmt"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/krmnn/zekventsar/pkg/zekventsar"
)

func main() {
	log.Println("zekventsar v0.1")

	defaultClip := zekventsar.NewClip()
	defaultClip.Randomize()

	defaultSequencer := zekventsar.NewSequencer(180)
	defaultSequencer.Play(defaultClip)
	// time.Sleep(5 * time.Second)
	// defaultSequencer.Stop()

	// time.Sleep(5 * time.Second)
	mainApp := app.New()
	// Show
	mainWindow(mainApp, defaultSequencer, defaultClip)
	// Run
	mainApp.Run()
	tidyUp()
}

func updateLabel(label *widget.Label, sequencer zekventsar.Sequencer, clip zekventsar.Clip) {

	text := clip.PrintSteps()
	pos := sequencer.Pos
	fmt.Println(text + "pos: " + fmt.Sprintf("%v", pos))
	label.SetText(text + "pos: " + fmt.Sprintf("%v", pos))
}

func mainWindow(app fyne.App, sequencer zekventsar.Sequencer, clip zekventsar.Clip) {
	mainWindow := app.NewWindow("zekventsar v0.1")

	label := widget.NewLabel("Hello World!")

	mainWindow.SetContent(label)
	// TODO: update content
	go func() {
		for range time.Tick(time.Second) {
			updateLabel(label, sequencer, clip)
		}
	}()

	mainWindow.Show()

}

func tidyUp() {
	fmt.Println("Exited")
}
