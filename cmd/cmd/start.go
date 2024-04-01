/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start session",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		startSession()
	},
}

func startSession() {
	serverAddress := getServerAddress()
	start := time.Now()

	fmt.Println("Session started at:", start)

	resp, err := http.Post(serverAddress+"/session", "application/json", nil)
	if err != nil {
		log.Fatalln("Failed to start session:", err)
	}
	defer resp.Body.Close()

	var sessionData struct {
		UserID    string `json:"user_id"`
		Username  string `json:"username"`
		SessionID string `json:"session_id"`
	}
	err = json.NewDecoder(resp.Body).Decode(&sessionData)
	if err != nil {
		log.Fatalln("Failed to decode session data:", err)
	}

	viper.Set("user_id", sessionData.UserID)
	if err := viper.WriteConfig(); err != nil {
		fmt.Println("Error saving token:", err)
		os.Exit(1)
	}
	fmt.Println("Token saved successfully.")

}

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/app/configs")
	viper.AddConfigPath("app")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", err)
		os.Exit(1)
	}
	rootCmd.AddCommand(startCmd)
}
