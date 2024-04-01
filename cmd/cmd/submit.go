/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type QuizResults struct {
	Score          int     `json:"score"`
	CorrectAnswers int     `json:"correct_answers"`
	Percentile     float32 `json:"percentile"`
	Message        string  `json:"message"`
}

// submitCmd represents the submit command
var submitCmd = &cobra.Command{
	Use:   "submit",
	Short: "Submit answers",
	Args:  cobra.ExactArgs(10),

	Run: func(cmd *cobra.Command, args []string) {
		Submit(args)
	},
}

func Submit(args []string) {
	userID := viper.GetString("user_id")
	if userID == "" {
		log.Fatalln("User ID not found. Please run 'start' command first.")
	}

	quizAnswers := make(map[string]string)
	for _, arg := range args {
		parts := strings.Split(arg, "=")
		if len(parts) != 2 {
			log.Fatalf("Invalid argument: %s\n", arg)
		}
		quizAnswers[parts[0]] = parts[1]
	}
	payload, err := json.Marshal(quizAnswers)
	if err != nil {
		log.Fatalf("Failed to marshal quiz answers: %v\n", err)
	}

	// Send POST request
	url := fmt.Sprintf("http://localhost:8080/session/%s/submit-quiz", userID)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Fatalf("Failed to submit quiz: %v\n", err)
	}
	defer resp.Body.Close()

	// Decode response
	var q QuizResults
	err = json.NewDecoder(resp.Body).Decode(&q)
	if err != nil {
		log.Fatalf("Failed to decode response: %v\n", err)
	}

	// Print result
	fmt.Println("Score:", q.Score)
	fmt.Println("Correct answers:", q.CorrectAnswers)
	fmt.Println("Percentile:", q.Percentile)
	fmt.Println("Message:", q.Message)

}

func init() {
	rootCmd.AddCommand(submitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// submitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// submitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
