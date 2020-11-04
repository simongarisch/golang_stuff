// go run text_to_speech.go
// https://github.com/simongarisch/go-texttospeech
package main

import "github.com/asticode/go-texttospeech/texttospeech"

func main() {
	tts := texttospeech.NewTextToSpeech()
	tts.Say("The quick brown fox jumped over the lazy dog")
	tts.Say("She sells sea shells by the sea shore")
}
