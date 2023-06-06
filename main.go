package main

import (
  "strconv"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

var todo string
var counter int

func main() {
    engine := html.New("./public", ".html") 
    app := fiber.New(fiber.Config{
        Views: engine,
    })

    app.Get("/", func(c *fiber.Ctx) error {
      if c.FormValue("todo") != "" {
        counter += 1 
        todo += strconv.Itoa(counter) + ") " + c.FormValue("todo") + "\n"
      }

      return c.Render("index", fiber.Map{
      "TODO": todo,});
  })

  app.Listen(":8000")
}

