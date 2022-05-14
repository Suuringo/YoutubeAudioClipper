package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Suuringo/voice-clip-studio/clipper"
)

func main() {
	args := os.Args[1:]

	if len(args) != 4 {
		printHelp()
		return
	}

	var id string = args[0]
	startS, err := strconv.ParseFloat(args[1], 32)
	if err != nil {
		log.Fatalln("error parsing start second")
	}
	endS, err := strconv.ParseFloat(args[2], 32)
	if err != nil {
		log.Fatalln("error parsing start second")
	}
	var name string = args[3]

	clipper := clipper.Clipper{
		YoutubeId:   id,
		StartSecond: float32(startS),
		EndSecond:   float32(endS),
		OutputName:  name,
	}

	if err := clipper.Run(); err != nil {
		log.Fatalf("error clipping: %v", err)
	}
}

func printHelp() {
	fmt.Fprintf(os.Stderr, "Usage: %s <youtubeVideoID> <startSecond> <endSecond> <outputName>\n", os.Args[0])
}
