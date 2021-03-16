package api

import (
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
	// discord.AddHandler(messageReaction)
	err = discord.Open()
	if err != nil {
		log.Println("Error opening Discord session: ", err)
		return false
	}
	// SERVICE.Greating(discord, inBotChannel[0])
	return true
}

// Close the Discord Session
func Close() {
	discord.Close()
}
