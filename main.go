package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Hello Turai Log World")

	ctx, cancel := context.WithCancel(context.Background())

	done := make(chan bool, 1)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func(ctx context.Context) {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		cancel()
		time.Sleep(3 * time.Second)
		done <- true
	}(ctx)

	tjo := &TuraiJsonOutputter{}
	go tjo.Run(ctx)

	<-done
	fmt.Println("exiting")
}
