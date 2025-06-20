//go:build wireinject
// +build wireinject

package server

import (
	"github.com/danielmesquitta/jwt-playground/internal/app/server/handler"
	"github.com/danielmesquitta/jwt-playground/internal/app/server/middleware"
	"github.com/danielmesquitta/jwt-playground/internal/app/server/router"
	"github.com/danielmesquitta/jwt-playground/internal/config/env"
	"github.com/danielmesquitta/jwt-playground/internal/domain/usecase"
	"github.com/danielmesquitta/jwt-playground/internal/pkg/jwtutil"
	"github.com/danielmesquitta/jwt-playground/internal/pkg/validator"
	"github.com/danielmesquitta/jwt-playground/internal/provider/cache"
	"github.com/danielmesquitta/jwt-playground/internal/provider/cache/redis"
	"github.com/google/wire"
)

func New() *Server {
	wire.Build(
		env.NewEnv,

		jwtutil.NewJWTUtil,
		wire.Bind(new(validator.Validator), new(*validator.Validation)),
		validator.New,

		wire.Bind(new(cache.Cache), new(*redis.Redis)),
		redis.NewRedis,

		usecase.NewSignInUseCase,
		usecase.NewRefreshUseCase,

		middleware.NewMiddleware,

		handler.NewAuthHandler,
		handler.NewUserHandler,

		router.NewRouter,
		newServer,
	)

	return &Server{}
}
