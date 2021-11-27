package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/mbazhlekova/go-todo/db"
	"github.com/mbazhlekova/go-todo/todo"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	db.ConnectDB()
	defer db.DB.Close()

	api := app.Group("/api")
	todo.Register(api, db.DB)

	log.Fatal(app.Listen(":5000"))
}
