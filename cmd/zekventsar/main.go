package main

import (
	"log"
	"time"

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
	exampleSequencer.Play(exampleClip)
	time.Sleep(20 * time.Second)
	// 	exampleClip2 := base.NewDefaultClip()

	// 	exampleClip2.Randomize()
	// 	exampleSequencer2 := base.Sequencer{Bpm: 180}
	// 	exampleSequencer2.Play(exampleClip2)

	// exampleSequencer3 := base.Sequencer{Bpm: 300}
	// exampleSequencer3.Play(exampleClip2)
}
