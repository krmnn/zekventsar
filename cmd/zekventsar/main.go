package main

import (
	"log"
	"github.com/krmnn/zekventsar/pkg/base"
	
)

func main() {
	log.Println("zekventsar v0.1")

	exampleNote := base.Note{
		Value: 45,
	}

	exampleNote.Print()

	exampleClip := base.Clip{}
	exampleClip.Init(8)
	exampleClip.Randomize()
	exampleClip.Print()

	exampleSequencer := base.Sequencer{Bpm: 96}
	exampleSequencer.Print()
	// exampleSequencer.play(exampleClip)

	// exampleSequencer.bpm = 180
	// exampleSequencer.print()

	exampleSequencer.Play(exampleClip)

}
