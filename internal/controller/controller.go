package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"projectWithPostgre/internal/repository"
	"projectWithPostgre/internal/service"
	"strconv"
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

// CreateUser создает нового пользователя.
// @Summary Create user
// @Description create new user
// @Tags CREATE
// @Accept  json
// @Produce  json
// @Param user body repository.User true "user"
// @Success 200 {string} string "create new user"
// @Router /api/users/create [post]
func (c *Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user repository.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := c.service.CreateUser(r.Context(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewDecoder(r.Body).Decode(&user)
	w.WriteHeader(http.StatusCreated)
}

// GetUser возвращает информацию о пользователе по его имени.
// @Summary get user
// @Description get user by username
// @Tags GET user
// @Accept  json
// @Produce  json
// @Param username query string true "username"
// @Success 200 {string} string "get user" repository.User
// @Router /api/users/{id} [get]
func (c *Controller) GetUser(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("username")
	user, err := c.service.GetUser(r.Context(), name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// UpdateUser обновляет информацию о пользователе.
// @Summary update user
// @Description update user
// @Tags GET user
// @Accept  json
// @Produce  json
// @Param username query string true "username"
// @Success 200 {string} string repository.User
// @Router /api/users/update [put]
func (c *Controller) UpdateUser(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("username")
	user, err := c.service.UpdateUser(r.Context(), name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// DeleteUser помечает пользователя как удаленного.
// @Summary delete user
// @Description mark user as deleted
// @Tags UPDATE user
// @Accept  json
// @Produce  json
// @Param username query string true "username"
// @Success 200 {string} string "user will be deleted"
// @Router /api/users/{id} [delete]
func (c *Controller) DeleteUser(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	err := c.service.DeleteUser(r.Context(), username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("User will be deleted"))
}

// StoreOrder сохраняет новый заказ.
// @Summary Create order
// @Description create new order
// @Tags CREATE
// @Accept  json
// @Produce  json
// @Param order body repository.Order true "Order"
// @Success 200 {string} string "create new order"
// @Router /api/orders/create [post]
func (c *Controller) StoreOrder(w http.ResponseWriter, r *http.Request) {
	var Order repository.Order

	if err := json.NewDecoder(r.Body).Decode(&Order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	Order, err := c.service.StoreOrder(r.Context(), Order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewDecoder(r.Body).Decode(&Order)
	w.WriteHeader(http.StatusCreated)
}

// GetOrder возвращает информацию о заказе по его ID.
// @Summary get order
// @Description get order by id
// @Tags GET order
// @Accept  json
// @Produce  json
// @Param id query string true "id"
// @Success 200 {string} string "get order" repository.Order
// @Router /api/orders/{id} [get]
func (c *Controller) GetOrder(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(query, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	order, err := c.service.GetOrder(r.Context(), int(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

// DeleteOrder помечает заказ как удаленный.
// @Summary delete order
// @Description mark order as deleted
// @Tags UPDATE order
// @Accept  json
// @Produce  json
// @Param id query string true "id"
// @Success 200 {string} string "order will be deleted"
// @Router /api/orders/{id} [delete]
func (c *Controller) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(query, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = c.service.DeleteOrder(r.Context(), int(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Order will be deleted"))
}

// CreatePet создает нового питомца.
// @Summary Create pet
// @Description create new pet
// @Tags CREATE
// @Accept  json
// @Produce  json
// @Param pet body repository.Pet true "Pet"
// @Success 200 {string} string "create new pet"
// @Router /api/pets/create [post]
func (c *Controller) CreatePet(w http.ResponseWriter, r *http.Request) {

	var pet repository.Pet
	if err := json.NewDecoder(r.Body).Decode(&pet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	pet, err := c.service.CreatePet(r.Context(), pet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewDecoder(r.Body).Decode(&pet)
	w.WriteHeader(http.StatusCreated)
}

// UpdatePet обновляет информацию о питомце.
// @Summary update pet
// @Description update pet
// @Tags GET pet
// @Accept  json
// @Produce  json
// @Param id body repository.Pet true "Pet"
// @Success 200 {string} string "Pet updated"
// @Router /api/pets/update [put]
func (c *Controller) UpdatePet(w http.ResponseWriter, r *http.Request) {
	var pet repository.Pet
	if err := json.NewDecoder(r.Body).Decode(&pet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	pet, err := c.service.UpdatePet(r.Context(), pet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewDecoder(r.Body).Decode(&pet)
	w.WriteHeader(http.StatusCreated)
}

// DeletePet помечает питомца как удаленного.
// @Summary delete pet
// @Description mark pet as deleted
// @Tags UPDATE pet
// @Accept  json
// @Produce  json
// @Param id query string true "id"
// @Success 200 {string} string "pet will be deleted"
// @Router /api/pets/{id} [delete]
func (c *Controller) DeletePet(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(query, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = c.service.DeletePet(r.Context(), int(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Pet will be deleted"))
}

// FindPetByStatus возвращает список питомцев по их статусу.
// @Summary Find pets by status
// @Description find pets by status
// @Tags SEARCH pet
// @Accept  json
// @Produce  json
// @Param status query string true "status"
// @Success 200 {string} string "get pets by status" repository.Pet
// @Router /api/pets/status [get]
func (c *Controller) FindPetByStatus(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query().Get("status")
	status := query
	pets, err := c.service.FindPetByStatus(r.Context(), status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pets)
}

// FindPetById возвращает информацию о питомце по его ID.
// @Summary get pet
// @Description get pet by id
// @Tags GET pet
// @Accept  json
// @Produce  json
// @Param id query string true "id"
// @Success 200 {string} string "get pet" repository.Pet
// @Router /api/pets/{id} [get]
func (c *Controller) FindPetById(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(query, 10, 64)
	pets, err := c.service.FindPetById(r.Context(), int(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pets)
}
