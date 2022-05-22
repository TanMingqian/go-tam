package strategy

import (
	"context"
	ginjwt "github.com/appleboy/gin-jwt/v2"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/marmotedu/errors"
	"github.com/tanmingqian/go-tam/pkg/auth"
	"github.com/tanmingqian/go-tam/pkg/code"
	"strings"
)

// AuthzAudience defines the value of jwt audience field.
const AuthzAudience = "tam.authz.tanmingqian.com"

// JWTStrategy defines jwt bearer authentication strategy.
type JWTStrategy struct {
	ginjwt.GinJWTMiddleware
	Authorizator func(data interface{}, ctx context.Context) bool
}

// NewJWTStrategy create jwt bearer strategy with GinJWTMiddleware
func NewJWTStrategy(gjwt ginjwt.GinJWTMiddleware) JWTStrategy {
	return JWTStrategy{gjwt, func(data interface{}, ctx context.Context) bool {
		if _, ok := data.(string); ok {
			return true
		}
		return false
	}}
}

// AuthFunc defines jwt bearer strategy as the gin authentication middleware.
func (j JWTStrategy) AuthFunc() auth.HandlerFunc {
	return func(ctx context.Context) error {
		claims, err := j.getClaimsFromJWT(ctx)
		if err != nil {
			return err
		}

		if claims["exp"] == nil {
			return errors.WithCode(code.ErrMissingExpField, "Authorization header exp field is missing.")
		}

		if _, ok := claims["exp"].(float64); !ok {
			return errors.WithCode(code.ErrSignatureInvalid, "Authorization header format is wrong.")
		}

		if int64(claims["exp"].(float64)) < j.TimeFunc().Unix() {
			return errors.WithCode(code.ErrExpired, "Token is expired")
		}

		identity := claims[j.IdentityKey]

		if !j.Authorizator(identity, ctx) {
			return errors.WithCode(code.ErrPermissionDenied, "Permission denied")
		}
		return nil
	}
}

func (j JWTStrategy) getClaimsFromJWT(ctx context.Context) (map[string]interface{}, error) {
	token, err := j.parseToken(ctx)

	if err != nil {
		return nil, err
	}

	claims := make(map[string]interface{})
	for key, value := range token.Claims.(jwt.MapClaims) {
		claims[key] = value
	}
	return claims, nil
}

func (j JWTStrategy) parseToken(ctx context.Context) (*jwt.Token, error) {
	var token string
	var err error

	methods := strings.Split(j.TokenLookup, ",")
	for _, method := range methods {
		if len(token) > 0 {
			break
		}
		parts := strings.Split(strings.TrimSpace(method), ":")
		k := strings.TrimSpace(parts[0])
		v := strings.TrimSpace(parts[1])
		switch k {
		case "header":
			token, err = j.jwtFromHeader(ctx, v)
		case "query":
			token, err = j.jwtFromQuery(ctx, v)
		case "cookie":
			token, err = j.jwtFromCookie(ctx, v)
		case "param":
			token, err = j.jwtFromParam(ctx, v)
		}
	}

	if err != nil {
		return nil, err
	}
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod(j.SigningAlgorithm) != token.Method {
			return nil, errors.WithCode(code.ErrSignatureInvalid, "Authorization header format is wrong.")
		}

		return j.Key, nil
	})
}

func (j JWTStrategy) jwtFromHeader(ctx context.Context, key string) (string, error) {
	var header transport.Header
	if tr, ok := transport.FromServerContext(ctx); ok {
		header = tr.RequestHeader()
	}
	authHeader := header.Get(key)

	if authHeader == "" {
		return "", errors.WithCode(code.ErrSignatureInvalid, "Authorization header key %s format is wrong.", key)
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == j.TokenHeadName) {
		return "", errors.WithCode(code.ErrSignatureInvalid, "Authorization header key %s format is wrong.", key)
	}

	return parts[1], nil
}

//TODO implement jwtFromQuery
func (j JWTStrategy) jwtFromQuery(ctx context.Context, key string) (string, error) {
	return "", errors.WithCode(code.ErrSignatureInvalid, "No Support")
}

func (j JWTStrategy) jwtFromCookie(ctx context.Context, key string) (string, error) {
	return "", errors.WithCode(code.ErrSignatureInvalid, "No Support")
}

func (j JWTStrategy) jwtFromParam(ctx context.Context, key string) (string, error) {
	return "", errors.WithCode(code.ErrSignatureInvalid, "No Support")
}
