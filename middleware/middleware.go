package middleware

import (
	"context"
	"net/http"

	"github.com/unrolled/render"
)

const (
	CtxKey = "render"
)

var ren *render.Render

func init() {
	ren = render.New()
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), CtxKey, ren)))
	})
}
