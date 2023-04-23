package zekventsar

import (
	"fmt"
	"time"
)

type Sequencer struct {
	Bpm     float64
	Running bool
	Loop    bool

	Pos uint8

	beatDurationMs float64
	midiCtx        *MidiContext

	clip   *Clip
	ticker *time.Ticker
	done   chan bool
}

func NewSequencer(bpm float64) Sequencer {
	sequencer := Sequencer{Bpm: bpm}
	sequencer.Init()
	return sequencer
}

func (sequencer *Sequencer) Init() {
	sequencer.midiCtx = &MidiContext{}
	sequencer.midiCtx.Init()
	sequencer.midiCtx.Panic()
	sequencer.beatDurationMs = 60.0 * 1000 / sequencer.Bpm
	sequencer.Running = false
	sequencer.Loop = true
	sequencer.done = make(chan bool)

}

func (sequencer *Sequencer) Load(clip Clip) {
	sequencer.clip = &clip
	sequencer.Running = false

}

func (sequencer *Sequencer) Play() {
	// fmt.Printf("play() @ %vbpm, %vms per beat\n", sequencer.Bpm, sequencer.beatDurationMs)

	sequencer.ticker = time.NewTicker(time.Duration(sequencer.beatDurationMs) * time.Millisecond)
	sequencer.Running = true

	go func() {
		i := uint8(0)
		for {
			select {
			case <-sequencer.done:
				fmt.Printf("end clip\n")
				sequencer.Running = false
				return
			case t := <-sequencer.ticker.C:
				fmt.Println("Current time: ", t)
				note, velocity := sequencer.clip.Next()
				sequencer.Pos = i
				go sequencer.midiCtx.Send(note.Value(), velocity)

				if i < sequencer.clip.Steps-1 {
					i++
				} else {
					i = 0
					if !sequencer.Loop {
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
	sequencer.Running = false
}
