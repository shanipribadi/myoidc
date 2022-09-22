package storage

import (
  "context"

  "gopkg.in/square/go-jose.v2"

  "github.com/zitadel/oidc/v2/pkg/op"
  "github.com/zitadel/oidc/v2/pkg/oidc"
)

func(s *Storage) GetClientByClientID(ctx context.Context, clientID string) (op.Client, error) {
  return nil, errNotImplemented

}
func(s *Storage) 	AuthorizeClientIDSecret(ctx context.Context, clientID, clientSecret string) error {
  return errNotImplemented

}
func(s *Storage) 	SetUserinfoFromScopes(ctx context.Context, userinfo oidc.UserInfoSetter, userID, clientID string, scopes []string) error {
  return errNotImplemented

}
func(s *Storage) 	SetUserinfoFromToken(ctx context.Context, userinfo oidc.UserInfoSetter, tokenID, subject, origin string) error {
  return errNotImplemented

}
func(s *Storage) 	SetIntrospectionFromToken(ctx context.Context, userinfo oidc.IntrospectionResponse, tokenID, subject, clientID string) error {
  return errNotImplemented

}
func(s *Storage) 	GetPrivateClaimsFromScopes(ctx context.Context, userID, clientID string, scopes []string) (map[string]interface{}, error) {
  return nil, errNotImplemented

}

func(s *Storage) 	GetKeyByIDAndUserID(ctx context.Context, keyID, userID string) (*jose.JSONWebKey, error) {
  return nil, errNotImplemented

}

func(s *Storage) 	ValidateJWTProfileScopes(ctx context.Context, userID string, scopes []string) ([]string, error) {
  return nil, errNotImplemented
}
