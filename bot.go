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
	user_repository "gitlab.com/back1ng1/question-bot/internal/database/repository"
	"gitlab.com/back1ng1/question-bot/internal/interval"
	"gitlab.com/back1ng1/question-bot/internal/routes/api"
	"gitlab.com/back1ng1/question-bot/internal/telegram"
)

var bot *tgbotapi.BotAPI

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

func bootstrap() {
	var err error
	godotenv.Load(".env")
	database.SetupConnection()

	bot, err = tgbotapi.NewBotAPI(os.Getenv("TGBOT_TOKEN"))
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	go runApi()

	if err != nil {
		log.Panic(err)
	}

	database.Database.DB.AutoMigrate(
		&models.Question{},
		&models.Answer{},
		&models.Preset{},
		&models.User{},
	)
}

func main() {
	bootstrap()

	ticker := time.NewTicker(30 * time.Second)
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				intervals := interval.GetActual(30)

				fmt.Println(intervals)
				for _, interval := range intervals {
					for _, user := range user_repository.FindByInterval(interval) {
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

	for update := range telegram.GetUpdates(bot) {
		if update.Message != nil {
			// создаем пользователя для его авторизации
			user := models.User{
				ChatId:   update.Message.Chat.ID,
				Nickname: update.Message.Chat.UserName,
			}

			// пользователь не может быть без пресета, даём ему один
			var presets []models.Preset
			database.Database.DB.Find(&presets)

			if len(presets) == 0 {
				preset := models.Preset{Title: "temporary name"}
				database.Database.DB.Create(&preset)
				user.PresetId = preset.Id
			} else {
				user.PresetId = presets[0].Id
			}

			// находим по айдишнику чата
			database.Database.DB.Find(
				&user,
				models.User{ChatId: user.ChatId},
			)

			// если такого нет - создаем
			if user.ID == 0 {
				database.Database.DB.Create(&user)
			}
		}

		if update.Poll != nil {
			// todo: записывать статистику ответов
			for id, option := range update.Poll.Options {
				if option.VoterCount > 0 {
					fmt.Printf("Selected{id: %v, value:%s}", id, option.Text)
				}
			}
		}
	}
}
