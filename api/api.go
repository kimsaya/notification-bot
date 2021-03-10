package api

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

var discord *discordgo.Session
var botChannel []string

// Initalize begin api
func Initalize(inToken string, inBotChannel []string) bool {
	botChannel = inBotChannel
	discord, err := discordgo.New("Bot " + inToken)
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
