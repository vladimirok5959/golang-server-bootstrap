package bootstrap

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

const path = "assets"

type someTestStruct struct {
	Name string
}

func handle() http.Handler {
	obj := &someTestStruct{Name: "TestValue"}

	b := new(path, func(w http.ResponseWriter, r *http.Request, o interface{}) {
		w.Header().Set("MyCustomHeaderName", "MyCustomHeaderValue")
	}, func(w http.ResponseWriter, r *http.Request, o interface{}) {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Content-Type", "text/html")

		var str string
		if m, ok := o.(*someTestStruct); ok {
			str = m.Name
		}

		w.Write([]byte(`Hello World! (` + str + `)`))
	}, obj)
	return http.HandlerFunc(b.handler)
}

func request(t *testing.T, file string) *httptest.ResponseRecorder {
	request, err := http.NewRequest("GET", file, nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handle().ServeHTTP(recorder, request)
	return recorder
}

func check_status_code(t *testing.T, r *httptest.ResponseRecorder) {
	if s := r.Code; s != http.StatusOK {
		t.Fatalf("handler return wrong status code: got (%v) want (%v)", s, http.StatusOK)
	}
}

func check_content_type(t *testing.T, r *httptest.ResponseRecorder, ctype string) {
	if c := r.Header().Get("Content-Type"); c != ctype {
		t.Fatalf("content type header not match: got (%v) want (%v)", c, ctype)
	}
}

func check_response_body(t *testing.T, r *httptest.ResponseRecorder, body []byte) {
	if r.Body.String() != string(body) {
		t.Fatalf("bad body response, not match")
	}
}

func check_resource(t *testing.T, file string, ctype string, body []byte) {
	r := request(t, "/"+path+"/"+file)
	check_status_code(t, r)
	check_content_type(t, r, ctype)
	check_response_body(t, r, body)
}

func TestBootstrapCss(t *testing.T) {
	check_resource(t, "bootstrap.css", "text/css", resource_bootstrap_css)
}

func TestBootstrapJs(t *testing.T) {
	check_resource(t, "bootstrap.js", "application/javascript; charset=utf-8", resource_bootstrap_js)
}

func TestJqueryJs(t *testing.T) {
	check_resource(t, "jquery.js", "application/javascript; charset=utf-8", resource_jquery_js)
}

func TestPopperJs(t *testing.T) {
	check_resource(t, "popper.js", "application/javascript; charset=utf-8", resource_popper_js)
}

func TestBeforeCallBack(t *testing.T) {
	r := request(t, "/")
	if c := r.Header().Get("MyCustomHeaderName"); c != "MyCustomHeaderValue" {
		t.Fatalf("content type header not match: got (%v) want (%v)", c, "MyCustomHeaderValue")
	}
}

func TestAfterCallBack(t *testing.T) {
	r := request(t, "/")
	if c := r.Header().Get("Cache-Control"); c != "no-cache, no-store, must-revalidate" {
		t.Fatalf("content type header not match: got (%v) want (%v)", c, "no-cache, no-store, must-revalidate")
	}
	if c := r.Header().Get("Content-Type"); c != "text/html" {
		t.Fatalf("content type header not match: got (%v) want (%v)", c, "text/html")
	}
	if r.Body.String() != "Hello World! (TestValue)" {
		t.Fatalf("bad body response, not match")
	}
}
