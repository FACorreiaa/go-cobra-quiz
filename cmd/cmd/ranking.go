/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

type Ranking struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Score    int    `json:"score"`
}

// rankingCmd represents the ranking command
var rankingCmd = &cobra.Command{
	Use:   "ranking",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		getRanking()
	},
}

func getRanking() {
	resp, err := http.Get("http://localhost:8080/session/ranking")
	if err != nil {
		fmt.Println("Failed to fetch questions:", err)
		return
	}
	defer resp.Body.Close()

	var ranking []Ranking
	err = json.NewDecoder(resp.Body).Decode(&ranking)
	if err != nil {
		fmt.Println("Failed to decode JSON response:", err)
		return
	}

	// Print the questions
	for _, r := range ranking {
		fmt.Println("User", r.Username)
		fmt.Println("Score:", r.Score)

	}
}

func init() {
	rootCmd.AddCommand(rankingCmd)
}
