package sample

import (
	"fmt"
	"kanbaru/kanbaru"
	"net/http"
)

func Patch(w http.ResponseWriter, r kanbaru.HttpReq) {
	fmt.Fprint(w, "Patch")
}
