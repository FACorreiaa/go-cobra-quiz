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
		Addr:         config.Server.Addr + ":" + config.Server.Port,
		WriteTimeout: config.Server.WriteTimeout,
		ReadTimeout:  config.Server.ReadTimeout,
		IdleTimeout:  config.Server.IdleTimeout,
		Handler:      Logger(Router()), // Router(),
	}

	go func() {
		fmt.Printf("Starting server %s on port %s\n", config.Server.Addr, config.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("ListenAndServe error:", err)
		}
	}()

	// select {
	// case userName = <-userNameChan:
	// 	// Continue starting the server
	// case <-ctx.Done():
	// 	return nil
	// }

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
