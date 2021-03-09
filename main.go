package main

import (
	"fmt"
	API "notification-bot/api"
	UTILS "notification-bot/utils"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("Hello World")
	UTILS.Initalize(LOG_DIR, LOG_DUR)
	API.Initalize(BOT_TOKEN)

	fmt.Println("Testing is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	API.Close()
}
