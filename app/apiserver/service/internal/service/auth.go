package service

import "github.com/tanmingqian/go-tam/pkg/auth"

const (
	// APIServerAudience defines the value of jwt audience field.
	APIServerAudience = "tam.api.tanmingqian.com"

	// APIServerIssuer defines the value of jwt issuer field.
	APIServerIssuer = "tam-apiserver"
)

func newBasicAuth() auth.AuthStrategy {
	return nil
}

func newJWTAuth() auth.AuthStrategy {
	return nil
}
