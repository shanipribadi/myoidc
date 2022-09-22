package storage

import (
	"context"
	"time"
  "errors"

	"gopkg.in/square/go-jose.v2"

	"github.com/zitadel/oidc/v2/pkg/oidc"
	"github.com/zitadel/oidc/v2/pkg/op"
)

func (s *Storage) CreateAuthRequest(context.Context, *oidc.AuthRequest, string) (op.AuthRequest, error) {
	return nil, errNotImplemented
}

func (s *Storage) AuthRequestByID(context.Context, string) (op.AuthRequest, error) {
	return nil, errNotImplemented
}

func (s *Storage) AuthRequestByCode(context.Context, string) (op.AuthRequest, error) {
	return nil, errNotImplemented
}

func (s *Storage) SaveAuthCode(context.Context, string, string) error {
	return errNotImplemented
}
func (s *Storage) DeleteAuthRequest(context.Context, string) error {
	return errNotImplemented
}

func (s *Storage) CreateAccessToken(context.Context, op.TokenRequest) (string, time.Time, error) {
	return "", time.Time{}, errNotImplemented
}
func (s *Storage) CreateAccessAndRefreshTokens(ctx context.Context, request op.TokenRequest, currentRefreshToken string) (accessTokenID string, newRefreshToken string, expiration time.Time, err error) {
	return "", "", time.Time{}, errNotImplemented
}
func (s *Storage) TokenRequestByRefreshToken(ctx context.Context, refreshToken string) (op.RefreshTokenRequest, error) {
	return nil, errNotImplemented

}
func (s *Storage) TerminateSession(ctx context.Context, userID string, clientID string) error {
	return errNotImplemented
}
func (s *Storage) RevokeToken(ctx context.Context, token string, userID string, clientID string) *oidc.Error {
	return oidc.DefaultToServerError(errNotImplemented, "not implemented")
}

func (s *Storage) SigningKey(context.Context) (op.SigningKey, error) {
	return nil, errNotImplemented
}
func (s *Storage) SignatureAlgorithms(context.Context) ([]jose.SignatureAlgorithm, error) {
	return nil, errNotImplemented
}
func (s *Storage) KeySet(context.Context) ([]op.Key, error) {
	return nil, errNotImplemented
}
