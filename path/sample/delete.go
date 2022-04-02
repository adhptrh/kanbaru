package sample

import (
	"fmt"
	"kanbaru/kanbaru"
	"net/http"
)

func Delete(w http.ResponseWriter, r kanbaru.HttpReq) {
	fmt.Fprint(w, "Delete")
}
