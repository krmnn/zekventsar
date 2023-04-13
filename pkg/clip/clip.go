package clip

import (
	"fmt"
	"math/rand"
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
	data     []Note
	iterator func() Note
	loop     bool
}

func NewClip() Clip {
	clip := Clip{}
	clip.Init(16, 4, true)
	return clip
}

func (clip *Clip) Init(steps int, bars int, loop bool) {
	clip.data = make([]Note, steps)
	clip.bars = bars
	clip.iterator = clip.getIterator()
	clip.loop = loop

}

func (clip *Clip) Randomize() {
	for i := 0; i < clip.Steps(); i++ {
		clip.data[i] = Note{Value: rand.Intn(100)}
	}
}
func (clip *Clip) SetNote(position int, note Note) {
	if position < clip.Steps() {
		clip.data[position] = note
	}
}

func (clip *Clip) Steps() int {
	return len(clip.data)
}

func (clip *Clip) Bars() int {
	return clip.bars
}

func (clip *Clip) Print() {
	fmt.Printf("clip={steps: %v, bars: %v, data: %v}\n", clip.Steps(), clip.Bars(), clip.data)
	for i := 0; i < clip.Steps(); i++ {
		fmt.Printf("%v ", clip.data[i])
	}
	fmt.Println()
}

func (clip *Clip) getIterator() func() Note {
	pos := 0

	return func() Note {
		// fmt.Printf("next() pos: %v\n", pos)
		next := clip.data[pos]
		next.Print()
		if pos < clip.Steps()-1 {
			pos++
		} else {
			pos = 0
		}
		return next
	}
}
func (clip *Clip) Next() Note {
	return clip.iterator()
}

func (clip *Clip) IsLoop() bool {
	return clip.loop
}
