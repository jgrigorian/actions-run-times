package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/miquella/ask"
	"io"
	"net/http"
	"os"
	"time"
)

type Run struct {
	TotalCount   int `json:"total_count"`
	WorkflowRuns []struct {
		ID           int64     `json:"id"`
		Name         string    `json:"name"`
		Path         string    `json:"path"`
		DisplayTitle string    `json:"display_title"`
		RunNumber    int       `json:"run_number"`
		Event        string    `json:"event"`
		Status       string    `json:"status"`
		Conclusion   string    `json:"conclusion"`
		WorkflowID   int       `json:"workflow_id"`
		URL          string    `json:"url"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
		RunStartedAt time.Time `json:"run_started_at"`
	} `json:"workflow_runs"`
}

// Colors
var Red = color.New(color.FgRed).PrintfFunc()
var Yellow = color.New(color.FgYellow).PrintfFunc()
var Green = color.New(color.FgGreen).PrintfFunc()
var Blue = color.New(color.FgCyan).PrintfFunc()

func main() {

	// Flags
	orgPtr := flag.String("org", "", "The organization that owns the repository")
	repoPtr := flag.String("repo", "", "The desired repository")
	workflowIdPtr := flag.String("workflowId", "", "The organization that owns the repository")
	branchPtr := flag.String("branch", "", "the desired branch")
	flag.Parse()

	url := fmt.Sprintf("https://api.github.com/repos/%v/%v/actions/workflows/%v/runs?branch=%v&status=success", *orgPtr, *repoPtr, *workflowIdPtr, *branchPtr)

	token := getGHToken()
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))

	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	jsonByte, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	var result Run
	jsonErr := json.Unmarshal(jsonByte, &result)
	if jsonErr != nil {
		panic(jsonErr)
	}

	var durationsSlice []time.Duration
	for _, item := range result.WorkflowRuns {
		//fmt.Println(" ")
		//fmt.Printf("ID: %v\n", item.ID)
		//fmt.Printf("Name: %v\n", item.Name)
		//
		//fmt.Printf("Started: %v\n", item.RunStartedAt.Second())
		//fmt.Printf("Ended: %v\n", item.UpdatedAt.Second())
		//fmt.Printf("Duration: %v\n", item.UpdatedAt.Sub(item.CreatedAt))

		durationsSlice = append(durationsSlice, item.UpdatedAt.Sub(item.CreatedAt))
	}

	var totalTime time.Duration
	for _, item := range durationsSlice {
		totalTime += item
	}

	runCount := int(result.TotalCount)
	avgTime := totalTime / time.Duration(runCount)

	Blue("Repository\t\t\tBranch\t\t\t\tSuccessful Runs\t\tAverage Build Time\n")
	//fmt.Println(strings.Repeat("-", 125))
	fmt.Printf("%v/%v\t\t\t%v\t\t%v\t\t\t%v\n", *orgPtr, *repoPtr, *branchPtr, result.TotalCount, avgTime)

	//Blue("Branch: ")
	//fmt.Printf("%v\n", *branchPtr)

	//Blue("Successful Runs: ")
	//fmt.Printf("%v\n", result.TotalCount)

	//Blue("Average Build Time: ")
	//fmt.Printf("%v\n", avgTime)
}

func getGHToken() string {
	//	Check if GH_TOKEN env var exists, if not, prompt for token
	//	Hint: You can store your token in your bash or zsh profile:
	//	export GH_TOKEN="<enter token here>"
	if _, ok := os.LookupEnv("GH_TOKEN"); ok {
		return os.Getenv("GH_TOKEN")
	} else {
		Yellow("WARNING: Could not find GH_TOKEN environment variable.\n")
		token, _ := ask.HiddenAsk("Please enter your Github token...\n")
		return token
	}
}
