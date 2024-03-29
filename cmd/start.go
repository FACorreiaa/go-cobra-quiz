package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type question struct {
	text          string
	options       []string
	correctOption int
}

type response struct {
	question question
	answer   int
}

var startQuizCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the quiz",
	Run: func(cmd *cobra.Command, args []string) {
		startQuiz()
	},
}

var setNameCmd = &cobra.Command{
	Use:   "setname [name]",
	Short: "Sets the user's name",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		fmt.Printf("Hello, %s! Let's start the quiz.\n", name)

		fmt.Printf("Name set to %s\n", name)
	},
}

func loadQuestions() []question {
	return []question{
		{
			text:          "What is the capital of France?",
			options:       []string{"London", "Paris", "Berlin"},
			correctOption: 2,
		},
	}
}

func startQuiz() {
	questions := loadQuestions()

	responses := make(chan response)

	go func() {
		for _, q := range questions {
			fmt.Println(q.text)
			for i, opt := range q.options {
				fmt.Printf("%d. %s\n", i+1, opt)
			}
			var ans int
			fmt.Print("Your answer: ")
			fmt.Scanln(&ans)
			responses <- response{question: q, answer: ans}
		}
		close(responses) // Close the channel after all questions are answered
	}()

	var score int
	for r := range responses {
		if r.question.correctOption == r.answer {
			score++
		}
	}

	fmt.Printf("Your score: %d out of %d\n", score, len(questions))
}
