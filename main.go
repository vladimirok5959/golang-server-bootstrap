package main

import (
	"context"
	"net/http"
	"time"

	"github.com/vladimirok5959/golang-server-bootstrap/bootstrap"
)

func main() {
	bootstrap.Start(
		&bootstrap.Opts{
			Host:    "127.0.0.1:8080",
			Timeout: 8 * time.Second,
			Path:    "assets",
			Before: func(ctx context.Context, w http.ResponseWriter, r *http.Request, o interface{}) {
				// Before callback
				w.Header().Set("Cache-Control", "public, max-age=31536000")
			},
			After: func(ctx context.Context, w http.ResponseWriter, r *http.Request, o interface{}) {
				// After callback
				w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
				w.Header().Set("Content-Type", "text/html")
				w.Write([]byte(`
					<div>Hello World!</div>
					<div><a href="/assets/bootstrap.css">/assets/bootstrap.css</a></div>
					<div><a href="/assets/bootstrap.js">/assets/bootstrap.js</a></div>
					<div><a href="/assets/jquery.js">/assets/jquery.js</a></div>
					<div><a href="/assets/popper.js">/assets/popper.js</a></div>
				`))
			},
		},
	)
}
