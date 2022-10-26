package service

import (
	"auth/internal/core/auth"
	"auth/internal/core/user"
	"auth/pkg/keystore"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"net/mail"
)

type AuthService struct {
	kid         string
	log         *zap.SugaredLogger
	userService user.Service
	tokenRepo   auth.TokenRepository
	keyStore    *keystore.KeyStore
}

func NewAuthService(log *zap.SugaredLogger, keystore *keystore.KeyStore, userService user.Service, tokenRepo auth.TokenRepository, kid string) auth.Service {
	return &AuthService{
		log:         log,
		userService: userService,
		tokenRepo:   tokenRepo,
		keyStore:    keystore,
		kid:         kid,
	}
}

func (a *AuthService) Login(ctx context.Context, login auth.Login) (auth.TokenPair, error) {
	var usr user.User
	var err error
	if _, ok := mail.ParseAddress(login.Login); ok != nil {
		usr, err = a.userService.GetOneByUsername(ctx, login.Login)
	} else {
		usr, err = a.userService.GetOneByEmail(ctx, login.Login)
	}

	if err != nil {
		return auth.TokenPair{}, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(usr.PassHash), []byte(login.Password)); err != nil {
		return auth.TokenPair{}, err
	}

	accessPayload := auth.NewPayload(usr.ID, string(usr.Role))
	at := jwt.NewWithClaims(jwt.SigningMethodRS256, &accessPayload)

	pkey, err := a.keyStore.PrivateKey(a.kid)
	if err != nil {
		return auth.TokenPair{}, fmt.Errorf("error when fetching private key [kid = %s] : %w", a.kid, err)
	}

	accessToken, err := at.SignedString(pkey)
	if err != nil {
		return auth.TokenPair{}, fmt.Errorf("error when signing access token : %w", err)
	}

	refreshPayload := auth.NewRefreshPayload(usr.ID)
	rt := jwt.NewWithClaims(jwt.SigningMethodRS256, &refreshPayload)

	refreshToken, err := rt.SignedString(pkey)
	if err != nil {
		return auth.TokenPair{}, fmt.Errorf("error when signing refresh token : %w", err)
	}

	return auth.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		AccessExp:    accessPayload.Exp,
		RefreshExp:   refreshPayload.Exp,
	}, nil
}

func (a *AuthService) ValidateAccess(ctx context.Context, accessToken string) (auth.Payload, error) {

	keyFunc := func(*jwt.Token) (interface{}, error) {
		return a.keyStore.PublicKey(a.kid)
	}

	tkn, err := jwt.ParseWithClaims(accessToken, &auth.Payload{}, keyFunc)
	if err != nil {
		return auth.Payload{}, fmt.Errorf("error when authorizing user : %w", err)
	}

	payload, ok := tkn.Claims.(*auth.Payload)
	if !ok || !tkn.Valid {
		return auth.Payload{}, auth.ErrInvalidAccessToken
	}

	return *payload, nil
}
