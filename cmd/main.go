package main

import (
	"context"
	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"log"
	"net/http"
	"os"
	_ "projectWithPostgre/cmd/docs"
	"projectWithPostgre/internal/controller"
	"projectWithPostgre/internal/repository"
	"projectWithPostgre/internal/service"
)

//	@title			CRUD API
//	@version		1.0
//	@description	This is a sample CRUD api.

// @host	localhost:8080/
func main() {

	// Инициализация репозитория
	repo := repository.New()
	repo.Migrate(context.Background())
	// Инициализация сервиса
	userService := service.New(repo)

	// Инициализация контроллера
	userController := controller.New(userService)
	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
		))
		r.Get("/api/users/", func(writer http.ResponseWriter, request *http.Request) {
			userController.GetUser(writer, request)
		})
		r.Post("/api/users/create", func(writer http.ResponseWriter, request *http.Request) {
			userController.CreateUser(writer, request)
		})
		r.Delete("/api/users/", func(writer http.ResponseWriter, request *http.Request) {
			userController.DeleteUser(writer, request)
		})

		r.Put("/api/users/update", func(writer http.ResponseWriter, request *http.Request) {
			userController.UpdateUser(writer, request)
		})
		r.Get("/api/list", func(writer http.ResponseWriter, request *http.Request) {
			userController.GetAllUsers(writer, request)
		})
	})

	// Запуск сервера
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server started on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
