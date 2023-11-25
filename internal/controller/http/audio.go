package http

import (
	"audio-storage/internal/service"
	"bytes"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

type audioController struct {
	AS service.AudioStorage
}

func newAudioController(AS service.AudioStorage) *audioController {
	return &audioController{AS: AS}
}

func NewAudioRoutes(r *mux.Router, as service.AudioStorage) {
	ctrl := newAudioController(as)
	r.HandleFunc("/audio/{id:[0-9]+}", ctrl.audio)
}

func (a *audioController) audio(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	audioBytes, err := a.AS.GetAudioByID(vars["id"])
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "audio/mpeg")
	audioChunks := bytes.NewBuffer(audioBytes)

	if _, err := audioChunks.WriteTo(w); err != nil {
		log.Println(errors.Wrap(err, "AS/Audio: fail to write stream"))
	}
}
