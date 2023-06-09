package zekventsar

import (
	"fmt"

	"gitlab.com/gomidi/midi/v2"
	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv"
)

func GetNoteStrings() []string {
	notes := make([]string, 128)
	notes[0] = "-"
	for i := uint8(1); i < 128; i++ {
		notes[i] = fmt.Sprintf("%v", midi.Note(i).String())
	}
	return notes
}

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

func (m *MidiContext) Send(note uint8, velocity uint8) {
	error := m.send(midi.NoteOn(m.channel, note, velocity))
	if error != nil {
		fmt.Println(error.Error())
	}
	// time.Sleep(time.Duration(duration_ms) * time.Millisecond)

	// error = m.send(midi.NoteOff(m.channel, note))
	// if error != nil {
	// 	fmt.Println(error.Error())
	// }
}
