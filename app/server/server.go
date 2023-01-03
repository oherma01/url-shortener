package server

import (
	"strconv"
	model "url-shortener/app/Model"
	tools "url-shortener/app/tools"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setupAndListen() {
	router := fiber.New()

	// setup routes
	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.Get("/r/:redirect", redirect)

	router.Get("/url-shortener", getAllRedirects)
	router.Get("/url-shortener/:id", getRedirect)
	router.Post("/url-shortener", createShortened)
	router.Patch("/url-shortener/:id", updateShortened)
	router.Delete("/url-shortener/:id", deleteShortened)
	router.Listen(":3000")

}

func redirect(ctx *fiber.Ctx) error {
	// get the shortened URL from the request
	shortened_url := ctx.Params("shortened")
	shortened, err := model.GetShortenedByShortURL(shortened_url)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error getting shortened url" + err.Error(),
		})
	}

	// grab properties from the shortened URL
	shortened.Clicked++

	err = model.UpdateShortened(shortened)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error updating shortened url" + err.Error(),
		})
	}

	// redirect to the original URL
	return ctx.Redirect(shortened.Redirect, fiber.StatusTemporaryRedirect)

}

func getAllRedirects(ctx *fiber.Ctx) error {
	shortened_urls, err := model.GetAllShortened()

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error getting shortened urls" + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(shortened_urls)

}

func getRedirect(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)

	// TODO: add error logging
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error parsing id" + err.Error(),
		})
	} else if id == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error parsing id, id is 0",
		})
	}

	shortened_url, err := model.GetShortened(id)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error getting shortened url" + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(shortened_url)

}

func createShortened(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	var shortened model.ShortURL
	err := ctx.BodyParser(&shortened)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error parsing body" + err.Error(),
		})
	}

	if shortened.Random {
		shortened.Short = tools.RandomURL(8)
	}

	err = model.CreateShortened(shortened)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creating shortened url" + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(shortened)

}

func updateShortened(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	var shortened model.ShortURL
	err := ctx.BodyParser(&shortened)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error parsing body" + err.Error(),
		})
	}

	err = model.UpdateShortened(shortened)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error updating shortened url" + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(shortened)

}

func deleteShortened(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error parsing id" + err.Error(),
		})
	}

	err = model.DeleteShortened(id)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error deleting shortened url" + err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Deleted shortened url with id " + ctx.Params("id"),
	})
}
