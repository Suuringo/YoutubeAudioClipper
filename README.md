# Youtube Audio Clipper

Youtube Audio Clipper is a small utility written in Golang to clip audio stream from a Youtube video.
`youtube-dl` and `ffmpeg` need to be in `$PATH` for Youtube Audio Clipper to work.


# Usage

In `cmd/clip-audio` :
`go run . <youtubeVideoID> <startSecond> <endSecond> <outputName>`
