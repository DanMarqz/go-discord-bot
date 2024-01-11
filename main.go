// https://pkg.go.dev/github.com/bwmarrin/discordgo

package main

import (
	"fmt"
	"os"
	"log"
	"syscall"
	"os/signal"
	"time"
	"crypto/tls"
	"net/http"
  "encoding/json"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/gocolly/colly/v2"
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

		embed := &discordgo.MessageEmbed {
			Author:      &discordgo.MessageEmbedAuthor{},
			Color:       0x00ff00, // Green
			Description: "This is me",
			Fields: []*discordgo.MessageEmbedField {
				&discordgo.MessageEmbedField {
					Name:   "I'm feelin cute",
					Value:  "Might be delete later",
					Inline: true,
				},
			},
			Image: &discordgo.MessageEmbedImage {
				URL: "https://res.cloudinary.com/practicaldev/image/fetch/s--r0-zDHWy--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/yhzx8nb3vlj172c3aq7z.png",
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail {
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
			case "Wallace, dame el precio de la tasa bcv, por favor":
				s.ChannelMessageSend(m.ChannelID, "Claro, toma: "+getDataBcv() )
			case "Wallace, dime tus pendientes":
				s.ChannelMessageSend(m.ChannelID, "Meterme chatgpt pa poder hablar,\nRefactorizar en POO,\nGracias.")
		}

	})

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	
	err = sess.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer sess.Close()

	// fmt.Println(getDataBcv())
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

type BCVDatata struct {
	USD  string `json:"usd_price"`
	EUR string `json:"eur_price"`
	VDate  string `json:"value_date"`
}

func getDataBcv() (string){
	// Crear un cliente HTTP personalizado con verificación TLS deshabilitada
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// Crear un nuevo colector con el cliente HTTP personalizado
	c := colly.NewCollector()
	c.WithTransport(httpClient.Transport)

	var priceUSD, priceEUR, fechaValor string

	// Definir la lógica para manejar los elementos extraídos
	c.OnHTML("#dolar", func(e *colly.HTMLElement) {
		priceUSD = e.DOM.Find("strong").Text()
	})

	c.OnHTML("#euro", func(e *colly.HTMLElement) {
		priceEUR = e.DOM.Find("strong").Text()
	})

	c.OnHTML(".dinpro", func(e *colly.HTMLElement) {
		fechaValor = e.DOM.Find("span").Text()
	})

	// Visitar la URL deseada
	err := c.Visit("https://bcv.org.ve")
	if err != nil {
		log.Fatal(err)
	}

  // Crear una instancia de la estructura BCVDatata con datos
	data := BCVDatata{
		USD:    priceUSD,
		EUR:    priceEUR,
		VDate:  fechaValor,
	}

	// Convertir la estructura a JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return "Ha ocurrido un error al obtener la data desde el bcv ..."
	}

	// Imprimir el JSON resultante
	// fmt.Println(string(jsonData)) 
	return string(jsonData)
}
