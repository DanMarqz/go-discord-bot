// https://pkg.go.dev/github.com/bwmarrin/discordgo

package main

import (
	"fmt"
	"os"
	"log"
	"syscall"
	"os/signal"
	"time"

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

		// if m.Content == "Wallace" {
		// 	s.ChannelMessageSend(m.ChannelID, "Fuck off!")
		// }

		embed := &discordgo.MessageEmbed{
    Author:      &discordgo.MessageEmbedAuthor{},
    Color:       0x00ff00, // Green
    Description: "This is me",
    Fields: []*discordgo.MessageEmbedField{
        &discordgo.MessageEmbedField{
            Name:   "I'm feelin cute",
            Value:  "Might be delete later",
            Inline: true,
        },
    },
    Image: &discordgo.MessageEmbedImage{
        URL: "https://res.cloudinary.com/practicaldev/image/fetch/s--r0-zDHWy--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/yhzx8nb3vlj172c3aq7z.png",
    },
    Thumbnail: &discordgo.MessageEmbedThumbnail{
        URL: "https://res.cloudinary.com/practicaldev/image/fetch/s--r0-zDHWy--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/yhzx8nb3vlj172c3aq7z.png",
    },
    Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
    Title:     "This is me :)",
}

		switch m.Content {
			case "Wallace":
				s.ChannelMessageSend(m.ChannelID, "Fuck off!")
			case "whoru?":
			// https://github-wiki-see.page/m/bwmarrin/discordgo/wiki/FAQ
				s.ChannelMessageSendEmbed(m.ChannelID, embed)
			case "digalo ahi Wallace":
				s.ChannelMessageSend(m.ChannelID, "Claro q si tremendo relambe webos es ese tal 🔱 prieto gang 🤑")
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