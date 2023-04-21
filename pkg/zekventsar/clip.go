package zekventsar

import (
	"fmt"
	"math/rand"
	"strings"

	"gitlab.com/gomidi/midi/v2"
)

type Clip struct {
	Notes      []midi.Note
	Velocities []uint8

	bars int
	pos  int
	loop bool

	iterator func() (midi.Note, uint8)
}

func NewClip() Clip {
	clip := Clip{}
	clip.Init(16, 4, true)
	return clip
}

func (clip *Clip) Init(steps int, bars int, loop bool) {
	clip.Notes = make([]midi.Note, steps)
	clip.Velocities = make([]uint8, steps)

	clip.bars = bars
	clip.pos = 0
	clip.iterator = func() (midi.Note, uint8) {
		// fmt.Printf("next() pos: %v\n", pos)
		next := clip.Notes[clip.pos]
		// next.Print()
		if clip.pos < clip.Steps()-1 {
			clip.pos++
		} else {
			clip.pos = 0
		}
		return next, clip.Velocities[clip.pos]
	}
	clip.loop = loop
}

func (clip *Clip) Randomize() {
	for i := 0; i < clip.Steps(); i++ {
		clip.Notes[i] = midi.Note(rand.Intn(127))
		clip.Velocities[i] = uint8(rand.Intn(127))

	}
}
func (clip *Clip) SetNote(position int, note midi.Note) {
	if position < clip.Steps() {
		clip.Notes[position] = note
	}
}

func (clip *Clip) Steps() int {
	return len(clip.Notes)
}

func (clip *Clip) Bars() int {
	return clip.bars
}
func (clip *Clip) Pos() int {
	return clip.pos
}

func (clip *Clip) PrintSteps() string {
	var sb strings.Builder
	for i := 0; i < clip.Steps(); i++ {
		if i == clip.pos {
			sb.WriteString(fmt.Sprintf("[%v] ", clip.Notes[i].Value()))

		} else {
			sb.WriteString(fmt.Sprintf("%v ", clip.Notes[i].Value()))

		}
	}
	fmt.Println(sb.String())
	return sb.String()
}

func (clip *Clip) Next() (midi.Note, uint8) {
	return clip.iterator()
}

func (clip *Clip) IsLoop() bool {
	return clip.loop
}
