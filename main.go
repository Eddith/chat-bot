package main

import (
	"context"
	"github.com/Andrew-peng/go-dalle2/dalle2"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
	"time"
)

func main() {
	// OpenAI API anahtarınızı ve model kimliğinizi burada ayarlayın
	apiKey := "OPENAI_API_KEY"
	// modelID := "MODEL_ID"

	// Discord botu oluşturun ve API anahtarınızı kullanarak OpenAI dil modeliyle bağlantı kurun
	bot, err := discordgo.New("Bot " + "DISCORD_BOT_TOKEN")
	if err != nil {
		panic(err)
	}
	// clientChat := openai.New(apiKey)
	client, err := dalle2.MakeNewClientV1(apiKey)
	if err != nil {
		log.Fatalf("Error initializing client: %s", err)
	}

	// Botun yanıt vereceği mesajları belirleyin
	bot.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == bot.State.User.ID {
			return
		}

		// Gelen mesajı OpenAI API'sine gönderin ve yanıtı alın
		// responseChat, err := clientChat.CompleteText(m.Content, modelID, 0, 1000)

		if strings.Contains(m.Content, "DISCORD_BOT_PREFIX") {
			response, err := client.Create(
				context.Background(),
				m.Content,
				dalle2.WithNumImages(1),
				dalle2.WithSize(dalle2.LARGE),
				dalle2.WithFormat(dalle2.URL),
			)

			if err != nil {
				log.Fatal(err)
			}

			var imgURL string

			for _, img := range response.Data {
				imgURL = img.Url
			}

			// Botun yanıtını kanala gönderin
			s.ChannelMessageSend(m.ChannelID, imgURL)
			s.ChannelMessage(m.ChannelID, imgURL)

			// Botun yanıtını kanala gönderin
			// s.ChannelMessageSend(m.ChannelID, responseChat.Choices[0].Text)
			// s.ChannelMessage(m.ChannelID, responseChat.Choices[0].Text)
		}

	})

	// Botu çalıştırın
	if err != nil {
		panic(err)
	}

	stop := make(chan struct{})
	defer close(stop)
	go bot.Open()
	for {
		time.Sleep(time.Minute)
	}
}
