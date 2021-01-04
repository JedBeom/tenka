package main

import (
	"io/ioutil"
	"log"
	"strings"
	"unicode/utf8"
)

var (
	lang = []string{"ja", "ja-ko", "ko"}
)

type Block struct {
	Timing   string `json:"timing"`
	ms       int
	Code     string            `json:"code"`
	Contents map[string]string `json:"contents"`
	Append   bool              `json:"append"`
	Stop     bool              `json:"stop"`
}

var musicTitle = "02. Love Addiction"

func main() {
	// filename, lang-code 두개 받자
	fileByte, err := ioutil.ReadFile(musicTitle + ".ttsl")
	if err != nil {
		panic(err)
	}

	file := string(fileByte)
	file = strings.Replace(file, "\r\n", "\n", -1) // for windows

	if file[len(file)-1] == '\n' {
		file = file[:len(file)-1]
	}

	mm := parseMeta(musicTitle + ".toml")
	bs := parse(file)
	playMusic(bs, mm, "ja-ko")
}

func parse(doc string) (blocks []Block) {

	blockStrings := strings.Split(doc, "\n\n")

	for i := range blockStrings {
		if len(blockStrings) == 0 {
			log.Println("block length 0")
			continue
		}

		mode := 0 // 0: default, 1: append, 2: stop
		switch blockStrings[i][0] {
		case '[':
			mode = 0
		case '+':
			mode = 1
			blockStrings[i] = blockStrings[i][1:]
		case '-':
			mode = 2
			blockStrings[i] = blockStrings[i][1:]
		default:
			runeValue, _ := utf8.DecodeRuneInString(blockStrings[i]) // 주어진 문자열의 처음부터의 유니코드 문자 하나를 반환한다.
			log.Fatalf("Unknown letter (Expected '[', '+', '-', Unexpected %c)\n(line %s)\n", runeValue, blockStrings[i])
		}

		blockLines := strings.Split(blockStrings[i], "\n")
		if mode != 2 && len(blockLines) != len(lang)+1 { // stop이 아니면서 줄 개수가 맞지 않을 때
			log.Fatalf("Not enough lines (Expected %d, Unexpected %d)\n", len(lang)+1, len(blockLines))
		}

		meta := blockLines[0]        // timing, code
		meta = meta[1 : len(meta)-1] // remove '[' ... ']'
		metas := strings.Split(meta, "][")

		code := "stop"
		var contents map[string]string
		if mode != 2 {
			contents = make(map[string]string)
			code = metas[1]
			contentLang := blockLines[1:]
			for i := range lang {
				if i != 0 && contentLang[i] == "=" {
					contents[lang[i]] = contentLang[0]
				} else {
					contents[lang[i]] = contentLang[i]
				}
			}
		}

		block := Block{
			Timing:   metas[0],
			ms:       parseTiming(metas[0]),
			Code:     code,
			Contents: contents,
		}

		if mode == 1 {
			block.Append = true
		} else if mode == 2 {
			block.Stop = true
		}

		blocks = append(blocks, block)
	}

	return
}
