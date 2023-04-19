package main

import (
	"log"
	"time"

	"github.com/krmnn/zekventsar/pkg/zekventsar"
)

func main() {
	log.Println("zekventsar v0.1")

	defaultClip := zekventsar.NewClip()
	defaultClip.Randomize()

	defaultSequencer := zekventsar.NewSequencer(180)
	defaultSequencer.Play(defaultClip)

	time.Sleep(5 * time.Second)

	defaultSequencer.Stop()

	time.Sleep(5 * time.Second)

}
