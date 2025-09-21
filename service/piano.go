package service

import (
	"fmt"
	"gitlab.com/gomidi/midi/v2"
	"gitlab.com/gomidi/midi/v2/drivers"
	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv"
	"sync"
	"time"
)

type KeyboardSignal struct {
	Value    uint8 `json:"value"`
	Velocity uint8 `json:"velocity"`
	Channel  uint8 `json:"channel"`
}

type Listener struct {
	Down    chan bool
	Started bool
}

type InMidiDevice struct {
	Device drivers.In `json:"device"`
	Name   string     `json:"name"`
	Value  int        `json:"value"`
}

type OutMidiDevice struct {
	Device drivers.Out `json:"-"`
	Name   string      `json:"name"`
	Value  int         `json:"value"`
}

type PedalSingal struct {
	DeviceID        int             `json:"deviceID"`
	DamperPedal     bool            `json:"damperPedal"`    // 延音踏板 64
	SostenutoPedal  bool            `json:"sostenutoPedal"` // 消音踏板 66
	SoftPedal       bool            `json:"softPedal"`      // 柔音踏板 67
	DamperPedalKeys []uint8         `json:"-"`
	DownKeys        map[uint8]uint8 `json:"-"`
}

type MidiDevices struct {
	InMidiPool        map[int]InMidiDevice  `json:"inMidiPool"`
	OutMidiPool       map[int]OutMidiDevice `json:"outMidiPool"`
	SelectedInDevice  int                   `json:"selectedInDevice"`
	SelectedOutDevice int                   `json:"selectedOutDevice"`
	PedalStatus       map[int]*PedalSingal  `json:"pedalStatus"`
	Listener          Listener              `json:"-"`
	Initialized       bool                  `json:"initialized"`
}

var Midis = MidiDevices{
	InMidiPool: map[int]InMidiDevice{
		-1: {
			Name:  "无",
			Value: -1,
		},
	},
	OutMidiPool: map[int]OutMidiDevice{
		-1: {
			Name:  "无",
			Value: -1,
		},
	},
	SelectedInDevice:  -1,
	SelectedOutDevice: -1,
	PedalStatus:       make(map[int]*PedalSingal),
	Listener: Listener{
		Down:    make(chan bool),
		Started: false,
	},
	Initialized: false,
}

type Keyboard struct{}

func CloseMidiDevice() {
	for i, _ := range Midis.InMidiPool {
		if i != -1 {
			Midis.InMidiPool[i].Device.Close()

		}
	}
	for i, _ := range Midis.OutMidiPool {
		if i != -1 {
			Midis.OutMidiPool[i].Device.Close()

		}
	}
	midi.CloseDriver()
	drivers.Close()
	fmt.Println("midi listener stop")
	App.Quit()
}

func CompareInDevices(inports midi.InPorts) {
	lastId := -1
	for i, _ := range inports {
		if _, ok := Midis.InMidiPool[inports[i].Number()]; ok {
			continue
		} else {
			Midis.InMidiPool[inports[i].Number()] = InMidiDevice{
				Device: inports[i],
				Name:   inports[i].String(),
				Value:  i,
			}
			Midis.PedalStatus[i] = &PedalSingal{
				DeviceID:        i,
				DamperPedal:     false,
				SostenutoPedal:  false,
				SoftPedal:       false,
				DamperPedalKeys: make([]uint8, 0),
				DownKeys:        make(map[uint8]uint8, 0),
			}
		}
		lastId = inports[i].Number()
	}
	if Midis.SelectedInDevice == -1 && len(inports) > 0 {
		Midis.SelectedInDevice = lastId
		go (&Keyboard{}).MidiListenerStart()
	}
	indexArray := make([]int, 0)
	for index, _ := range Midis.InMidiPool {
		indexArray = append(indexArray, index)
	}
	for _, i := range indexArray {
		exsit := false
		for j, _ := range inports {
			if inports[j].Number() == i {
				exsit = true
				break
			}
		}
		if !exsit {
			if i == Midis.SelectedInDevice {
				Midis.SelectedInDevice = -1
			}
			if i != -1 {
				Midis.InMidiPool[i].Device.Close()
				(&Keyboard{}).MidiListenerStop()
				delete(Midis.PedalStatus, i)
				delete(Midis.InMidiPool, i)
			}

		}
	}

}

func CompareOutDevices(outports midi.OutPorts) {
	lastId := -1
	for i, _ := range outports {
		if _, ok := Midis.OutMidiPool[outports[i].Number()]; ok {
			continue
		} else {
			Midis.OutMidiPool[outports[i].Number()] = OutMidiDevice{
				Device: outports[i],
				Name:   outports[i].String(),
				Value:  i,
			}
			outports[i].Open()
		}
		lastId = outports[i].Number()
	}
	if Midis.SelectedOutDevice == -1 && len(outports) > 0 {
		Midis.SelectedOutDevice = lastId
	}
	indexArray := make([]int, 0)
	for index, _ := range Midis.OutMidiPool {
		indexArray = append(indexArray, index)
	}
	for _, i := range indexArray {
		exsit := false
		for j, _ := range outports {
			if outports[j].Number() == i {
				exsit = true
				break
			}
		}
		if !exsit {
			if i == Midis.SelectedOutDevice {
				Midis.SelectedOutDevice = -1
			}
			if i != -1 {
				Midis.OutMidiPool[i].Device.Close()
				delete(Midis.OutMidiPool, i)
			}
		}
	}
}

func ListenMidiDevices() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		CompareInDevices(midi.GetInPorts())
	}()
	go func() {
		defer wg.Done()
		CompareOutDevices(midi.GetOutPorts())
	}()
	wg.Wait()

	App.Event.Emit("devices", Midis)

}

func (k *Keyboard) GetMidiDevices() MidiDevices {
	return Midis
}

func (k *Keyboard) MidiListenerStart() {
	if Midis.SelectedInDevice == -1 {
		return
	}
	fmt.Println("midi listener start")
	stop, err := midi.ListenTo(Midis.InMidiPool[Midis.SelectedInDevice].Device, func(msg midi.Message, timestampms int32) {
		var bt []byte
		var ch, key, vel, con uint8
		switch {
		case msg.GetSysEx(&bt):
			//fmt.Println("got sysex: % X\n", bt)
		case msg.GetNoteStart(&ch, &key, &vel):
			Keydown(int32(ch), int32(key), int32(vel))
			Midis.PedalStatus[Midis.SelectedInDevice].DamperPedalKeys = append(Midis.PedalStatus[Midis.SelectedInDevice].DamperPedalKeys, midi.Note(key).Value())
			Midis.PedalStatus[Midis.SelectedInDevice].DownKeys[midi.Note(key).Value()] = midi.Note(key).Value()
			App.Event.Emit("down", &KeyboardSignal{
				Value:    midi.Note(key).Value(),
				Velocity: vel,
				Channel:  ch,
			})
		case msg.GetControlChange(&ch, &con, &vel):
			if con == 64 {
				Midis.PedalStatus[Midis.SelectedInDevice].DamperPedal = vel > 0
				if !Midis.PedalStatus[Midis.SelectedInDevice].DamperPedal {
					for _, key := range Midis.PedalStatus[Midis.SelectedInDevice].DamperPedalKeys {
						if _, ok := Midis.PedalStatus[Midis.SelectedInDevice].DownKeys[key]; !ok {
							App.Event.Emit("up", &KeyboardSignal{
								Value:    key,
								Velocity: 0,
								Channel:  ch,
							})
						}

					}
					Midis.PedalStatus[Midis.SelectedInDevice].DamperPedalKeys = make([]uint8, 0)
				}
			} else if con == 66 {
				Midis.PedalStatus[Midis.SelectedInDevice].SostenutoPedal = vel > 0
			} else if con == 67 {
				Midis.PedalStatus[Midis.SelectedInDevice].SoftPedal = vel > 0
			}
			if con == 64 || con == 66 || con == 67 {
				App.Event.Emit("pedal", Midis.PedalStatus[Midis.SelectedInDevice])
			}
		case msg.GetNoteEnd(&ch, &key):
			if !Midis.PedalStatus[Midis.SelectedInDevice].DamperPedal {
				App.Event.Emit("up", &KeyboardSignal{
					Value:    midi.Note(key).Value(),
					Velocity: 0,
					Channel:  ch,
				})
				Keyup(int32(ch), int32(key))

			}
			delete(Midis.PedalStatus[Midis.SelectedInDevice].DownKeys, midi.Note(key).Value())
			//fmt.Println(midi.Note(key), midi.Note(key).Value(), midi.Note(key).Octave(), midi.Note(key).Name(), midi.Note(key).Base(), ch)
		default:
			// ignore
		}
	}, midi.UseSysEx())

	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}
	Midis.Listener.Started = true

	go func() {
		<-Midis.Listener.Down
		if len(Midis.InMidiPool) > 1 {
			stop()
		} else {
			stop = nil
		}

		fmt.Println("midi listener stop")
	}()
	<-Midis.Listener.Down
}

func (k *Keyboard) KeyboardPlay(key uint8) {
	//fmt.Println("down:", key)
	Keydown(0, int32(key), UserConfig.Volume)

	if !Midis.Listener.Started || Midis.SelectedOutDevice == -1 {
		return
	}
	noteOn := midi.NoteOn(uint8(Midis.SelectedOutDevice), key, UserConfig.Velocity)
	if err := Midis.OutMidiPool[Midis.SelectedOutDevice].Device.Send(noteOn); err != nil {
		fmt.Println(err)
	}
}

func (k *Keyboard) KeyboardStop(key uint8) {
	Keyup(0, int32(key))

	if !Midis.Listener.Started || Midis.SelectedOutDevice == -1 {
		return
	}
	noteOff := midi.NoteOff(uint8(Midis.SelectedOutDevice), key)
	if err := Midis.OutMidiPool[Midis.SelectedOutDevice].Device.Send(noteOff); err != nil {
		fmt.Println(err)
	}
}

func (k *Keyboard) MidiListenerStop() {
	if !Midis.Listener.Started {
		return
	}
	Midis.Listener.Down <- true
	close(Midis.Listener.Down)
	Midis.Listener.Started = false
	Midis.Listener.Down = make(chan bool)
}

func (k *Keyboard) ChangeDevice(deviceType string, deviceID int) bool {
	if deviceType == "in" && len(Midis.InMidiPool) > deviceID {
		Midis.SelectedInDevice = deviceID
	} else if deviceType == "out" && len(Midis.OutMidiPool) > deviceID {
		Midis.SelectedOutDevice = deviceID
	} else {
		return false
	}
	fmt.Println(Midis)
	return true
}

func ListenDevices() {
	for {
		ListenMidiDevices()
		time.Sleep(3 * time.Second)
	}
}
