package ctxutil

import (
	"context"

	"github.com/rs/zerolog/log"
)

type contextKey struct {
	name string
}

var userContextKey = contextKey{name: "user"}

type User = interface{}

func SetUser(ctx context.Context, user User) context.Context {
	return context.WithValue(ctx, userContextKey, user)
}

func GetUser(ctx context.Context) User {
	user, ok := ctx.Value(userContextKey).(User)
	if !ok {
		log.Panic().Msg("could not find user in context")
	}

	return user
}
