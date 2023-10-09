package main

import (
	"github.com/gofiber/fiber/v2"
)

func handleRequest(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"method": c.Method(),
		"url":    c.OriginalURL(),
		"name":   c.Route().Name,
		"params": c.AllParams(),
	})
}

func main() {
	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
	})

	app.Get("/", handleRequest).Name("index")

	app.Get("/a/:a_id", handleRequest).Name("a")
	app.Post("/b/:bId", handleRequest).Name("b")

	// group without param
	c := app.Group("/c")
	c.Get("", handleRequest).Name("c.get")
	c.Post("", handleRequest).Name("c.post")
	c.Get("/d", handleRequest).Name("c.get.d")

	// group with params
	d := app.Group("/d/:d_id")
	d.Get("", handleRequest).Name("d.get")
	d.Post("", handleRequest).Name("d.post")

	// group with camel case param
	e := app.Group("/e/:eId")
	e.Get("", handleRequest).Name("e.get")
	e.Post("", handleRequest).Name("e.post")
	e.Get("f", handleRequest).Name("e.get.f")

	// using real world example
	postGroup := app.Group("/post/:postId")
	postGroup.Get("", handleRequest).Name("post.get")
	postGroup.Post("", handleRequest).Name("post.update")

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
