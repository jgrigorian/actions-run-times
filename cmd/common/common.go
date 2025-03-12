package common

import (
	"github.com/fatih/color"
	"github.com/miquella/ask"
	"os"
)

func GetGHToken() string {
	//	Check if GH_TOKEN env var exists, if not, prompt for token
	//	Hint: You can store your token in your bash or zsh profile:
	//	export GH_TOKEN="<enter token here>"
	if _, ok := os.LookupEnv("GH_TOKEN"); ok {
		return os.Getenv("GH_TOKEN")
	} else {
		color.Yellow("WARNING: Could not find GH_TOKEN environment variable.\n")
		token, _ := ask.HiddenAsk("Please enter your Github token...\n")
		return token
	}
}
