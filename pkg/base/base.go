package base

import (
	"fmt"
	"math/rand"
	"time"
)

type Note struct {
	Value int
}

func (n *Note) Print() {
	fmt.Printf("note={midiValue: %v}\n", n.Value)
}

type Clip struct {
	data     []Note
	iterator func() Note
}

func (c *Clip) Init(steps int) {
	c.data = make([]Note, steps)
	c.iterator = c.getIterator()
}

func (c *Clip) Randomize() {
	for i := 0; i < c.len(); i++ {
		c.data[i] = Note{Value: rand.Intn(100)}
	}
}

func (c *Clip) len() int {
	return len(c.data)
}

func (c *Clip) Print() {
	fmt.Printf("clip={steps: %v}\n", c.data)
	for i := 0; i < c.len(); i++ {
		fmt.Printf("%v ", c.data[i])
	}
	fmt.Println()
}

type Sequencer struct {
	Bpm int
}

func (s *Sequencer) Print() {
	fmt.Printf("sequencer={Bpm: %v}\n", s.Bpm)
}

func (s *Sequencer) Play(c Clip) {
	beat_duration_ms := 60.0 * 1000 / float64(s.Bpm)
	fmt.Printf("play() @ %vBpm, %vms per beat\n", s.Bpm, beat_duration_ms)

	for i := 0; i < c.len(); i++ {
		fmt.Printf("%v ", c.next())
		time.Sleep(time.Duration(beat_duration_ms) * time.Millisecond)
	}
}

// close over current position
func (c *Clip) getIterator() func() Note {
	pos := 0

	return func() Note {
		fmt.Printf("next() pos: %v\n", pos)
		next := c.data[pos]
		next.Print()
		if pos < c.len()-1 {
			pos++
		} else {
			pos = 0
		}
		return next
	}
}
func (c *Clip) next() Note {
	return c.iterator()
}
