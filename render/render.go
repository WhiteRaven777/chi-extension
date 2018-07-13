package render

import (
	"net/http"

	mw "github.com/WhiteRaven777/chi-extension/middleware"
	"github.com/unrolled/render"
)

func Render(r *http.Request) (ren *render.Render) {
	return r.Context().Value(mw.CtxKey).(*render.Render)
}
