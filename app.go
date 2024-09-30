package test1

import "github.com/gofiber/fiber/v3"

func NewApp(config fiber.Config) *fiber.App {
	return fiber.New(config)
}
