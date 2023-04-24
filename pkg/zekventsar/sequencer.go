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
	sequencer.Pos = 0
	sequencer.done = make(chan bool)

}

func (sequencer *Sequencer) Load(clip Clip) {
	sequencer.clip = &clip
	sequencer.Pos = 0
	sequencer.Running = false

}

func (sequencer *Sequencer) Play() {
	// fmt.Printf("play() @ %vbpm, %vms per beat\n", sequencer.Bpm, sequencer.beatDurationMs)

	sequencer.ticker = time.NewTicker(time.Duration(sequencer.beatDurationMs) * time.Millisecond)
	sequencer.Running = true

	lastBpm := sequencer.Bpm
	go func() {
		for {
			select {
			case <-sequencer.done:
				fmt.Printf("end clip\n")
				sequencer.Running = false
				return
			case t := <-sequencer.ticker.C:
				fmt.Println("Current time: ", t)
				note, velocity := sequencer.clip.Next()
				go sequencer.midiCtx.Send(note.Value(), velocity)

				if sequencer.Bpm != lastBpm {
					fmt.Printf("adapt bpm to %v period %v\n", sequencer.Bpm, sequencer.beatDurationMs)
					sequencer.SyncBpm(sequencer.Bpm)
					sequencer.ticker.Reset(time.Duration(sequencer.beatDurationMs) * time.Millisecond)
					lastBpm = sequencer.Bpm
				}

				if sequencer.Pos < sequencer.clip.Steps-1 {
					sequencer.Pos++
				} else {
					sequencer.Pos = 0
					if !sequencer.Loop {
						return
					}
				}
			}
		}
	}()
}
func (sequencer *Sequencer) SyncBpm(bpm float64) {
	sequencer.beatDurationMs = 60.0 * 1000 / bpm
	// sequencer.Bpm = bpm data binding!
}
func (sequencer *Sequencer) Pause() {
	sequencer.ticker.Stop()
	sequencer.done <- true
	sequencer.Running = false
}
func (sequencer *Sequencer) Stop() {
	sequencer.ticker.Stop()
	sequencer.done <- true
	sequencer.Running = false
	sequencer.Pos = 0
}
