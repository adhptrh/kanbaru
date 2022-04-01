package sample

import (
	"fmt"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Delete")
}
