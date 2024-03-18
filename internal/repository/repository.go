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
	Create(ctx context.Context, user User) error
	GetByID(ctx context.Context, id string) (User, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, c Conditions) ([]User, error)
}

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
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
func (r *Repository) Create(ctx context.Context, user User) error {
	conn, err := Connect(ctx)
	defer conn.Close(context.Background())
	_, err = conn.Exec(ctx, "INSERT INTO users (id, name, password) VALUES ($1, $2, $3)", user.ID, user.Name, user.Password)
	if err != nil {
		return err
	}
	return nil
}
func (r *Repository) GetByID(ctx context.Context, id string) (User, error) {
	conn, err := Connect(ctx)
	defer conn.Close(context.Background())
	var user User
	err = conn.QueryRow(ctx, "SELECT id, name, password FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Password)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
func (r *Repository) Update(ctx context.Context, user User) error {
	conn, err := Connect(ctx)
	defer conn.Close(context.Background())
	_, err = conn.Exec(ctx, "UPDATE users SET name = $1, password = $2 WHERE id = $3", user.Name, user.Password, user.ID)
	if err != nil {
		return err
	}
	return nil
}
func (r *Repository) Delete(ctx context.Context, id string) error {
	conn, err := Connect(ctx)
	defer conn.Close(context.Background())
	_, err = conn.Exec(ctx, "UPDATE users SET deleted_at = now() WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
func (r *Repository) List(ctx context.Context, c Conditions) ([]User, int, error) {
	conn, err := Connect(ctx)
	var count int
	defer conn.Close(context.Background())
	var users []User
	rows, err := conn.Query(ctx, "SELECT id, name, password FROM users LIMIT $1 OFFSET $2", c.Limit, c.Offset)
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Name, &user.Password)
		if err != nil {
			return nil, count, err
		}
		users = append(users, user)
	}
	if err != nil {
		return nil, count, err
	}
	err = conn.QueryRow(context.Background(), "SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		return nil, count, err
	}
	return users, count, nil
}
