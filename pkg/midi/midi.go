package midicontext

import (
	"fmt"

	"github.com/krmnn/zekventsar/pkg/base"
	"gitlab.com/gomidi/midi/v2"
)

type midiContext struct {
	send func(msg midi.Message) error
}

func (m *midiContext) Init() {
	fmt.Println(midi.GetOutPorts())
	fmt.Printf("\n\n")

	var out, _ = midi.OutPort(0)
	// var out = OutByName("my synth")

	// creates a sender function to the out port
	m.send, _ = midi.SendTo(out)

	defer midi.CloseDriver()

}

func (m *midiContext) Send(note base.Note) { // send some messages
	// channel, key, velocity
	m.send(midi.NoteOn(0, midi.Db(5), 100))
	m.send(midi.NoteOff(0, midi.Db(5)))

}
