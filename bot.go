package main

import (
	"fmt"
	"log"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"gitlab.com/back1ng1/question-bot/internal/database"
	"gitlab.com/back1ng1/question-bot/internal/database/models"
	"gitlab.com/back1ng1/question-bot/internal/interval"
	"gitlab.com/back1ng1/question-bot/internal/routes/api"
)

func runApi() {
	app := fiber.New()

	app.Use(cors.New())

	// register routes
	api.QuestionRoutes(app)
	api.PresetRoutes(app)
	api.AnswerRoutes(app)
	api.UserRoutes(app)

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

	database.Database.DB.AutoMigrate(
		&models.Question{},
		&models.Answer{},
		&models.Preset{},
		&models.User{},
	)

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	ticker := time.NewTicker(30 * time.Second)
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				intervals := interval.GetActual(30)

				fmt.Println(intervals)
				for _, interval := range intervals {

					users := []models.User{}
					database.Database.DB.Find(&users, models.User{Interval: interval, IntervalEnabled: true})

					for _, user := range users {
						question := user.GetQuestion()
						poll := question.CreatePoll(user.ChatId)
						bot.Send(poll)
					}
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	for update := range updates {
		if update.Message != nil {
			user := models.User{
				ChatId:   update.Message.Chat.ID,
				Nickname: update.Message.Chat.UserName,
			}

			var presets []models.Preset
			database.Database.DB.Find(&presets)

			if len(presets) == 0 {
				preset := models.Preset{Title: "temporary name"}
				database.Database.DB.Create(&preset)
				user.PresetId = preset.Id
			} else {
				user.PresetId = presets[0].Id
			}

			database.Database.DB.Find(&user, models.User{ChatId: user.ChatId})
			if user.ID == 0 {
				database.Database.DB.Create(&user)
			}
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
