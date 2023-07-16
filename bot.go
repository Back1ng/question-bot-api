package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gitlab.com/back1ng1/question-bot/internal/database"
	presets "gitlab.com/back1ng1/question-bot/internal/presets/api"
	presets_models "gitlab.com/back1ng1/question-bot/internal/presets/models"
	test_questions "gitlab.com/back1ng1/question-bot/internal/questions"
	questions "gitlab.com/back1ng1/question-bot/internal/questions/api"
	questions_models "gitlab.com/back1ng1/question-bot/internal/questions/models"
)

func runApi() {
	app := fiber.New()

	// register routes
	presets.Routes(app)
	questions.Routes(app)

	app.Listen(":3000")
}

func main() {
	godotenv.Load(".env")
	database.SetupConnection()
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TGBOT_TOKEN"))

	if err != nil {
		log.Panic(err)
	}

	go runApi()

	// example Gorm
	// first_question := models.Question{Title: "Название вопроса №1"}
	// second_question := models.Question{Title: "Название вопроса №2"}

	database.Database.DB.AutoMigrate(&questions_models.Answer{}, &questions_models.Question{}, &presets_models.Preset{})

	// db.Create(&models.Question{Title: "Название вопроса №1"})
	// db.Create(&models.Question{Title: "Название вопроса №2"})
	// db.Create(&models.Preset{
	// Title:     "Название пресета",
	// Questions: []models.Question{first_question, second_question},
	// })

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {

			question := test_questions.GetRandom()

			poll := question.CreatePoll(update.Message.Chat.ID)

			bot.Send(poll)

			// log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			// msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			// msg.ReplyToMessageID = update.Message.MessageID

			// bot.Send(msg)
		}

		if update.Poll != nil {
			for id, option := range update.Poll.Options {
				if option.VoterCount > 0 {
					fmt.Printf("Selected{id: %v, value:%s}", id, option.Text)
				}
			}
		}
	}
}
