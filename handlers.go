package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var hiList = []string{"Hi", "Hello", "Hey", "hi", "hello", "hey", "sup", "Sup"}

func initHandler(discord *discordgo.Session, ready *discordgo.Ready) {
	err := discord.UpdateStatus(0, "Kicking some A$$")
	if err != nil {
		fmt.Println("Error attempting to set my status")
	}
	servers := discord.State.Guilds
	fmt.Printf("TheDonBot has started on %d servers", len(servers))
}

func reactionHandler(discord *discordgo.Session, react *discordgo.MessageReactionAdd) {
	user := react.MessageReaction.UserID
	if user == botID {
		//Do nothing because the bot is treacting
		return
	}

	//do something

	fmt.Printf("Reaction: %+v || From: %s\n", react.MessageReaction.Emoji.Name, react.MessageReaction.UserID)
}

func hiHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	user := message.Author
	if user.ID == botID || user.Bot {
		//Do nothing because the bot is talking
		return
	}

	content := message.Content
	if contains(hiList, content) {
		discord.ChannelMessageSend(message.ChannelID, "It looks like you just said a greeting and nothing else. Please contribute more to the conversation rather than filling the channel with greetings. \nhttp://www.nohello.com/")
	}

	fmt.Printf("Message: %+v || From: %s\n", content, user)
}

func typingHandler(discord *discordgo.Session, typingEvent *discordgo.TypingStart) {
	fmt.Printf("%s is typing\n", typingEvent.UserID)
	// discord.ChannelMessageSend(typingEvent.ChannelID, "Someone is typing... and it's not YA BOI GUZMA!")
	//do something
}
