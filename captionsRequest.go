package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetCaptions(videoId, apiKey string) string {
	captionsApiURL := fmt.Sprintf("https://youtube.googleapis.com/youtube/v3/captions?part=snippet&videoId=%s&key=%s", videoId, apiKey)
	resp, err := http.Get(captionsApiURL)
	if err != nil {
		log.Fatalln(err)
	}

	parsedBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	stringifiedBody := string(parsedBody)
	return stringifiedBody
}
