package service

import (
	"audio-storage/internal/model"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

type audioService struct {
	apiirl string
}

func (a *audioService) GetAudioByID(id string) ([]byte, error) {
	client := &http.Client{
		Timeout: 60 * time.Second,
	}

	files, err := os.ReadDir("assets")
	if err != nil {
		return nil, errors.Wrap(err, "AS/GetAudioByID: fail to find file .mp3")
	}
	for _, f := range files {
		if f.Name() == id+".mp3" {
			fmt.Println("хуй")
			dataFile, err := os.ReadFile("assets/" + id + ".mp3")
			if err != nil {
				return nil, errors.Wrap(err, "AS/GetAudioByID: fail to read mp3")
			}
			return dataFile, nil
		}
	}
	req, err := client.Get(a.apiirl + id)
	if err != nil {
		return nil, errors.Wrap(err, "AS/GetAudioByID: fail to request for API")
	}

	bodys, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, errors.Wrap(err, "AS/GetAudioByID: fail to read request body")
	}
	bell := &model.BellInfo{}

	err = json.Unmarshal(bodys, bell)
	if err != nil {
		return nil, errors.Wrap(err, "AS/GetAudioByID: fail to read request body")
	}

	req, err = client.Get(bell.ContactAudio)
	if err != nil {
		return nil, errors.Wrap(err, "AS/GetAudioByID: fail to request for contactaudio")
	}

	idi := strconv.Itoa(bell.ID)

	newFile, err := os.Create("assets/" + idi + ".mp3")
	if err != nil {
		return nil, errors.Wrap(err, "AS/GetAudioByID: fail to create file")
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, req.Body)
	if err != nil {
		return nil, errors.Wrap(err, "AS/GetAudioByID: fail to request for copy")
	}

	dataFile, err := os.ReadFile("assets/" + idi + ".mp3")
	if err != nil {
		return nil, errors.Wrap(err, "AS/GetAudioByID: fail to read mp3")
	}

	return dataFile, nil

}

func newAudioService(apiirl string) *audioService {
	return &audioService{apiirl: apiirl}
}
