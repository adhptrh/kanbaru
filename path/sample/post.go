package sample

import (
	"fmt"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post")
}
