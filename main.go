package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gen2brain/beeep"
)

const (
	qiauLoginBaseURL = "https://net.qiau.ac.ir:8443/lanlogin"
)

var (
	qiauHotspotUsername         string
	qiauHotspotPassword         string
	qiauHotspotLoginFormAddress string

	rnd string
)

func init() {
	qiauHotspotUsername = os.Getenv("QIAU_USERNAME")
	qiauHotspotPassword = os.Getenv("QIAU_PASSWORD")
	if qiauHotspotUsername == "" || qiauHotspotPassword == "" {
		beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
		beeep.Notify("Oops!", "Please set all QIAU_* environment variables!", "")
		os.Exit(1)
	}
	rand.Seed(time.Now().UnixNano())
	rnd = fmt.Sprintf("%d", 100000000+rand.Intn(899999999))
	qiauHotspotLoginFormAddress = fmt.Sprintf("%s?rnd=%s", qiauLoginBaseURL, rnd)
}

func main() {
	response, err := http.PostForm(qiauHotspotLoginFormAddress, url.Values{
		"username": {qiauHotspotUsername},
		"password": {qiauHotspotPassword},
		"rnd":      {rnd},
		"login":    {"1"},
		"h":        {"google.com"},
		"m":        {"24e9b323507c"},
		"i":        {"c0a80fe8"},
		"u":        {"2f"},
	})
	if err != nil {
		beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
		beeep.Notify("Oops!", "Error sending post request to login form!", "")
		os.Exit(1)
	}
	defer response.Body.Close()
	beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	beeep.Notify("Done!", "Successfully logged in to QIAU hotspot!", "")
}
