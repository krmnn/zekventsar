package main

import (
	"log"
	"time"

	"github.com/krmnn/zekventsar/pkg/base"
)

func main() {
	log.Println("zekventsar v0.1")

	exampleClip := base.Clip{}
	exampleClip.Init(8, 2, true)
	exampleClip.Randomize()
	exampleClip.Print()

	exampleSequencer := base.Sequencer{Bpm: 96}
	exampleSequencer.Init()
	exampleSequencer.Print()

	exampleSequencer.Play(exampleClip)
	time.Sleep(3 * time.Second)
	exampleSequencer.Stop()

	time.Sleep(3 * time.Second)

	exampleClip2 := base.NewClip()
	exampleClip2.Randomize()
	exampleClip2.Print()

	exampleSequencer2 := base.NewSequencer(180)
	exampleSequencer2.Play(exampleClip2)

	time.Sleep(2 * time.Second)
	exampleSequencer2.Stop()

	time.Sleep(2 * time.Second)

	exampleSequencer2.Play(exampleClip2)
	time.Sleep(30 * time.Second)
}
