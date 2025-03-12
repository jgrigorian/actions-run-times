package list

import (
	"encoding/json"
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"os"
	"strconv"

	//"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"

	"io"
	"net/http"
	//"strconv"
	"time"
)

type Workflow struct {
	TotalCount int `json:"total_count"`
	Workflows  []struct {
		ID        int64     `json:"id"`
		Name      string    `json:"name"`
		Path      string    `json:"path"`
		State     string    `json:"state"`
		URL       string    `json:"url"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"workflows"`
}

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

func Workflows(c *cli.Context) {
	owner := c.String("owner")
	repo := c.String("repo")

	url := fmt.Sprintf("https://api.github.com/repos/%v/%v/actions/workflows", owner, repo)

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	jsonByte, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	var result Workflow
	jsonErr := json.Unmarshal(jsonByte, &result)
	if jsonErr != nil {
		panic(jsonErr)
	}

	runCount := int(result.TotalCount)
	if runCount == 0 {
		color.Yellow("No workflow runs found")
		os.Exit(0)
	}

	// Table
	t := table.New().Border(lipgloss.HiddenBorder()).StyleFunc(func(row, col int) lipgloss.Style {
		return lipgloss.NewStyle().Padding(0, 2)
	})

	t.Headers(
		color.HiMagentaString("Repository"),
		color.HiMagentaString("Workflow"),
		color.HiMagentaString("ID"),
		color.HiMagentaString("Successful Runs"),
		color.HiMagentaString("Average Run Times"))

	for i, r := range result.Workflows {

		wfRuns, totalRunCount, avgTime := workflowRuns(owner, repo, result.Workflows[i].ID)
		t.Row(fmt.Sprintf("%v/%v", owner, repo), wfRuns.WorkflowRuns[i].Name, strconv.FormatInt(r.ID, 10), totalRunCount, avgTime.String())
	}

	fmt.Println(t.Render())
}

func workflowRuns(owner string, repo string, workflowId int64) (*Run, string, time.Duration) {

	url := fmt.Sprintf("https://api.github.com/repos/%v/%v/actions/workflows/%v/runs", owner, repo, workflowId)

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

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
		durationsSlice = append(durationsSlice, item.UpdatedAt.Sub(item.CreatedAt))
	}

	var totalTime time.Duration
	for _, item := range durationsSlice {
		totalTime += item
	}

	runCount := int(result.TotalCount)
	var avgTime time.Duration

	if runCount == 0 {
		fmt.Println("No workflow runs found")
		avgTime = 0
	} else {
		avgTime = totalTime / time.Duration(runCount)
	}

	return &result, strconv.Itoa(result.TotalCount), time.Duration(avgTime).Round(time.Second)
}
