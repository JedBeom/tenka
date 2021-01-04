package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func playMusic(bs []Block, mm MusicMeta, lang string) {
	f, err := os.Open("musics/02. Love Addiction.mp3")
	if err != nil {
		log.Fatalln(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatalln(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	mutex := &sync.Mutex{}

	fmt.Printf("%s - %s [%s]\n\n", mm.Artist, mm.Titles[lang], mm.Duration)

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	i := 0
	calibration := +500
	for {
		select {
		case <-done:
			log.Println("music ended")
			return
		case <-time.After(time.Millisecond * 10):
			if i == len(bs) {
				continue
			}
			speaker.Lock()
			pos := format.SampleRate.D(streamer.Position()).Milliseconds()
			speaker.Unlock()

			mutex.Lock()
			if int64(bs[i].ms+calibration) <= pos {
				fmt.Printf("%s: %s\n", mm.Singers[bs[i].Code][lang], bs[i].Contents[lang])
				i++
			}
			mutex.Unlock()
		}
	}

}
