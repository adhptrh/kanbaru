package sample

import (
	"net/http"
)

func Main(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		Get(w, r)
	case "POST":
		Post(w, r)
	case "PATCH":
		Patch(w, r)
	case "DELETE":
		Delete(w, r)
	}
}
