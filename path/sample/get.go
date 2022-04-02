package sample

import (
	"fmt"
	"kanbaru/kanbaru"
	"net/http"
)

func Get(w http.ResponseWriter, r kanbaru.HttpReq) {
	fmt.Fprint(w, "Get")
}
