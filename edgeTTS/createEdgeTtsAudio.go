package edgeTTS

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type SpeechMark struct {
	Time  int    `json:"time"`
	Value string `json:"value"`
}

type TextToSpeechResponse struct {
	SpeechMarks []*SpeechMark `json:"speech_marks"`
	AudioFile   string        `json:"audio_file"`
	Text        string        `json:"text"`
}

func CreateEdgeTtsAudio(ctx context.Context, text string, voice string) (*TextToSpeechResponse, error) {

	fileName := fmt.Sprintf("%s.mp3", uuid.New().String())

	ttsPayload := Args{
		Text:       text,
		Voice:      voice,
		Rate:       "1.2",
		WriteMedia: fileName,
	}
	synthesizer := NewTTS(ttsPayload)

	synthesizer.AddText(ttsPayload.Text, ttsPayload.Voice, ttsPayload.Rate, "")

	var speechMarks []*SpeechMark
	for _, word := range synthesizer.Speak() {
		speechMarks = append(speechMarks, &SpeechMark{
			Time:  word.Time,
			Value: word.Value,
		})
	}

	return &TextToSpeechResponse{
		SpeechMarks: speechMarks,
		AudioFile:   fileName,
		Text:        text,
	}, nil

}
