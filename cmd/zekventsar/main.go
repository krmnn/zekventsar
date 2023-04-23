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
	// defaultClip.Init(1, 4, false)
	defaultClip.Randomize()

	defaultSequencer := zekventsar.NewSequencer(180)
	defaultSequencer.Load(defaultClip)

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

	clipView := container.New(layout.NewGridLayout(8))
	options := zekventsar.GetNoteStrings()
	for i := uint8(0); i < clip.Steps; i++ {
		new := widget.NewSelect(options, func(selected string) {
			log.Println(selected, " selected!")

		})
		log.Println("want to select: ", clip.Notes[i].String())
		new.Alignment = fyne.TextAlignCenter
		new.SetSelected(clip.Notes[i].String())

		velocitySlider := widget.NewSlider(0, 127.0)
		velocitySlider.Orientation = widget.Vertical
		velocitySlider.Value = float64(clip.Velocities[i])
		velocitySlider.Step = 1

		stepBox := container.New(layout.NewVBoxLayout(), new, velocitySlider)

		clipView.Add(stepBox)

	}

	statusLabel := widget.NewLabel("0 / 0")
	bpmLabel := widget.NewLabel("100 bpm")
	bpmSlider := widget.NewSliderWithData(0, 400.0, binding.BindFloat(&sequencer.Bpm))
	stopButton := widget.NewButton("stop", func() {
		log.Println("stop!")
		sequencer.Stop()

	})
	pauseButton := widget.NewButton("pause", func() {
		log.Println("Pause!")
	})
	playButton := widget.NewButton("play", func() {
		log.Println("play!")
		sequencer.Play()

	})
	transport := container.New(layout.NewHBoxLayout(), stopButton, pauseButton, playButton)

	button := widget.NewButton("randomize!", func() {
		log.Println("randomize!")
		clip.Randomize()
	})
	tools := container.New(layout.NewVBoxLayout(), transport, button, bpmLabel, bpmSlider)

	content := container.New(layout.NewHBoxLayout(), clipView, tools)
	mainWindow.SetContent(content)

	go func() {
		for range time.Tick(time.Second) {
			notes, _ := boundClip.GetValue("Notes")
			velocities, _ := boundClip.GetValue("Velocities")
			// pos, _ := boundSequencer.GetValue("Pos")
			log.Printf("notes      = %v", notes)
			log.Printf("velocities = %v", velocities)
			// label.SetText(fmt.Sprintf("%v\npos: %v", notes, pos))

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
