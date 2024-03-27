package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/FACorreiaa/go-cobra-quiz/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := server.Run(ctx); err != nil {
			fmt.Println("Error running server:", err)
			cancel()
		}
	}()

	////wait server
	//time.Sleep(100 * time.Millisecond)
	//
	//wg.Add(1)
	//// start CLI
	//go func() {
	//	defer wg.Done()
	//	// Execute the CLI commands
	//	cmd.Execute()
	//}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	cancel()

	wg.Wait()
}
