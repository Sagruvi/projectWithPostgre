package repository

import (
	"context"
	"database/sql"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"log"
	"os"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user User) error
	GetUser(ctx context.Context, name string) (User, error)
	UpdateUser(ctx context.Context, name string) (User, error)
	DeleteUser(ctx context.Context, name string) error
	StoreOrder(ctx context.Context, order Order) (Order, error)
	GetOrder(ctx context.Context, id int) (Order, error)
	DeleteOrder(ctx context.Context, id int) error
	CreatePet(ctx context.Context, pet Pet) (Pet, error)
	UpdatePet(ctx context.Context, pet Pet) (Pet, error)
	FindPetByStatus(ctx context.Context, status string) (Pet, error)
	FindPetById(ctx context.Context, id int) (Pet, error)
	DeletePet(ctx context.Context, id int) error
}

type User struct {
	ID         int    `json:"id"`
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	UserStatus int    `json:"userstatus"`
}
type Pet struct {
	id       int             `json:"id"`
	category Category        `json:"category"`
	photourl map[string]bool `json:"photoUrls"`
	tags     map[Tag]bool    `json:"tags"`
	name     string          `json:"name"`
	status   string          `json:"status"`
}
type Tag struct {
	id   int    `json:"id"`
	name string `json:"name"`
}
type Category struct {
	id   int    `json:"id"`
	name string `json:"name"`
}
type Order struct {
	id       int    `json:"id"`
	petId    int    `json:"petId"`
	quantity int    `json:"quantity"`
	shipDate string `json:"shipDate"`
	status   string `json:"status"`
	complete bool   `json:"complete"`
}
type Repository struct {
	conn *pgx.Conn
}
type Conditions struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

func New() *Repository {
	return &Repository{}
}
func (r *Repository) Migrate(ctx context.Context) error {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	db, err := sql.Open("postgres", string("host="+os.
		Getenv("DB_HOST")+" port="+os.
		Getenv("DB_PORT")+" user="+os.
		Getenv("DB_USER")+" password="+os.
		Getenv("DB_PASSWORD")+" dbname="+os.
		Getenv("DB_NAME")+" sslmode=disable"))
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}
	defer db.Close()

	if err := goose.Up(db, "migrations/"); err != nil {
		log.Fatalf("Error applying migrations: %s", err)
	}
	return err
}
func Connect(ctx context.Context) (*pgx.Conn, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	conn, err := pgx.Connect(ctx, string("host="+os.
		Getenv("DB_HOST")+" port="+os.
		Getenv("DB_PORT")+" user="+os.
		Getenv("DB_USER")+" password="+os.
		Getenv("DB_PASSWORD")+" dbname="+os.
		Getenv("DB_NAME")+" sslmode=disable"))
	if err != nil {
		return nil, err
	}
	return conn, nil
}
func (r *Repository) CreateUser(ctx context.Context, user User) (User, error) {
	conn, err := Connect(ctx)
	defer conn.Close(context.Background())
	_, err = conn.Exec(ctx, "INSERT INTO users (id, name, password, userstatus, "+
		"email, phone, lastname) VALUES ($1, $2, $3)",
		user.ID, user.FirstName, user.Password,
		user.UserStatus, user.Email, user.Phone, user.LastName)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
func (r *Repository) GetUser(ctx context.Context, name string) (User, error) {
	conn, err := Connect(ctx)
	defer conn.Close(context.Background())
	var user User
	err = conn.QueryRow(ctx, "SELECT id, name, password, userstatus, email, phone, lastname"+
		" FROM users WHERE name = $1", name).Scan(&user.ID,
		&user.FirstName, &user.Password, &user.UserStatus, &user.Email, &user.Phone, &user.LastName)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
func (r *Repository) UpdateUser(ctx context.Context, name string) (User, error) {
	conn, err := Connect(ctx)
	defer conn.Close(context.Background())
	_, err = conn.Exec(ctx, "UPDATE users SET name = $1", name)
	if err != nil {
		return User{}, err
	}
	var res User
	err = conn.QueryRow(ctx, "SELECT id, name, password, userstatus, email, phone, lastname"+
		" FROM users WHERE name = $1", name).Scan(&res.ID,
		&res.FirstName, &res.Password, &res.UserStatus, &res.Email, &res.Phone, &res.LastName)
	if err != nil {
		return User{}, err
	}
	return res, nil
}
func (r *Repository) DeleteUser(ctx context.Context, name string) error {
	conn, err := Connect(ctx)
	defer conn.Close(context.Background())
	_, err = conn.Exec(ctx, "UPDATE users SET deleted_at = now() WHERE name = $1", name)
	if err != nil {
		return err
	}
	return nil
}
func (r *Repository) StoreOrder(ctx context.Context, order Order) (Order, error) {
	conn, err := Connect(ctx)
	defer conn.Close(context.Background())
	_, err = conn.Exec(ctx, "INSERT INTO orders (id, petId, quantity, shipDate, status, complete)"+
		" VALUES ($1, $2, $3, $4, $5, $6)",
		order.id, order.petId, order.quantity, order.shipDate, order.status, order.complete)
	if err != nil {
		return Order{}, err
	}
	return order, nil
}

func (r *Repository) GetOrder(ctx context.Context, id int) (Order, error) {
	conn, err := Connect(ctx)
	defer conn.Close(context.Background())
	var order Order
	err = conn.QueryRow(ctx, "SELECT id, petId, quantity, shipDate, status, complete"+
		" FROM orders WHERE id = $1", id).Scan(&order.id, &order.petId, &order.quantity, &order.shipDate, &order.status, &order.complete)
	if err != nil {
		return Order{}, err
	}
	return order, nil
}
func (r *Repository) DeleteOrder(ctx context.Context, id int) error {
	conn, err := Connect(ctx)
	defer conn.Close(context.Background())
	_, err = conn.Exec(ctx, "DELETE FROM orders WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
func (r *Repository) CreatePet(ctx context.Context, pet Pet) (Pet, error) {
	conn, err := Connect(ctx)
	defer conn.Close(context.Background())
	_, err = conn.Exec(ctx, "INSERT INTO pets (id, category, photoUrls, tags, name, status)"+
		" VALUES ($1, $2, $3, $4, $5, $6)",
		pet.id, pet.category, pet.photourl, pet.tags, pet.name, pet.status)
	if err != nil {
		return Pet{}, err
	}
	return pet, nil
}
func (r *Repository) UpdatePet(ctx context.Context, pet Pet) (Pet, error) {
	conn, err := Connect(ctx)
	defer conn.Close(context.Background())
	_, err = conn.Exec(ctx, "UPDATE pets SET name = $1", pet.name)
	if err != nil {
		return Pet{}, err
	}
	return pet, nil
}
func (r *Repository) FindPetByStatus(ctx context.Context, status string) (Pet, error) {
	conn, err := Connect(ctx)
	defer conn.Close(context.Background())
	var pet Pet
	err = conn.QueryRow(ctx, "SELECT id, category, photoUrls, tags, name, status"+
		" FROM pets WHERE status = $1", status).Scan(&pet.id, &pet.category, &pet.photourl, &pet.tags, &pet.name, &pet.status)
	if err != nil {
		return Pet{}, err
	}
	return pet, nil
}
func (r *Repository) FindPetById(ctx context.Context, id int) (Pet, error) {
	conn, err := Connect(ctx)
	defer conn.Close(context.Background())
	var pet Pet
	err = conn.QueryRow(ctx, "SELECT id, category, photoUrls, tags, name, status"+
		" FROM pets WHERE id = $1", id).Scan(&pet.id, &pet.category, &pet.photourl, &pet.tags, &pet.name, &pet.status)
	if err != nil {
		return Pet{}, err
	}
	return pet, nil
}

func (r *Repository) DeletePet(ctx context.Context, id int) error {
	conn, err := Connect(ctx)
	defer conn.Close(context.Background())
	_, err = conn.Exec(ctx, "DELETE FROM pets WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
