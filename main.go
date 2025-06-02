package main

import (
	"github.com/gofiber/fiber/v2"
)

// Struct User
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	app := fiber.New()

	// Endpoint GET /user
	app.Get("/user", func(c *fiber.Ctx) error {
		user := User{ID: 1, Name: "Miqdam syiam ", Age: 20}
		return c.JSON(user)
	})

	app.Post("/user", func(c *fiber.Ctx) error {
	var user User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON")
	}

	return c.JSON(fiber.Map{
		"message": "User berhasil dibuat ya",
		"data":    user,
	})
})


	app.Listen(":3000")
}
