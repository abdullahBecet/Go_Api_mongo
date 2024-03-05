package main

import (
	"ApiGo/app"
	"ApiGo/configs"
	"ApiGo/repository"
	"ApiGo/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	appRoute := fiber.New()
	configs.ConnectDB()
	dbClient := configs.GetCollection(configs.DB, "todos")

	TodoRepostoryDb := repository.NewTodoRepositoryDb(dbClient)

	td := app.TodoHandler{Service: services.NewTodoService(TodoRepostoryDb)}

	appRoute.Post("/api/todo", td.CreateTodo)
	appRoute.Get("/api/todos", td.GetAllTodo)
	appRoute.Delete("/api/todo/:id", td.DeleteTodo)
	appRoute.Listen(":8080")

}
