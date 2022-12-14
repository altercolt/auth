package postgres

import (
	"auth/internal/core/user"
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

var (
	entity = "user"
)

type UserRepository struct {
	db  *pgxpool.Pool
	log *zap.SugaredLogger
}

func NewUserRepository(db *pgxpool.Pool, log *zap.SugaredLogger) user.Repository {
	return UserRepository{
		db:  db,
		log: log,
	}
}

func (u UserRepository) Create(ctx context.Context, m user.Model) error {
	query := `insert into users (username, email, passhash) values ($1, $2)`

	_, err := u.db.Exec(ctx, query, m.Email, m.PassHash)
	if err != nil {
		u.log.Error(err)
		return wrapError(entity, err)
	}

	return nil
}

func (u UserRepository) Update(ctx context.Context, m user.Model) error {
	query := `update users set 
                 email = coalesce($1, email),
                 passhash = coalesce($2, passhash)`

	_, err := u.db.Exec(ctx, query, m.Email, m.PassHash)
	if err != nil {
		u.log.Error(err)
		return wrapError(entity, err)
	}

	return nil
}

func (u UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `delete from users where id = $1`

	tx, err := u.db.Begin(ctx)
	if err != nil {
		u.log.Error(err)
		return wrapError(entity, err)
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, query, id)
	if err != nil {
		u.log.Error(err)
		return wrapError(entity, err)
	}

	if err = tx.Commit(ctx); err != nil {
		u.log.Error(err)
		return wrapError(entity, err)
	}

	return nil
}

func (u UserRepository) FetchOne(ctx context.Context, f user.SingleFilter) (user.User, error) {
	query := `select users.id, users.email, users.passhash, role.name from users
				inner join role on users.role_id = role.id
				`

}

func (u UserRepository) Fetch(ctx context.Context, f user.Filter) ([]user.User, error) {
	query := `select users.id, users.email, users.passhash, role.name from users
				inner join role on users. = role.id where 
				and ($1::uuid[] IS NULL OR users.id = ANY($1)),
				and ($2::varchar[] IS NULL OR users.email = ANY($2))`

	res, err := u.db.Query(ctx, query, f.IDs, f.Email)
	if err != nil {
		u.log.Error(err)
		return nil, wrapError(entity, err)
	}

	var users []user.User
	for res.Next() {
		var usr user.User
		if err := res.Scan(&usr.ID, &usr.Email, &usr.PassHash, &usr.Role); err != nil {
			u.log.Error(err)
			return nil, wrapError(entity, err)
		}
		users = append(users, usr)
	}

	return users, nil
}
