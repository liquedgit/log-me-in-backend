package routes

import (
	"log-me-in/middleware"
	"log-me-in/model"
	"log-me-in/service"

	"github.com/gofiber/fiber/v2"
)

func SetupNoteRoutes(api fiber.Router) {
	note := api.Group("/note")
	note.Use(middleware.TokenMiddleware)
	note.Get("/", func(ctx *fiber.Ctx) error {
		user := ctx.Locals("user").(*model.User)
		notes, err := service.GetNotesByUserid(user.Id)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponseDTO{
				Error: "There is something wrong with the server",
			})
		}

		var dtoNotes = make([]model.NoteResponseDTO, 0)
		for _, note := range *notes {
			dtoNote := model.NoteResponseDTO{
				Id:          note.Id,
				UserId:      note.UserId,
				Title:       note.Title,
				Description: note.Description,
				ImageUrl:    note.ImageUrl,
			}
			dtoNotes = append(dtoNotes, dtoNote)
		}

		return ctx.Status(fiber.StatusOK).JSON(model.GetUserNotesResponseDTO{
			Notes: &dtoNotes,
		})

	})

	note.Get("/:noteid", func(ctx *fiber.Ctx) error {
		noteid := ctx.Params("noteid")
		noteDetails, err := service.GetNoteDetailsById(noteid)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(model.ErrorResponseDTO{
				Error: "There is something wrong with the server",
			})
		}
		return ctx.Status(fiber.StatusOK).JSON(model.NoteResponseDTO{
			Id:          noteDetails.Id,
			UserId:      noteDetails.UserId,
			Title:       noteDetails.Title,
			Description: noteDetails.Description,
			ImageUrl:    noteDetails.ImageUrl,
		})
	})

	note.Post("/", func(ctx *fiber.Ctx) error {
		user := ctx.Locals("user").(*model.User)
		form, err := ctx.MultipartForm()
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponseDTO{
				Error: err.Error(),
			})
		}
		response, err := service.CreateNote(form, user)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponseDTO{
				Error: err.Error(),
			})
		}

		return ctx.Status(fiber.StatusCreated).JSON(response)
	})
}
