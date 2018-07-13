# chi-extension
This package is an extension of go-chi.
If you want to build the REST API on AWS Lambda in a single binary, using this package makes implementation easier.

# How to use.
go get -u github.com/WhiteRaven777/chi-extension

## code sample
``` go
import (
	mw "github.com/WhiteRaven777/chi-extension/middleware"
	"github.com/WhiteRaven777/chi-extension/render"
	"github.com/go-chi/chi"
)

r := chi.NewRouter()

r.Use(mw.Middleware)
r.Get("/", func(w http.ResponseWriter, r *http.Request) {
    render.Render(r).JSON(w, http.StatusOK, map[string]string{"message": "hello"})
})
```