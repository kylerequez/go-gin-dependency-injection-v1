package middleware

import "github.com/kylerequez/go-gin-dependency-injection-v1/api/service"

type Middleware struct {
	js *service.JwtService
}

func NewMiddleware(js *service.JwtService) *Middleware {
	return &Middleware{
		js: js,
	}
}
