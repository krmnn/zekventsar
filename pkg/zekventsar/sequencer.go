package zekventsar

import (
	"time"
)

type Sequencer struct {
	Bpm       int
	IsStarted bool

	beatDurationMs float64
	midiCtx        *MidiContext

	ticker *time.Ticker
	done   chan bool
}

func NewSequencer(bpm int) Sequencer {
	sequencer := Sequencer{Bpm: bpm}
	sequencer.Init()
	return sequencer
}

func (sequencer *Sequencer) Init() {
	sequencer.midiCtx = &MidiContext{}
	sequencer.midiCtx.Init()
	sequencer.midiCtx.Panic()
	sequencer.beatDurationMs = 60.0 * 1000 / float64(sequencer.Bpm)
	sequencer.IsStarted = false
}

func (sequencer *Sequencer) Play(clip Clip) {
	note_duration_ms := 400.0 // TODO: user param
	// fmt.Printf("play() @ %vbpm, %vms per beat\n", sequencer.Bpm, sequencer.beatDurationMs)

	sequencer.ticker = time.NewTicker(time.Duration(sequencer.beatDurationMs) * time.Millisecond)
	sequencer.done = make(chan bool)
	sequencer.IsStarted = true

	go func() {
		i := 0
		for {
			select {
			case <-sequencer.done:
				// fmt.Printf("end clip\n")
				sequencer.IsStarted = false
				return
			case <-sequencer.ticker.C:
				// fmt.Println("Current time: ", t)
				note := clip.Next()

				go sequencer.midiCtx.Send(uint8(note.Value), note_duration_ms)

				if i < clip.Steps()-1 {
					i++
				} else {
					i = 0
					if !clip.IsLoop() {
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
	// fmt.Printf("Ticker stopped\n")
}

func (sequencer *Sequencer) Print() {
	// fmt.Printf("sequencer={Bpm: %v}\n", sequencer.Bpm)
}
