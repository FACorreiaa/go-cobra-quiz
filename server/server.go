package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/FACorreiaa/go-cobra-quiz/configs"
)

func Run(ctx context.Context) error {
	config, err := configs.InitConfig()
	if err != nil {
		fmt.Printf("Error initializing config: %s", err)
		panic(err)
	}

	srv := &http.Server{
		Addr:         config.Server.Addr,
		WriteTimeout: config.Server.WriteTimeout,
		ReadTimeout:  config.Server.ReadTimeout,
		IdleTimeout:  config.Server.IdleTimeout,
		Handler:      Router(),
	}

	go func() {
		fmt.Println("Starting server " + config.Server.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("ListenAndServe error:", err)
		}
	}()

	<-ctx.Done()

	// Shutdown server
	ctx, cancel := context.WithTimeout(context.Background(), config.Server.GracefulTimeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Server shutdown error:", err)
	}
	fmt.Println("Shutting down")
	return nil
}
