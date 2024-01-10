// https://pkg.go.dev/github.com/bwmarrin/discordgo

package main

import (
	"fmt"
	"os"
	"log"
	"syscall"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"

)

func main() {
	loadEnvs()
	
	sess, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	
	if err != nil {
		log.Fatal(err)
	}

	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Content == "Wallace" {
			s.ChannelMessageSend(m.ChannelID, "Fuck off!")
		}
	})

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	
	err = sess.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer sess.Close()

	fmt.Println("Wallace is pissing off!")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func loadEnvs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}