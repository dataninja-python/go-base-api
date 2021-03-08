package books

import "github.com/gofiber/fiber"

func GetBooks(c *fiber.Ctx) {
	c.Send("All Books")
}

func GetBook(c *fiber.Ctx) {
	c.Send("A Book")
}

func NewBook(c *fiber.Ctx) {
	c.Send("Adds a new book")
}

func UpdateBook(c *fiber.Ctx) {
	c.Send("Updates a book")
}

func DeleteBook(c *fiber.Ctx) {
	c.Send("Deletes a book")
}
