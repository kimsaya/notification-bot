package main

import (
	"fmt"
	"log"
	API "notification-bot/api"
	REPO "notification-bot/repository"
	UTILS "notification-bot/utils"
	"os"
	"os/signal"
	"syscall"
)

var version = "0.0.1 ALPHA"

func main() {
	fmt.Println("Hello World")
	UTILS.Initalize(LOG_DIR, LOG_DUR)
	REPO.Initalize(DAT_DIR)
	API.Initalize(BOT_TOKEN, BOT_CHANNEL)

	log.Println("Notification Bot:", version)
	fmt.Println("Testing is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	API.Close()
}
