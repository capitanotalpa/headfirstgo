package main

import "github.com/capitanotalpa/headfirstgo/gadget"

type Player interface {
	Play(string)
	Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	}
}

func main() {
	player := gadget.TapePlayer{Batteries: "123"}
	recorder := gadget.TapeRecorder{Microphones: 2}
	mixtape := []string{"KENDRICK OMAR", "LIL SHRIMP", "WU TANG CLAM"}
	playList(player, mixtape)
	playList(recorder, mixtape)
	TryOut(recorder)
	TryOut(player)
}
