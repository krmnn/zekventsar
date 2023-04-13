package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type note struct {
	midiValue int
}

func (n *note) print() {
	fmt.Printf("note={midiValue: %v}\n", n.midiValue)
}

type clip struct {
	data     []note
	iterator func() note
}

func (c *clip) init(steps int) {
	c.data = make([]note, steps)
	c.iterator = c.getIterator()
}

func (c *clip) randomize() {
	for i := 0; i < c.len(); i++ {
		c.data[i] = note{midiValue: rand.Intn(100)}
	}
}

func (c *clip) len() int {
	return len(c.data)
}

func (c *clip) print() {
	fmt.Printf("clip={steps: %v}\n", c.data)
	for i := 0; i < c.len(); i++ {
		fmt.Printf("%v ", c.data[i])
	}
	fmt.Println()
}

type sequencer struct {
	bpm int
}

func (s *sequencer) print() {
	fmt.Printf("sequencer={bpm: %v}\n", s.bpm)
}

func (s *sequencer) play(c clip) {
	beat_duration_ms := 60.0 * 1000 / float64(s.bpm)
	fmt.Printf("play() @ %vbpm, %vms per beat\n", s.bpm, beat_duration_ms)

	for i := 0; i < c.len(); i++ {
		fmt.Printf("%v ", c.next())
		time.Sleep(time.Duration(beat_duration_ms) * time.Millisecond)
	}
}

// close over current position
func (c *clip) getIterator() func() note {
	pos := 0

	return func() note {
		fmt.Printf("next() pos: %v\n", pos)
		next := c.data[pos]
		next.print()
		if pos < c.len()-1 {
			pos++
		} else {
			pos = 0
		}
		return next
	}
}
func (c *clip) next() note {
	next := c.iterator()
	return next
}

func main() {
	log.Println("zekventsar v0.1")

	exampleNote := note{
		midiValue: 45,
	}

	exampleNote.print()

	exampleClip := clip{}
	exampleClip.init(8)
	exampleClip.randomize()
	exampleClip.print()

	exampleSequencer := sequencer{bpm: 96}
	exampleSequencer.print()
	// exampleSequencer.play(exampleClip)

	// exampleSequencer.bpm = 180
	// exampleSequencer.print()

	exampleSequencer.play(exampleClip)

}
