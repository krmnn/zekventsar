package zekventsar

import (
	"fmt"
	"time"

	"gitlab.com/gomidi/midi/v2"
	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv"
)

type MidiContext struct {
	channel uint8
	send    func(msg midi.Message) error
}

func (m *MidiContext) Init() {

	m.channel = 0

	// fmt.Println(midi.GetOutPorts())
	var out, _ = midi.OutPort(0)
	m.send, _ = midi.SendTo(out)
	// defer midi.CloseDriver()
}

func (m *MidiContext) Panic() {
	for i := 0; i < 128; i++ {
		m.send(midi.NoteOff(m.channel, uint8(i)))
	}
}

func (m *MidiContext) Send(note uint8, duration_ms float64) {
	error := m.send(midi.NoteOn(m.channel, note, 100))
	if error != nil {
		fmt.Println(error.Error())
	}
	time.Sleep(time.Duration(duration_ms) * time.Millisecond)

	error = m.send(midi.NoteOff(m.channel, note))
	if error != nil {
		fmt.Println(error.Error())
	}
}
