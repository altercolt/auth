package postgres

import (
	"auth/internal/core/user"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) user.Repository {
	return UserRepository{
		db: db,
	}
}

func (r UserRepository) Create(ctx context.Context, m user.Model) error {
	query := `INSERT INTO users (email, username, passhash) 
						VALUES($1, $2, $3);`

	_, err := r.db.Exec(ctx, query, m.Email, m.Username, m.PassHash)
	if err != nil {
		return err
	}

	return nil
}

func (r UserRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id = $1`

	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r UserRepository) Update(ctx context.Context, m user.Model) error {
	query := `UPDATE users SET 
                 email = COALESCE($2)
                 username = COALESCE($3)
                 passhash = COALESCE($4)
                 WHERE id = $1`

	_, err := r.db.Exec(ctx, query, m.ID, m.Email, m.Username, m.PassHash)
	if err != nil {
		return err
	}

	return nil
}

func (r UserRepository) Fetch(ctx context.Context, filter user.Filter) ([]user.User, error) {
	query := `SELECT users.id, roles.name, users.email, users.username, users.passhash FROM users INNER JOIN roles ON users.role_id = roles.id
				WHERE $1 IS NOT NULL OR users.id = ANY($1)
				AND $2 IS NOT NULL OR users.email = ANY($2)
				AND $3 IS NOT NULL OR users.username = ANY($3);`

	curr, err := r.db.Query(ctx, query, filter.IDs, filter.Emails, filter.Usernames)
	if err != nil {
		return nil, err
	}

	var res []user.User
	for curr.Next() {
		var usr user.User
		if err = curr.Scan(&usr.ID, &usr.Role, &usr.Email, &usr.Username, &usr.PassHash); err != nil {
			return nil, err
		}
		res = append(res, usr)
	}

	return res, nil
}

func (r UserRepository) FetchOne(ctx context.Context, filter user.Filter) (user.User, error) {
	query := `SELECT users.id, roles.name, users.email, users.username, users.passhash FROM users INNER JOIN roles ON users.role_id = roles.id
				WHERE $1 IS NOT NULL OR users.id = ANY($1)
				AND $2 IS NOT NULL OR users.email = ANY($2)
				AND $3 IS NOT NULL OR users.username = ANY($3);`

	res := r.db.QueryRow(ctx, query, filter.IDs, filter.Emails, filter.Usernames)

	var usr user.User
	if err := res.Scan(&usr.ID, &usr.Role, &usr.Email, &usr.Username, &usr.PassHash); err != nil {
		return user.User{}, err
	}

	return usr, nil
}
