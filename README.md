# golang-server-bootstrap
Simple http server wrapper + bootstrap + jquery

## What inside
* Smooth server shutdown
* Bootstrap v4.1.3 mounted as resource
* jQuery v3.3.1 mounted as resource
* Popper mounted as resource

## How to use
```
go get github.com/vladimirok5959/golang-server-bootstrap
```
```
package main

import (
	"context"
	"net/http"
	"time"

	"github.com/vladimirok5959/golang-server-bootstrap/bootstrap"
)

func main() {
	bootstrap.Start(
		context.Background(),
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
```
Extra headers can be located in **before** callback. For example for cache control or server name. Logic can be located in **after** callback. If mounted resource file will pushed, **after** callback will not fired. Mounted resources in priority.
```
type callbackBeforeAfter func(ctx context.Context, w http.ResponseWriter, r *http.Request, o interface{})
func Start(ctx context.Context, opts *bootstrap.Opts)
```
Where **host** is server ip and port (127.0.0.1:8080), **timeout** is time in seconds to force server shutdown if server don't want stopping, **path** is virtual path for mounted files (for example if path will be "assets", then bootstrap css file will be located at http://127.0.0.1:8080/assets/bootstrap.css or if "system/assets", then http://127.0.0.1:8080/system/assets/bootstrap.css), **before** and **after** is callback functions for integration.

## In result
This files will be available:
* http://127.0.0.1:8080/assets/bootstrap.css
* http://127.0.0.1:8080/assets/bootstrap.js
* http://127.0.0.1:8080/assets/jquery.js
* http://127.0.0.1:8080/assets/popper.js
