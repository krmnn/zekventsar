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
	bpmSlider := widget.NewSliderWithData(0, 300.0, binding.BindFloat(&sequencer.Bpm))

	// widgets with data binding

	button := widget.NewButton("randomize!", func() {
		log.Println("randomize!")
		clip.Randomize()
	})
	content := container.New(layout.NewHBoxLayout(), label, bpmSlider, button, layout.NewSpacer())
	mainWindow.SetContent(content)

	go func() {
		for range time.Tick(time.Second) {
			notes, _ := boundClip.GetValue("Notes")
			velocities, _ := boundClip.GetValue("Velocities")

			pos, _ := boundSequencer.GetValue("Pos")

			log.Printf("notes      = %v", notes)
			log.Printf("velocities = %v", velocities)

			label.SetText(fmt.Sprintf("%v\npos: %v", notes, pos))
		}
	}()

	mainWindow.Show()

}

func tidyUp() {
	fmt.Println("Exited")
}
