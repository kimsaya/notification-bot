package api

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

var discord *discordgo.Session

// Initalize begin api
func Initalize(token string) bool {
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Println("[ER] discordgo.New:", err)
	}
	discord.AddHandler(messageCreate)

	err = discord.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}

	return true
}

// Close the Discord Session
func Close() {
	discord.Close()
}
