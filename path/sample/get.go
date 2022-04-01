package sample

import (
	"fmt"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Get")
}
