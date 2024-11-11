package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)

	go grpcServerStart(stopCh)
	go httpServerStart(stopCh)

	<-stopCh
	log.Println("graceful shutdown")
}
