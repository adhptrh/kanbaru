package sample

import (
	"fmt"
	"kanbaru/kanbaru"
	"net/http"
)

func Main(w http.ResponseWriter, r kanbaru.HttpReq) {
	fmt.Fprintf(w, "Method : %v", r.Method)
}
