package main

import (
	"log"
	"strconv"
)

func parseTiming(timing string) int {
	switch len(timing) {
	case 8: // 00:00.00
		timing = "00:" + timing // add hour
	case 11: // 00:00:00.00
	default:
		log.Fatalln("Unexpected timing format...", timing)
	}
	hours, err := strconv.Atoi(timing[0:2])
	if err != nil {
		log.Fatalln(err)
	}
	minutes, err := strconv.Atoi(timing[3:5])
	if err != nil {
		log.Fatalln(err)
	}

	if len(timing[6:]) != 5 {
		log.Fatalln("Unexpected second format...", timing[6:])
	}

	sec, err := strconv.Atoi(timing[6:8])
	if err != nil {
		log.Fatalln(err)
	}
	ms10, err := strconv.Atoi(timing[9:])
	if err != nil {
		log.Fatalln(err)
	}

	return hours*3600000 + minutes*60000 + sec*1000 + ms10*10
}
