package server

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func InitPprof(addr, port string) error {
	pprofAddr := fmt.Sprintf("%s:%s", addr, port)
	go func() {
		log.Printf("Running pprof on %s\n", pprofAddr)
		if err := http.ListenAndServe(pprofAddr, nil); err != nil {
			log.Printf("Error running pprof server: %v\n", err)
		}
	}()
	return nil
}
