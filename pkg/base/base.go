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
func (clip *Clip) next() Note {
	return clip.iterator()
}

type Sequencer struct {
	Bpm int

	beatDurationMs float64
	midiCtx        *midicontext.MidiContext

	ticker *time.Ticker
	done   chan bool
}

func NewSequencer(bpm int) Sequencer {
	sequencer := Sequencer{Bpm: bpm}
	sequencer.Init()
	return sequencer
}

func (sequencer *Sequencer) Init() {
	sequencer.midiCtx = &midicontext.MidiContext{}
	sequencer.midiCtx.Init()
	sequencer.midiCtx.Panic()
	sequencer.beatDurationMs = 60.0 * 1000 / float64(sequencer.Bpm)
}

func (sequencer *Sequencer) Play(clip Clip) {
	note_duration_ms := 400.0 // FIXME: user param
	fmt.Printf("play() @ %vbpm, %vms per beat\n", sequencer.Bpm, sequencer.beatDurationMs)

	sequencer.ticker = time.NewTicker(time.Duration(sequencer.beatDurationMs) * time.Millisecond)
	sequencer.done = make(chan bool)

	go func() {
		i := 0
		for {
			select {
			case <-sequencer.done:
				fmt.Printf("end clip\n")
				return
			case t := <-sequencer.ticker.C:
				fmt.Println("Current time: ", t)
				note := clip.next()

				go sequencer.midiCtx.Send(uint8(note.Value), note_duration_ms)

				if i < clip.Steps()-1 {
					i++
				} else {
					i = 0
					if !clip.loop {
						return
					}
				}
			}
		}
	}()
}

func (sequencer *Sequencer) Stop() {
	sequencer.ticker.Stop()
	sequencer.done <- true
	fmt.Printf("Ticker stopped\n")
}

func (sequencer *Sequencer) Print() {
	fmt.Printf("sequencer={Bpm: %v}\n", sequencer.Bpm)
}
