package main

import (
	"fmt"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/krmnn/zekventsar/pkg/zekventsar"
)

func main() {
	log.Println("zekventsar v0.1")

	defaultClip := zekventsar.NewClip()
	defaultClip.Randomize()

	defaultSequencer := zekventsar.NewSequencer(180)
	defaultSequencer.Play(defaultClip)

	mainApp := app.New()
	mainWindow(mainApp, defaultSequencer, defaultClip)
	mainApp.Run()
	tidyUp()
}

func mainWindow(app fyne.App, sequencer zekventsar.Sequencer, clip zekventsar.Clip) {
	// setup data bindings
	boundClip := binding.BindStruct(&clip)
	boundSequencer := binding.BindStruct(&sequencer)

	mainWindow := app.NewWindow("zekventsar v0.1")

	label := widget.NewLabel("zekventsar v0.1")
	statusLabel := widget.NewLabel("0 / 0")
	bpmLabel := widget.NewLabel("100 bpm")
	bpmSlider := widget.NewSliderWithData(0, 300.0, binding.BindFloat(&sequencer.Bpm))

	// widgets with data binding
	stopButton := widget.NewButton("stop", func() {
		log.Println("stop!")
	})
	playButton := widget.NewButton("play", func() {
		log.Println("play!")
	})
	transport := container.New(layout.NewHBoxLayout(), stopButton, playButton)

	button := widget.NewButton("randomize!", func() {
		log.Println("randomize!")
		clip.Randomize()
	})
	tools := container.New(layout.NewVBoxLayout(), transport, button, bpmLabel, bpmSlider)
	clipView := container.New(layout.NewVBoxLayout(), label, statusLabel)
	content := container.New(layout.NewHBoxLayout(), clipView, tools)

	mainWindow.SetContent(content)

	go func() {
		for range time.Tick(time.Second) {
			notes, _ := boundClip.GetValue("Notes")
			velocities, _ := boundClip.GetValue("Velocities")
			pos, _ := boundSequencer.GetValue("Pos")
			log.Printf("notes      = %v", notes)
			log.Printf("velocities = %v", velocities)
			label.SetText(fmt.Sprintf("%v\npos: %v", notes, pos))

			bars, _ := boundClip.GetValue("Bars")
			steps, _ := boundClip.GetValue("Steps")
			statusLabel.SetText(fmt.Sprintf("%v / %v", bars, steps))

			bpm, _ := boundSequencer.GetValue("Bpm")
			bpmLabel.SetText(fmt.Sprintf("%v bpm", bpm))
		}
	}()

	mainWindow.Show()

}

func tidyUp() {
	fmt.Println("Exited")
}
