package main

import (
	"log"
	"time"

	clip "github.com/krmnn/zekventsar/pkg/clip"
	sequencer "github.com/krmnn/zekventsar/pkg/sequencer"
)

func main() {
	log.Println("zekventsar v0.1")

	exampleClip := clip.Clip{}
	exampleClip.Init(8, 2, true)
	exampleClip.Randomize()
	exampleClip.Print()

	exampleSequencer := sequencer.Sequencer{Bpm: 96}
	exampleSequencer.Init()
	exampleSequencer.Print()

	exampleSequencer.Play(exampleClip)
	time.Sleep(1 * time.Second)
	exampleSequencer.Play(exampleClip)
	// exampleSequencer.Stop()

	time.Sleep(3 * time.Second)

	exampleClip2 := clip.NewClip()
	exampleClip2.Randomize()
	exampleClip2.Print()

	exampleSequencer2 := sequencer.NewSequencer(180)
	exampleSequencer2.Play(exampleClip2)

	time.Sleep(2 * time.Second)
	exampleSequencer2.Stop()

	time.Sleep(2 * time.Second)

	exampleSequencer2.Play(exampleClip2)
	time.Sleep(30 * time.Second)
	exampleSequencer2.Stop()
}
