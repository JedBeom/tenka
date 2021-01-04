package main

import (
	"io/ioutil"
	"log"

	"github.com/BurntSushi/toml"
)

type MusicMeta struct {
	Title     string
	Languages []string
	Titles    map[string]string

	Artist, Album, Composer, Duration string
	Singers                           map[string]map[string]string
}

func parseMeta(filename string) (mm MusicMeta) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	_, err = toml.Decode(string(file), &mm)
	if err != nil {
		log.Fatal(err)
	}

	return
}
