package user_handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/back1ng1/question-bot-api/app/usecase/crud_user"
	"gitlab.com/back1ng1/question-bot-api/entity"
)

type handler struct {
	crud_user_uc crud_user.UseCase
}

func NewHandler(crud_user_uc crud_user.UseCase) RestHandler {
	return &handler{
		crud_user_uc: crud_user_uc,
	}
}

func (h *handler) GetByInterval(c *fiber.Ctx) error {
	// todo rename id to interval
	interval, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	users, err := h.crud_user_uc.GetByInterval(interval)
	if err != nil {
		return err
	}

	if len(users) == 0 {
		return c.JSON([]string{})
	}

	return c.JSON(users)
}

func (h *handler) GetByChatId(c *fiber.Ctx) error {
	chatId, err := strconv.Atoi(c.Params("chat_id"))
	if err != nil {
		return err
	}

	user, err := h.crud_user_uc.GetByChatId(chatId)
	if err != nil {
		return err
	}

	if user.ID == 0 {
		return c.JSON([]string{})
	}

	return c.JSON(user)
}

func (h *handler) Create(c *fiber.Ctx) error {
	var user entity.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	out, err := h.crud_user_uc.Create(user)
	if err != nil {
		return err
	}

	return c.JSON(out)
}

func (h *handler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return err
	}

	user := entity.User{ChatId: int64(id)}
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	out, err := h.crud_user_uc.Update(user)
	if err != nil {
		return err
	}

	return c.JSON(out)
}
