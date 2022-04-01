package sample

import (
	"fmt"
	"net/http"
)

func Patch(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Patch")
}
