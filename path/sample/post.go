package sample

import (
	"fmt"
	"kanbaru/kanbaru"
	"net/http"
)

func Post(w http.ResponseWriter, r kanbaru.HttpReq) {
	fmt.Fprint(w, "Post")
}
