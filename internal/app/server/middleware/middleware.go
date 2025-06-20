package middleware

import "github.com/danielmesquitta/jwt-playground/internal/pkg/jwtutil"

type Middleware struct {
	j *jwtutil.JWTUtil
}

func NewMiddleware(j *jwtutil.JWTUtil) *Middleware {
	return &Middleware{
		j: j,
	}
}
