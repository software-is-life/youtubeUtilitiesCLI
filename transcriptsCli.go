package main

import (
	"flag"
	"fmt"
	"os"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	youtubeAPIKey := os.Getenv("YOUTUBE_API_KEY")

	if youtubeAPIKey == "" {
		fmt.Printf("You must set both Youtube API's Client Secret and Key. One or both are left unset")
		os.Exit(1)
	}

	transcriptCmd := flag.NewFlagSet("transcript", flag.ExitOnError)
	transcriptVideoId := transcriptCmd.String("videoId", "", "abc123")
	transcriptWriteFileName := transcriptCmd.String("outputFileName", "", "transcript.txt")

	if len(os.Args) < 2 {
		fmt.Println("expected 'transcript', 'videoId', and 'outputFileName'")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "transcript":
		transcriptCmd.Parse(os.Args[2:])
		transcriptDetails := GetCaptions(*transcriptVideoId, youtubeAPIKey)
		outputFileName := "./transcripts/transcript.txt"
		if *transcriptWriteFileName != "" {
			outputFileName = fmt.Sprintf("./transcripts/%s", *transcriptWriteFileName)
		}
		f, err := os.Create(outputFileName)
		checkError(err)

		defer f.Close()
		transcriptBody, err := f.WriteString(transcriptDetails)
		checkError(err)

		fmt.Printf("This transcript information: %s\n\n was writen to this file path: %s", transcriptBody, transcriptWriteFileName)
	default:
		fmt.Println("something went awry.")
		os.Exit(1)
	}

}
