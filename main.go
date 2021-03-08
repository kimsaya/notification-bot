package main

import (
	"fmt"
	"log"
	UTILS "notification-bot/utils"
)

func main() {
	fmt.Println("Hello World")
	UTILS.InitalizeLog(LOG_DIR)
	log.Println("")
}
