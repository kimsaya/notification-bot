package main

import (
	"fmt"
	"log"
	"net/http"
	API "notification-bot/api"
	REPO "notification-bot/repository"
	UTILS "notification-bot/utils"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var version = "0.0.1 ALPHA"

func main() {
	fmt.Println("Hello World")
	UTILS.Initalize(LOG_DIR, LOG_DUR)

	for !connected() {
		time.Sleep(5 * time.Second)
	}

	REPO.Initalize(DAT_DIR)
	API.Initalize(BOT_TOKEN, BOT_CHANNEL)

	log.Println("Notification Bot:", version)
	fmt.Println("Ai is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	API.Close()
}

func connected() (ok bool) {
	_, err := http.Get("http://clients3.google.com/generate_204")
	if err != nil {
		return false
	}
	return true
}
