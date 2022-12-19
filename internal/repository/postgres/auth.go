package postgres

import (
	"auth/internal/core/auth"
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type TokenRepository struct {
	db  *pgxpool.Pool
	log *zap.SugaredLogger
}

func NewTokenRepository(db *pgxpool.Pool, log *zap.SugaredLogger) auth.TokenRepository {
	return TokenRepository{
		db:  db,
		log: log,
	}
}

func (r TokenRepository) Create(ctx context.Context, token auth.RefreshToken) error {
	query := `insert into tokens (user_id, refresh_token, expiration_time)
									values ($1, $2, $3)`

	_, err := r.db.Exec(ctx, query, token.UserID, token.RefreshToken, token.ExpirationTime)
	if err != nil {
		return wrapError()
	}

	return nil
}

func (r TokenRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `delete from tokens where id = $1`

	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r TokenRepository) FetchOne(ctx context.Context, singleFilter auth.SingleFilter) (auth.RefreshToken, error) {
	query := `select * from tokens 
         		where `

}

func (r TokenRepository) Fetch(ctx context.Context, filter auth.Filter) ([]auth.RefreshToken, error) {
	query := `select * from tokens 
         where 
             `

	res, err := r.db.Query(ctx, query, filter.IDs, filter.Users, filter.Tokens)
	if err != nil {
		return nil, err
	}

	var tokens []auth.RefreshToken

	for res.Next() {
		var token auth.RefreshToken

		if err = res.Scan(&token.ID, &token.UserID, &token.RefreshToken, &token.ExpirationTime); err != nil {
			return nil, err
		}

		tokens = append(tokens, token)
	}

	return tokens, nil
}
