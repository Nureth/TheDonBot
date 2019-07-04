package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

var (
	commandPrefix string
	botID         string
)

func main() {
	bot := "Bot " + os.Getenv("DON_BOT_KEY")
	discord, err := discordgo.New(bot)
	errCheck("error creating discord session", err)
	user, err := discord.User("@me")
	errCheck("error retrieving account", err)

	botID = user.ID
	discord.AddHandler(reactionHandler)
	discord.AddHandler(initHandler)
	discord.AddHandler(hiHandler)
	discord.AddHandler(typingHandler)

	err = discord.Open()
	errCheck("Error opening connection to Discord", err)
	defer discord.Close()

	commandPrefix = "!"

	<-make(chan struct{})

}

func errCheck(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %+v", msg, err)
		panic(err)
	}
}
