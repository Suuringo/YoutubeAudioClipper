package clipper

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type Clipper struct {
	YoutubeId   string  `json:"youtubeId"`
	StartSecond float32 `json:"startSecond"`
	EndSecond   float32 `json:"endSecond"`
	OutputName  string  `json:"outputName"`
}

func (c *Clipper) Run() error {
	ytCmd := exec.Command("youtube-dl", "-g", "https://www.youtube.com/watch?v="+c.YoutubeId)
	ytLinks, err := ytCmd.Output()
	if err != nil {
		return fmt.Errorf("error running yt-dl to get video stream & audio stream links: %w", err)
	}

	// [0] is video stream, [1] is audio stream
	streams := strings.Split(string(ytLinks), "\n")

	if len(streams) < 2 {
		return errors.New("expected two streams from youtube-dl -g command")
	}

	ffmpegArgs := []string{
		"-y",
		"-ss", fmt.Sprintf("%.2f", c.StartSecond),
		"-i", streams[1],
		"-t", fmt.Sprintf("%.2f", c.EndSecond-c.StartSecond),
		"-c", "copy",
		fmt.Sprintf("%s.aac", c.OutputName),
	}

	log.Printf("Running ffmpeg %s", strings.Join(ffmpegArgs, " "))

	ffCmd := exec.Command("ffmpeg", ffmpegArgs...)

	if err := ffCmd.Run(); err != nil {
		return fmt.Errorf("error running ffmpeg to download clip: %w", err)
	}

	return nil
}
