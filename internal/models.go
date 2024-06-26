package internal

import "github.com/google/uuid"

type Session struct {
	ID   uuid.UUID `json:"id"`
	User []User
}

type User struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Answers []Answer  `json:"answers"`
	Score   int       `json:"score"`
}

type Question struct {
	ID      int      `json:"id"`
	Text    string   `json:"text"`
	Answers []string `json:"answers"`
	Correct int      `json:"correct"`
}

type Answer struct {
	QuestionID int    `json:"question_id"`
	Answer     string `json:"answer"`
}

type MultipleChoiceQuestion struct {
	ID          int      `json:"id"`
	Question    string   `json:"text"`
	Options     []string `json:"options"`
	CorrectAns  string   `json:"correct_answer"`
	UserAnswer  string   `json:"user_answer,omitempty"`
	Explanation string   `json:"explanation,omitempty"`
}

// change this

type CreateUser struct {
	UserID    uuid.UUID `json:"user_id"`
	Username  string    `json:"username"`
	SessionID uuid.UUID `json:"session_id"`
}

type UpdateUser struct {
	UserID   uuid.UUID `json:"user_id"`
	Username string    `json:"username"`
}

type QuestionsReply struct {
	Score          int     `json:"score"`
	CorrectAnswers int     `json:"correct_answers"`
	Percentile     float64 `json:"percentile"`
	Message        string  `json:"message"`
}

type Ranking struct {
	UserID   uuid.UUID `json:"user_id"`
	Username string    `json:"username"`
	Score    int       `json:"score"`
}

var MultipleChoiceQuestions = []MultipleChoiceQuestion{
	{
		ID:          1,
		Question:    "What command is used to run a Go program? ",
		Options:     []string{"A. go run", "B. go start", "C. go exec", "D. go boot"},
		CorrectAns:  "A. go run",
		Explanation: "go run main.go",
	},
	{
		ID:          2,
		Question:    "What is the largest planet in the solar system?",
		Options:     []string{"A. Earth", "B. Mars", "C. Jupiter", "D. Saturn"},
		CorrectAns:  "D. Jupiter",
		Explanation: "Jupiter is the largest planet in the solar system.",
	},
	{
		ID:          3,
		Question:    "Which country has higher population?",
		Options:     []string{"A. India", "B. China", "C. USA", "D. Brasil"},
		CorrectAns:  "A. India",
		Explanation: "India is the most populous country in the world.",
	},
	{
		ID:          4,
		Question:    "What is the capital of Japan?",
		Options:     []string{"A. Beijing", "B. Seoul", "C. Tokyo", "D. Bangkok"},
		CorrectAns:  "C. Tokyo",
		Explanation: "Tokyo is the capital of Japan",
	},
	{
		ID:          5,
		Question:    "Which planet is known as the Red Planet?",
		Options:     []string{"A. Earth", "B. Mars", "C. Jupiter", "D. Saturn"},
		CorrectAns:  "B. Mars",
		Explanation: "Mars is the red planet."},
	{
		ID:          6,
		Question:    "What is the boiling point of water in Celsius?",
		Options:     []string{"A. 0", "B. 100", "C. 212", "D. 373"},
		CorrectAns:  "B. 100",
		Explanation: "Water boils at 100 degrees Celsius.",
	},
	{
		ID:          7,
		Question:    "What is the largest planet in the solar system?",
		Options:     []string{"A. Earth", "B. Mars", "C. Jupiter", "D. Saturn"},
		CorrectAns:  "C. Jupiter",
		Explanation: "Jupiter is the largest planet in the solar system.",
	},
	{
		ID:          8,
		Question:    "How is the mascot of Pokemon called?",
		Options:     []string{"A. Pikachu", "B. Squirtle", "C. Bulbasaur", "D. Charmander"},
		CorrectAns:  "A. Pikachu",
		Explanation: "Pikachu is the mascot of Pokemon",
	},
	{
		ID:          9,
		Question:    "What is the capital of Portugal",
		Options:     []string{"A. New York", "B. Lisbon", "C. Madrid", "D. Rome"},
		CorrectAns:  "B. Lisbon",
		Explanation: "Lisbon is the capital of Portugal",
	},
	{
		ID:          10,
		Question:    "Who plays on Madison Square Garden?",
		Options:     []string{"A. New York Knicks", "B. Miami Heat", "C. Real Madrid", "D. Inter Miami"},
		CorrectAns:  "A. New York Knicks",
		Explanation: "New York Knicks play on Madison Square Garden (MSG)",
	},
}

type QuestionList struct {
	ID       int      `json:"id"`
	Options  []string `json:"options"`
	Question string   `json:"question"`
}
