package auth

import (
	"context"
)

type HandlerFunc func(ctx context.Context) error

type AuthStrategy interface {
	AuthFunc() HandlerFunc
}

type AuthOperator struct {
	strategy AuthStrategy
}

func (o *AuthOperator) SetStrategy(strategy AuthStrategy) {
	o.strategy = strategy
}

func (o *AuthOperator) AuthFunc() HandlerFunc {
	return o.strategy.AuthFunc()
}
