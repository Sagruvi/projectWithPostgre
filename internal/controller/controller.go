package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"projectWithPostgre/internal/repository"
	"projectWithPostgre/internal/service"
)

type Controllerer interface {
	GetUser(ctx context.Context, w http.ResponseWriter, r *http.Request) error
	CreateUser(ctx context.Context, w http.ResponseWriter, r *http.Request) error
	UpdateUser(ctx context.Context, w http.ResponseWriter, r *http.Request) error
	DeleteUser(ctx context.Context, w http.ResponseWriter, r *http.Request) error
	ListUsers(ctx context.Context, w http.ResponseWriter, r *http.Request) error
}

type Controller struct {
	service service.Servicer
}
type UserResponse struct {
	users []repository.User
	count int
}

func New(userService service.Servicer) *Controller {
	return &Controller{
		service: userService,
	}
}

// Create godoc
// @Summary Create user
// @Description create new user
// @Tags CREATE
// @Accept  json
// @Produce  json
// @Param id body string true "id"
// @Param username body string true "username"
// @Param password body string true "password"
// @Success 200 {string} string "create new user"
// @Router /api/users/create [post]
func (c *Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user repository.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := c.service.CreateUser(r.Context(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write([]byte(fmt.Sprintf("User %s created", user.Name)))
	w.WriteHeader(http.StatusCreated)
}

// get godoc
// @Summary get user
// @Description get user by id
// @Tags GET user
// @Accept  json
// @Produce  json
// @Param id query string true "id"
// @Success 200 {string} string "get user" repository.User
// @Router /api/users/{id} [get]
func (c *Controller) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	user, err := c.service.GetUser(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// update godoc
// @Summary update user
// @Description update user
// @Tags GET user
// @Accept  json
// @Produce  json
// @Param id body string true "id"
// @Param username body string true "username"
// @Param password body string true "password"
// @Success 200 {string} string "User updated"
// @Router /api/users/update [put]
func (c *Controller) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user repository.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := c.service.UpdateUser(r.Context(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write([]byte("User created"))
	w.WriteHeader(http.StatusAccepted)
}

// delete godoc
// @Summary delete user
// @Description mark user as deleted
// @Tags UPDATE user
// @Accept  json
// @Produce  json
// @Param id query string true "id"
// @Success 200 {string} string "user will be deleted"
// @Router /api/users/{id} [delete]
func (c *Controller) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	err := c.service.DeleteUser(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("User will be deleted"))
}

// list godoc
// @Summary all users
// @Description list of all users
// @Tags GET users
// @Accept  json
// @Produce  json
// @Param limit body string true "limit"
// @Param offset body string true "offset"
// @Success 200 {string} string "users"
// @Router /api/list [get]
func (c *Controller) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var conditions repository.Conditions
	if err := json.NewDecoder(r.Body).Decode(&conditions); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	users, count, err := c.service.ListUsers(r.Context(), conditions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var responseBody UserResponse
	responseBody.users = users
	responseBody.count = count
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseBody)
}
