package service

import (
	_ "embed"
	"fmt"
	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/effects"
	"github.com/gopxl/beep/v2/speaker"
	"github.com/sinshu/go-meltysynth/meltysynth"
	"os"
)

type Sythesizer struct {
	Synth      *meltysynth.Synthesizer
	SampleRate int32
	BufferSize int32
	Streamer   beep.Streamer
}

var PianoPlayer *Sythesizer

func InitSpeaker() {
	if err := speaker.Init(beep.SampleRate(PianoPlayer.SampleRate), int(PianoPlayer.BufferSize)); err != nil {
		fmt.Println(err)
	}
	PianoPlayer.Streamer = beep.StreamerFunc(func(samples [][2]float64) (n int, ok bool) {
		left := make([]float32, len(samples))
		right := make([]float32, len(samples))
		PianoPlayer.Synth.Render(left, right)
		for i := range samples {
			samples[i][0] = float64(left[i])
			samples[i][1] = float64(right[i])
		}
		return len(samples), true
	})
	volumeStreamer := effects.Volume{
		Streamer: PianoPlayer.Streamer,
		Base:     7,
		Volume:   1,
		Silent:   false,
	}
	speaker.Play(&volumeStreamer)
}

func LoadSoundFont(path string, sampleRate, bufferSize int32) {
	sf2, _ := os.Open(path)
	soundfont, err := meltysynth.NewSoundFont(sf2)
	if err != nil {
		fmt.Println(err)
	}
	settings := meltysynth.NewSynthesizerSettings(sampleRate)
	synthesizer, err := meltysynth.NewSynthesizer(soundfont, settings)
	if err != nil {
		fmt.Println(err)
	}
	PianoPlayer = &Sythesizer{
		Synth:      synthesizer,
		SampleRate: sampleRate,
		BufferSize: bufferSize,
	}
}

func InitSoundFont(path string) {
	LoadSoundFont(path, 44100, 1024)
	InitSpeaker()
}

func Keydown(channel, key, velocity int32) {
	PianoPlayer.Synth.NoteOn(channel, key, velocity)
}

func Keyup(channel, key int32) {
	PianoPlayer.Synth.NoteOff(channel, key)
}
