package index

import (
	"net/http"
)

func Main(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is index"))
}
