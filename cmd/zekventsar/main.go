package main

import (
	"log"

	"github.com/krmnn/zekventsar/pkg/base"
)

func main() {
	log.Println("zekventsar v0.1")

	exampleClip := base.Clip{}
	exampleClip.Init(8, 2)
	exampleClip.Randomize()
	exampleClip.Print()

	exampleSequencer := base.Sequencer{Bpm: 96}
	exampleSequencer.Print()
	// exampleSequencer.play(exampleClip)

	exampleSequencer.Play(exampleClip)

	exampleClip.Randomize()
	exampleSequencer2 := base.Sequencer{Bpm: 180}
	exampleSequencer2.Play(exampleClip)

	exampleSequencer3 := base.Sequencer{Bpm: 300}
	exampleSequencer3.Play(exampleClip)

}
