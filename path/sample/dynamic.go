package sample

import (
	"fmt"
	"kanbaru/kanbaru"
	"net/http"
)

func Dynamic(w http.ResponseWriter, r kanbaru.HttpReq) {
	res, _ := r.GetVar("inidata")
	fmt.Fprint(w, res)
}
