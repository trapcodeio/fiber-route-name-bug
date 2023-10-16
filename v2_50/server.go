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

	// snake case param works (name is registered)
	app.Get("/a/:a_id", handleRequest).Name("a")
	// camel case param doesn't work (name is not registered)
	app.Post("/b/:bId", handleRequest).Name("b")

	// group without param
	c := app.Group("/c")
	c.Get("", handleRequest).Name("c.get")
	// for some reason `c.get` is not registered and `c.post` is registered
	// if you comment out the `c.post` route, `c.get` will be registered
	c.Post("", handleRequest).Name("c.post")
	// this works as expected
	c.Get("/d", handleRequest).Name("c.get.d")

	// group with params
	d := app.Group("/d/:d_id")
	// works as expected
	d.Get("", handleRequest).Name("d.get")
	// for some reason `d.get` is not registered and `d.post` is registered
	// if you comment out the `d.post` route, `d.get` will be registered
	d.Post("", handleRequest).Name("d.post")

	// group with camel case param
	e := app.Group("/e/:eId")
	// All route names in this group are not registered
	e.Get("", handleRequest).Name("e.get")
	e.Post("", handleRequest).Name("e.post")
	e.Get("f", handleRequest).Name("e.get.f")

	// using real world example
	postGroup := app.Group("/post/:postId")
	// All route names in this group are not registered just like the `e` group
	postGroup.Get("", handleRequest).Name("post.get")
	postGroup.Post("", handleRequest).Name("post.update")

	err := app.Listen(":3002")
	if err != nil {
		panic(err)
	}
}
