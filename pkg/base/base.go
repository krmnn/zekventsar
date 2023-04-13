package base

import (
	"fmt"
	"math/rand"
	"time"

	midicontext "github.com/krmnn/zekventsar/pkg/midi"
)

type Note struct {
	Value int
}

func (n *Note) Print() {
	fmt.Printf("note={midiValue: %v}\n", n.Value)
}

type Clip struct {
	bars     int
	data     []Note
	iterator func() Note
}

func NewDefaultClip() Clip {
	clip := Clip{}
	clip.Init(16, 4)
	return clip
}

func (c *Clip) Init(steps int, bars int) {
	c.data = make([]Note, steps)
	c.bars = bars
	c.iterator = c.getIterator()
}

func (c *Clip) Randomize() {
	for i := 0; i < c.Steps(); i++ {
		c.data[i] = Note{Value: rand.Intn(100)}
	}
}

func (c *Clip) Steps() int {
	return len(c.data)
}

func (c *Clip) Bars() int {
	return c.bars
}

func (c *Clip) Print() {
	fmt.Printf("clip={steps: %v, bars: %v, data: %v}\n", c.Steps(), c.Bars(), c.data)
	for i := 0; i < c.Steps(); i++ {
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

	// FIXME
	m := midicontext.MidiContext{}
	m.Init()

	beat_duration_ms := 60.0 * 1000 / float64(s.Bpm)
	note_duration_ms := 150.0 // FIXME: user param

	fmt.Printf("play() @ %vbpm, %vms per beat\n", s.Bpm, beat_duration_ms)

	ticker := time.NewTicker(time.Duration(beat_duration_ms) * time.Millisecond)
	i := 0
	for range ticker.C {
		if i < c.Steps() {
			cur := c.next()
			fmt.Printf("%v ", cur.Value)
			go m.Send(uint8(cur.Value), note_duration_ms)
			i++
		} else {
			break
		}
	}

	m.Panic()
}

// close over current position
func (c *Clip) getIterator() func() Note {
	pos := 0

	return func() Note {
		fmt.Printf("next() pos: %v\n", pos)
		next := c.data[pos]
		next.Print()
		if pos < c.Steps()-1 {
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
