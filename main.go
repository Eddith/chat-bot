package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/frankzhao/openai-go"
	"time"
)

func main() {
	// OpenAI API anahtarınızı ve model kimliğinizi burada ayarlayın
	apiKey := "OPEN_AI_API_KEY"
	modelID := "OPEN_AI_MODEL_ID"

	// Discord botu oluşturun ve API anahtarınızı kullanarak OpenAI dil modeliyle bağlantı kurun
	bot, err := discordgo.New("Bot " + "DISCORD_BOT_TOKEN")
	if err != nil {
		panic(err)
	}
	client := openai.New(apiKey)
	if err != nil {
		panic(err)
	}

	// Botun yanıt vereceği mesajları belirleyin
	bot.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == bot.State.User.ID {
			return
		}

		// Gelen mesajı OpenAI API'sine gönderin ve yanıtı alın
		response, err := client.CompleteText(m.Content, modelID, 0, 1000)

		if err != nil {
			panic(err)
		}

		// Botun yanıtını kanala gönderin
		s.ChannelMessageSend(m.ChannelID, response.Choices[0].Text)
		s.ChannelMessage(m.ChannelID, response.Choices[0].Text)
	})

	// Botu çalıştırın
	if err != nil {
		panic(err)
	}

	stop := make(chan struct{})
	defer close(stop)
	go bot.Open()
	for {
		time.Sleep(time.Second)
	}
}
