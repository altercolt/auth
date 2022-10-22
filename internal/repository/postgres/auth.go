package postgres

import (
	"auth/internal/core/auth"
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TokenRepository struct {
	db *pgxpool.Pool
}

func NewTokenRepository(db *pgxpool.Pool) auth.TokenRepository {
	return TokenRepository{
		db: db,
	}
}

func (r TokenRepository) Create(ctx context.Context, token auth.RefreshToken) error {
	query := `INSERT INTO tokens (id, user_id, refresh_token, expiration_time) 
						VALUES ($1, $2, $3, $4); `

	_, err := r.db.Exec(ctx, query, &token.ID, &token.UserID, &token.RefreshToken, &token.ExpirationTime)
	if err != nil {
		return err
	}

	return nil
}

func (r TokenRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM tokens WHERE id = $1`
	_, err := r.db.Exec(ctx, query, &id)
	if err != nil {
		return err
	}

	return nil
}

func (r TokenRepository) Fetch(ctx context.Context, filter auth.Filter) ([]auth.RefreshToken, error) {
	query := `SELECT * FROM tokens WHERE
                         $1::uuid IS NULL OR id = ANY($1)
						AND $2::int IS NULL OR user_id = ANY($2)
						AND $3::varchar IS NULL OR refresh_token = ANY($3);
                         `
	curr, err := r.db.Query(ctx, query, filter.IDs, filter.Users, filter.Tokens)
	if err != nil {
		return nil, err
	}

	var result []auth.RefreshToken

	for curr.Next() {
		var token auth.RefreshToken
		err = curr.Scan(&token.ID, &token.UserID, &token.RefreshToken, &token.ExpirationTime)
		if err != nil {
			return nil, err
		}
		result = append(result, token)
	}

	return result, nil
}

func (r TokenRepository) FetchOne(ctx context.Context, filter auth.Filter) (auth.RefreshToken, error) {
	query := `SELECT * FROM tokens WHERE 
                         $1 IS NOT NULL OR id = ANY($1)
                         AND $2 IS NOT NULL OR refresh_token = ANY($2);`

	res := r.db.QueryRow(ctx, query, filter.IDs, filter.Tokens)
	var token auth.RefreshToken
	if err := res.Scan(&token.ID, &token.UserID, &token.RefreshToken, &token.ExpirationTime); err != nil {
		return auth.RefreshToken{}, err
	}

	return token, nil
}
