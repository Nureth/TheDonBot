package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

func sendFile(discord *discordgo.Session, id string, filename string, contentType string) {
	r, err := os.Open(filename)
	if err != nil {
		fmt.Println("error in opening file", err)
		return
	}

	message := &discordgo.MessageSend{}
	file := &discordgo.File{}
	file.Name = filename
	file.ContentType = contentType
	file.Reader = r
	message.Files = []*discordgo.File{file}

	sent, err := discord.ChannelMessageSendComplex(id, message)
	fmt.Println(sent)
	return
}

func getFileFromName(name string) (path string) {
	// TODO
	return ""
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
