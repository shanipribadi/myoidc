package storage

import (
	"context"

	"github.com/zitadel/oidc/v2/pkg/op"
)

type Storage struct{}

var _ op.Storage = (*Storage)(nil)
var _ op.OPStorage = (*Storage)(nil)
var _ op.AuthStorage = (*Storage)(nil)

func (s *Storage) Health(context.Context) error {
  return nil
}

var errNotImplemented = errors.New("not implemented")

