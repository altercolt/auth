package postgres

import (
	"auth/internal/core/user"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func getConn(ctx context.Context) *pgxpool.Pool {
	dsn := "postgres://altercolt:1952@localhost:5432/auth"
	db, _ := pgxpool.New(ctx, dsn)
	return db
}

func toPtr[T comparable](obj T) *T {
	var def T
	if obj != def {
		return &obj
	}

	return nil
}

func TestUserRepository_Create(t *testing.T) {
	ctx := context.Background()
	db := getConn(ctx)

	userRepo := NewUserRepository(db)
	password, _ := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.DefaultCost)

	err := userRepo.Create(ctx, user.Model{
		Email:    toPtr("aybarrel@gmail.com"),
		Username: toPtr("helloworld"),
		PassHash: toPtr(string(password)),
	})

	if err != nil {
		t.Fatalf("err : %v", err)
	}
}

func TestUserRepository_FetchOne(t *testing.T) {
	ctx := context.Background()
	db := getConn(ctx)
	userRepo := NewUserRepository(db)
	user, err := userRepo.FetchOne(ctx, user.Filter{
		Emails: []string{"helloworld@gmail.com"},
	})
	if err != nil {
		t.Fatalf("err : %v", err)
	}

	t.Logf("user : %v", user)

}
