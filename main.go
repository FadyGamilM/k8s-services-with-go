package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// define build for build flag later
var build = "develop"

func main() {

	log.Printf("go service is starting .. [%v]\n", build)

	defer log.Println("go service is ended !")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// wait for it
	<-quit

	log.Printf("go service is stopping .. [%v]\n", build)
}
