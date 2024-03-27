package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/FACorreiaa/go-cobra-quiz/configs"
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

var name string
var userNameChan = make(chan string)

var setNameCmd = &cobra.Command{
	Use:   "setname [name]",
	Short: "Sets the user's name",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		config, _ := configs.InitConfig()
		// name = args[0]
		// fmt.Printf("Hello, %s! Let's start the quiz.\n", name)
		// userNameChan <- name
		//name := args[0]

		// Make an HTTP request to the /hello endpoint of your server
		resp, err := http.Get("http://localhost" + config.Server.Addr)
		if err != nil {
			fmt.Printf("Error setting name: %v\n", err)
			return
		}
		defer resp.Body.Close()

		// Read and process the response
		var jsonResponse map[string]string
		if err := json.NewDecoder(resp.Body).Decode(&jsonResponse); err != nil {
			fmt.Printf("Error decoding response: %v\n", err)
			return
		}

		// Display the message from the response
		message, ok := jsonResponse["message"]
		if !ok {
			fmt.Println("Unexpected response format")
			return
		}
		fmt.Println(message)

	},
}

func loadQuestions() []question {
	return []question{
		{
			text:          "What is the capital of France?",
			options:       []string{"London", "Paris", "Berlin"},
			correctOption: 2,
		},
		// Add more questions as needed
	}
}

func startQuiz() {
	// Load questions
	questions := loadQuestions()

	// Initialize a channel to receive quiz responses
	responses := make(chan response)

	// Start the quiz
	go func() {
		for _, q := range questions {
			// Display question
			fmt.Println(q.text)
			// Display options
			for i, opt := range q.options {
				fmt.Printf("%d. %s\n", i+1, opt)
			}
			// Ask for answer
			var ans int
			fmt.Print("Your answer: ")
			fmt.Scanln(&ans)
			// Send response to channel
			responses <- response{question: q, answer: ans}
		}
		close(responses) // Close the channel after all questions are answered
	}()

	// Process responses
	var score int
	for r := range responses {
		if r.question.correctOption == r.answer {
			score++
		}
	}

	// Display score
	fmt.Printf("Your score: %d out of %d\n", score, len(questions))
}
