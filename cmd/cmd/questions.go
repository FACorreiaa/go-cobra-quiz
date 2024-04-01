package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var getQuestionsCmd = &cobra.Command{
	Use:   "list",
	Short: "Get all multiple-choice questions",
	Run: func(cmd *cobra.Command, args []string) {
		getAllQuestions()
	},
}

func init() {
	rootCmd.AddCommand(getQuestionsCmd)
}

type Question struct {
	ID   int      `json:"id"`
	Text string   `json:"question"`
	Opts []string `json:"options"`
}

func getAllQuestions() {
	serverAddress := getServerAddress()

	resp, err := http.Get(serverAddress + "/quiz/list")
	if err != nil {
		fmt.Println("Failed to fetch questions:", err)
		return
	}
	defer resp.Body.Close()

	var questions []Question
	err = json.NewDecoder(resp.Body).Decode(&questions)
	if err != nil {
		fmt.Println("Failed to decode JSON response:", err)
		return
	}

	// Print the questions
	for _, q := range questions {
		fmt.Printf("%v. Question: %v\n", q.ID, q.Text)
		fmt.Println("Options:")
		options := q.Opts
		for _, option := range options {
			fmt.Println("-", option)
		}
		fmt.Println()
	}
}
