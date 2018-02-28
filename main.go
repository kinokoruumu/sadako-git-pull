package main

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func run() error {
	f, err := os.Open("sadako.mp3")
	if err != nil {
		return err
	}
	s, format, err := mp3.Decode(f)
	if err != nil {
		return err
	}
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(s)
	select {}
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
