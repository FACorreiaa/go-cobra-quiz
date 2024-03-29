package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/FACorreiaa/go-cobra-quiz/api"
	"github.com/FACorreiaa/go-cobra-quiz/configs"
)

func Run(ctx context.Context) error {
	config, err := configs.InitConfig()
	if err != nil {
		fmt.Printf("Error initializing config: %s", err)
		panic(err)
	}
	repo := api.NewRepositoryStore()
	service := api.NewServiceStore(repo)

	srv := &http.Server{
		Addr:              config.Server.Addr + ":" + config.Server.Port,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      config.Server.WriteTimeout,
		ReadTimeout:       config.Server.ReadTimeout,
		IdleTimeout:       config.Server.IdleTimeout,
		Handler:           Logger(Router(service)),
	}

	go func() {
		fmt.Printf("Starting server %s on port %s\n", config.Server.Addr, config.Server.Port)
		if err := srv.ListenAndServe(); err != nil {
			fmt.Println("ListenAndServe error:", err)
		}
	}()

	err = InitPprof(config.Pprof.Addr, config.Pprof.Port)
	if err != nil {
		fmt.Printf("Error initializing pprof config: %s", err)
		panic(err)
	}

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
