package app

import (
	services "ApiGo/mocks/service"
	"ApiGo/models"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var td TodoHandler

func TestTodoHandler_GetAllTodo(t *testing.T) {
	var FakeDataForHandler = []models.Todo{
		{primitive.NewObjectID(), "Title 1", "Content1"},
		{primitive.NewObjectID(), "Title 2", "Content2"},
		{primitive.NewObjectID(), "Title 3", "Content 3"},
	}
	ctrl := gomock.NewController(t)
	mockService := services.NewMockTodoService(ctrl)
	td = TodoHandler{mockService}
	router := fiber.New()
	router.Get("/api/todos", td.GetAllTodo)

	mockService.EXPECT().TodoGetAll().Return(FakeDataForHandler, nil)
	req := httptest.NewRequest("GET", "/api/todos", nil)
	resp, _ := router.Test(req, 1)
	assert.Equal(t, 200, resp.StatusCode)

}
