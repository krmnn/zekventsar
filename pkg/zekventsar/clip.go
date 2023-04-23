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

	Bars  uint8
	Steps uint8

	Pos uint8

	iterator func() (midi.Note, uint8)
}

func NewClip() Clip {
	clip := Clip{}
	clip.Init(16, 4, true)
	return clip
}

func (clip *Clip) Init(steps uint8, bars uint8, loop bool) {
	clip.Steps = steps
	clip.Notes = make([]midi.Note, steps)
	clip.Velocities = make([]uint8, steps)

	clip.Bars = bars
	clip.Pos = 0
	clip.iterator = func() (midi.Note, uint8) {
		// fmt.Printf("next() pos: %v\n", pos)
		next := clip.Notes[clip.Pos]
		// next.Print()
		if clip.Pos < clip.Steps-1 {
			clip.Pos++
		} else {
			clip.Pos = 0
		}
		return next, clip.Velocities[clip.Pos]
	}
}

func (clip *Clip) Randomize() {
	for i := uint8(0); i < clip.Steps; i++ {
		clip.Notes[i] = midi.Note(rand.Intn(127))
		clip.Velocities[i] = uint8(rand.Intn(127))

	}
}

func (clip *Clip) SetStep(position uint8, note midi.Note, velocity uint8) {
	if (uint8(0) <= position) && (position <= clip.Steps) {
		clip.Notes[position] = note
	}
}

func (clip *Clip) PrintSteps() string {
	var sb strings.Builder
	for i := uint8(0); i < clip.Steps; i++ {
		if i == clip.Pos {
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
