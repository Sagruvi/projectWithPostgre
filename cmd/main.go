package main

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
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
		r.Get("/user", func(writer http.ResponseWriter, request *http.Request) {
			userController.GetUser(writer, request)
		})
	})
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(userController.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/user", func(writer http.ResponseWriter, request *http.Request) {
			userController.CreateUser(writer, request)
		})
		r.Delete("/user", func(writer http.ResponseWriter, request *http.Request) {
			userController.DeleteUser(writer, request)
		})

		r.Put("/user", func(writer http.ResponseWriter, request *http.Request) {
			userController.UpdateUser(writer, request)
		})
	})
	r.Group(func(r chi.Router) {
		r.Post("/pet", func(writer http.ResponseWriter, request *http.Request) {
			userController.CreatePet(writer, request)
		})
		r.Get("/pet", func(writer http.ResponseWriter, request *http.Request) {
			userController.FindPetById(writer, request)
		})
		r.Put("/pet", func(writer http.ResponseWriter, request *http.Request) {
			userController.UpdatePet(writer, request)
		})
		r.Get("/pet/findByStatus", func(writer http.ResponseWriter, request *http.Request) {
			userController.FindPetByStatus(writer, request)
		})
		r.Delete("/pet", func(writer http.ResponseWriter, request *http.Request) {
			userController.DeletePet(writer, request)
		})
	})
	r.Group(func(r chi.Router) {
		r.Post("/store/order", func(writer http.ResponseWriter, request *http.Request) {
			userController.StoreOrder(writer, request)
		})
		r.Get("/store/order", func(writer http.ResponseWriter, request *http.Request) {
			userController.GetOrder(writer, request)
		})
		r.Delete("/store/order", func(writer http.ResponseWriter, request *http.Request) {
			userController.DeleteOrder(writer, request)
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
