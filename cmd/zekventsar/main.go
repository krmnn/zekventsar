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

	exampleSequencer.Play(exampleClip)

	exampleClip.Randomize()
	exampleSequencer2 := base.Sequencer{Bpm: 180}
	exampleSequencer2.Play(exampleClip)

	exampleClip3 := base.Clip{}
	exampleClip3.Init(80)
	exampleClip3.Randomize()
	exampleSequencer.Play(exampleClip3)

}
