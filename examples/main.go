package main

import (
	"context"
	"fmt"
	"log"

	"github.com/neoblaze/edge-tts-go/edgeTTS"
)

func main() {

	// voices, err := edgeTTS.ListVoices()
	// if err != nil {
	// 	log.Fatalf("Error listing voices: %v", err)
	// }
	// for _, voice := range voices {
	// 	if voice.Locale == "en-US" {
	// 		fmt.Printf("Name: %s, ShortName: %s, Gender: %s\n",
	// 			voice.Name, voice.ShortName, voice.Gender)
	// 	}
	// }

	text := "The quick brown fox jumps over the lazy dog near the shimmering lake at dawn."
	voice := "en-IN-NeerjaNeural"

	response, err := edgeTTS.CreateEdgeTtsAudio(context.Background(), text, voice)
	if err != nil {
		log.Fatalf("Error creating edge tts audio: %v", err)
	}

	// speechMarks := response.SpeechMarks

	fmt.Println(response.AudioFile)
	fmt.Println(response.Text)
	// for _, mark := range speechMarks {
	// 	fmt.Println("Time: ", mark.Time)
	// 	fmt.Println("Value: ", mark.Value)
	// }

}
