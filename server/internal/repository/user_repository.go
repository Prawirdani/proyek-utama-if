package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prawirdani/golang-restapi/internal/entity"
	"github.com/prawirdani/golang-restapi/pkg/httputil"
)

var (
	ErrorUsernameExists = httputil.ErrConflict("Username already exists")
	ErrorUserNotFound   = httputil.ErrNotFound("User not found")
)

type UserRepository interface {
	InsertUser(ctx context.Context, u entity.User) error
	SelectWhere(ctx context.Context, field string, searchVal any) (entity.User, error)
	Select(ctx context.Context) ([]entity.User, error)
	Update(ctx context.Context, u entity.User) error
}

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(pgpool *pgxpool.Pool) userRepository {
	return userRepository{
		db: pgpool,
	}
}

func (r userRepository) InsertUser(ctx context.Context, u entity.User) error {
	query := "INSERT INTO users(nama, username, password) VALUES($1, $2, $3)"
	_, err := r.db.Exec(ctx, query, u.Nama, u.Username, u.Password)

	if err != nil {
		// Unique constraint duplicate err by PG error code.
		if strings.Contains(err.Error(), "23505") {
			return ErrorUsernameExists
		}
		return err
	}
	return nil
}

func (r userRepository) SelectWhere(ctx context.Context, field string, searchVal any) (entity.User, error) {
	var user entity.User
	query := fmt.Sprintf("SELECT id, nama, username, password, role, aktif, created_at, updated_at FROM users WHERE %s=$1", field)

	err := r.db.QueryRow(ctx, query, searchVal).Scan(
		&user.ID,
		&user.Nama,
		&user.Username,
		&user.Password,
		&user.Role,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return user, ErrorUserNotFound
		}
		return user, err
	}

	return user, nil
}

func (r userRepository) Select(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	query := "SELECT id, nama, username, password, role, aktif, created_at, updated_at FROM users"

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user entity.User
		err := rows.Scan(
			&user.ID,
			&user.Nama,
			&user.Username,
			&user.Password,
			&user.Role,
			&user.Active,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r userRepository) Update(ctx context.Context, u entity.User) error {
	now := time.Now()
	query := "UPDATE users SET nama=$1, username=$2, password=$3, aktif=$4, updated_at=$5 WHERE id=$6"
	_, err := r.db.Exec(ctx, query, u.Nama, u.Username, u.Password, u.Active, now, u.ID)

	if err != nil {
		return err
	}
	return nil
}
