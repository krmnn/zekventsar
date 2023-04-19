package zekventsar

import (
	"fmt"
	"math/rand"
	"strings"
)

// TODO: replace with midi lib
type Note struct {
	Value int
}

func (note *Note) Print() {
	fmt.Printf("note={midiValue: %v}\n", note.Value)
}

type Clip struct {
	bars     int
	Notes    []Note
	pos      int
	iterator func() Note
	loop     bool
}

func NewClip() Clip {
	clip := Clip{}
	clip.Init(16, 4, true)
	return clip
}

func (clip *Clip) Init(steps int, bars int, loop bool) {
	clip.Notes = make([]Note, steps)
	clip.bars = bars
	clip.pos = 0
	clip.iterator = func() Note {
		// fmt.Printf("next() pos: %v\n", pos)
		next := clip.Notes[clip.pos]
		// next.Print()
		if clip.pos < clip.Steps()-1 {
			clip.pos++
		} else {
			clip.pos = 0
		}
		return next
	}
	clip.loop = loop
}

func (clip *Clip) Randomize() {
	for i := 0; i < clip.Steps(); i++ {
		clip.Notes[i] = Note{Value: rand.Intn(100)}
	}
}
func (clip *Clip) SetNote(position int, note Note) {
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
			sb.WriteString(fmt.Sprintf("[%v] ", clip.Notes[i].Value))

		} else {
			sb.WriteString(fmt.Sprintf("%v ", clip.Notes[i].Value))

		}
	}
	fmt.Println(sb.String())
	return sb.String()
}

func (clip *Clip) Next() Note {
	return clip.iterator()
}

func (clip *Clip) IsLoop() bool {
	return clip.loop
}
