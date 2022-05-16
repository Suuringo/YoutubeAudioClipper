package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path"

	"github.com/Suuringo/voice-clip-studio/clipper"
)

func GetServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", getClip)

	return mux
}

func getClip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var clipParams clipper.Clipper

	err := json.NewDecoder(r.Body).Decode(&clipParams)
	if err != nil {
		http.Error(w, fmt.Errorf("error parsing request: %w", err).Error(), http.StatusBadRequest)
		return
	}

	if clipParams.EndSecond-clipParams.StartSecond > 20 {
		http.Error(w, "clip can't be more than 20 seconds long", http.StatusBadRequest)
		return
	}

	clipParams.OutputName = randStringBytes(10)

	tmpDir, err := ioutil.TempDir("", "clip")
	if err != nil {
		http.Error(w, fmt.Errorf("error creating temporary directory: %w", err).Error(), http.StatusInternalServerError)
		return
	}
	defer os.RemoveAll(tmpDir)

	// Output in temporary directory
	clipParams.OutputName = path.Join(tmpDir, clipParams.OutputName)
	if err := clipParams.Run(); err != nil {
		http.Error(w, fmt.Errorf("error running clip binaries: %w", err).Error(), http.StatusInternalServerError)
		return
	}

	filename := path.Base(clipParams.OutputName)

	w.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.aac", filename))
	w.Header().Set("Content-Type", "audio/aac")

	http.ServeFile(w, r, fmt.Sprintf("%s.aac", clipParams.OutputName))
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
