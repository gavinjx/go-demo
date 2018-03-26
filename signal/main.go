package main

import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

func main(){
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGALRM, syscall.SIGTERM, syscall.SIGUSR1, syscall.SIGHUP, syscall.SIGKILL)



	// block until signal comes
	sig := <-c
	fmt.Printf("Get Sig %+v", sig)
}