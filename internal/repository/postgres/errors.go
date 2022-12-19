package postgres

import (
	"auth/internal/repository"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

var (
	errUniqueConstraint = "23505"
)

func wrapError(entityName string, err error) error {
	if err == pgx.ErrNoRows {
		return repository.NewNotFoundError("not found: ", entityName, err)
	}

	res := err
	if pgErr, ok := err.(*pgconn.PgError); ok {
		switch pgErr.Code {
		case errUniqueConstraint:
			res = repository.NewUniqueError(fmt.Sprintf("%s already exists, field: ", entityName), pgErr.ColumnName, err)
		default:
			return err
		}
	}

	return res
}
