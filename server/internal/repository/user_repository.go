package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prawirdani/golang-restapi/internal/entity"
	"github.com/prawirdani/golang-restapi/pkg/httputil"
)

var (
	ErrorEmailExists  = httputil.ErrConflict("Email already exists")
	ErrorUserNotFound = httputil.ErrNotFound("User not found")
)

type UserRepository interface {
	InsertUser(ctx context.Context, u entity.User) error
	SelectWhere(ctx context.Context, field string, searchVal any) (entity.User, error)
}

type userRepository struct {
	tableName string
	db        *pgxpool.Pool
}

func NewUserRepository(pgpool *pgxpool.Pool, tableName string) userRepository {
	return userRepository{
		tableName: tableName,
		db:        pgpool,
	}
}

func (r userRepository) InsertUser(ctx context.Context, u entity.User) error {
	query := fmt.Sprintf("INSERT INTO %s(id, name, email, password) VALUES($1, $2, $3, $4)", r.tableName)
	_, err := r.db.Exec(ctx, query, u.ID, u.Name, u.Email, u.Password)

	if err != nil {
		// Unique constraint duplicate err by PG error code.
		if strings.Contains(err.Error(), "23505") {
			return ErrorEmailExists
		}
		return err
	}
	return nil
}

func (r userRepository) SelectWhere(ctx context.Context, field string, searchVal any) (entity.User, error) {
	var user entity.User
	query := fmt.Sprintf("SELECT id, name, email, password, created_at, updated_at FROM %s WHERE %s=$1", r.tableName, field)

	err := r.db.QueryRow(ctx, query, searchVal).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
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
