package strategy

import (
	"context"
	"encoding/base64"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/marmotedu/errors"
	"github.com/tanmingqian/go-tam/pkg/auth"
	"github.com/tanmingqian/go-tam/pkg/code"
	"strings"
)

// BasicStrategy defines Basic authentication strategy.
type BasicStrategy struct {
	compare func(username string, password string) bool
}

var _ auth.AuthStrategy = &BasicStrategy{}

// NewBasicStrategy create basic strategy with compare function.
func NewBasicStrategy(compare func(username string, password string) bool) BasicStrategy {
	return BasicStrategy{
		compare: compare,
	}
}

// AuthFunc defines basic strategy as the gin authentication middleware.
func (b BasicStrategy) AuthFunc() auth.HandlerFunc {
	return func(ctx context.Context) error {
		var header transport.Header
		if tr, ok := transport.FromServerContext(ctx); ok {
			header = tr.RequestHeader()
		}

		auth := strings.SplitN(header.Get("Authorization"), " ", 2)
		if len(auth) != 2 || auth[0] != "Basic" {
			return errors.WithCode(code.ErrSignatureInvalid, "Authorization header format is wrong.")
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 || !b.compare(pair[0], pair[1]) {
			return errors.WithCode(code.ErrSignatureInvalid, "Authorization header format is wrong.")
		}
		return nil
	}
}
