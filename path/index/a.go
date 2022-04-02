package index

import (
	"kanbaru/kanbaru"
	"net/http"
)

func Main(w http.ResponseWriter, r kanbaru.HttpReq) {
	w.Write([]byte("This is index"))
}
