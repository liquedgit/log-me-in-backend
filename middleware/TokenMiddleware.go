package middleware

import (
	"log-me-in/model"
	"log-me-in/service"
	"log-me-in/utils"

	"github.com/gofiber/fiber/v2"
)

func TokenMiddleware(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")
	if len(token) <= 7 {
		return ctx.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponseDTO{
			Error: "You are not authorized to view this content",
		})
	}
	token = token[7:]
	jwtObj, err := utils.DecodeToken(token)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(model.ErrorResponseDTO{
			Error: "You are not authorized to view this content",
		})
	}
	user_id := jwtObj.Claims["id"].(string)
	user, err := service.GetUserById(user_id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponseDTO{
			Error: "There is something wrong with the server",
		})
	}
	// fmt.Println(jwtObj.Claims["id"])
	ctx.Locals("user", user)
	return ctx.Next()
}
