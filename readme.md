# edge-tts-go

`edge-tts-go` is a Go module that allows you to use Microsoft Edge's online text-to-speech service from within your Go code.

## Installation

To install it, run the following command:

```sh
go get github.com/neoblaze/edge-tts-go
```

## Usage

### Basic Usage

Here is a basic example of how to convert text to speech:

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/neoblaze/edge-tts-go/edgeTTS"
)

func main() {
	text := "The quick brown fox jumps over the lazy dog near the shimmering lake at dawn."
	voice := "en-IN-NeerjaNeural"

	response, err := edgeTTS.CreateEdgeTtsAudio(context.Background(), text, voice)
	if err != nil {
		log.Fatalf("Error creating edge tts audio: %v", err)
	}

	fmt.Printf("Audio saved to: %s\n", response.AudioFile)
	fmt.Printf("Original text: %s\n", response.Text)

	// The response also contains speech marks for word timings
	// for _, mark := range response.SpeechMarks {
	// 	fmt.Printf("Word: '%s' at %dms\n", mark.Value, mark.Time)
	// }
}
```
This will create an mp3 file with a random UUID as its name (e.g., `xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx.mp3`) in the current directory.

### Listing Available Voices

You can get a list of all available voices that can be used.

```go
package main

import (
	"fmt"
	"log"

	"github.com/neoblaze/edge-tts-go/edgeTTS"
)

func main() {
	voices, err := edgeTTS.ListVoices()
	if err != nil {
		log.Fatalf("Failed to get voices: %v", err)
	}

	for _, voice := range voices {
		if voice.Locale == "en-US" {
			fmt.Printf("Name: %s, ShortName: %s, Gender: %s\n",
				voice.Name, voice.ShortName, voice.Gender)
		}
	}
}
```

This will output something like:
```
Name: Microsoft Server Speech Text to Speech Voice (en-US, AnaNeural), ShortName: en-US-AnaNeural, Gender: Female
Name: Microsoft Server Speech Text to Speech Voice (en-US, AriaNeural), ShortName: en-US-AriaNeural, Gender: Female
Name: Microsoft Server Speech Text to Speech Voice (en-US, ChristopherNeural), ShortName: en-US-ChristopherNeural, Gender: Male
...
```

You can use the `ShortName` (e.g., `en-US-AriaNeural`) for the `voice` parameter in `CreateEdgeTtsAudio`.

## Acknowledgements

This library is based on the work from:
* https://github.com/rany2/edge-tts
