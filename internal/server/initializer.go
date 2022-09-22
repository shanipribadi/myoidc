package server

import (
	"context"

	"github.com/zitadel/oidc/v2/pkg/op"
	"golang.org/x/text/language"

  "myoidc/internal/storage"
)

func newOP(ctx context.Context, storage op.Storage, key [32]byte, insecure bool, issuer string) (*op.Provider, error) {
	var opts []op.Option
	if insecure {
		opts = append(opts, op.WithAllowInsecure())
	}
    cfg := &op.Config{
			CryptoKey:                key,
			DefaultLogoutRedirectURI: "/login",
			CodeMethodS256:           true,
			AuthMethodPost:           true,
			AuthMethodPrivateKeyJWT:  true,
			GrantTypeRefreshToken:    true,
			RequestObjectSupported:   true,
			SupportedUILocales:       []language.Tag{language.English},
		}
    if len(issuer) > 0 {
      return op.NewOpenIDProvider(ctx, issuer, cfg, storage, opts...)
    }
    return op.NewDynamicOpenIDProvider(ctx, "/", cfg, storage, opts...)
}

func newStorage() *storage.Storage {
  return &storage.Storage{}
}
