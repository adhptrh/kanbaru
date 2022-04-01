package main

import (
	"flag"
	"fmt"
	"kanbaru/path/index"
	"net/http"
	"strings"
)

func main() {
	port := flag.Int("port", 1337, "Port for the server")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		for i := 0; i < len(Paths); i++ {
			if strings.EqualFold("/", r.URL.Path) {
				index.Main(w, r)
				return
			}
			if strings.EqualFold(Paths[i]["path"].(string), r.URL.Path) {
				Paths[i]["target"].(func(http.ResponseWriter, *http.Request))(w, r)
				return
			}
		}

		fmt.Fprintf(w, "Not found")
	})

	fmt.Println("Listening on port", *port)
	http.ListenAndServe(":"+fmt.Sprint(*port), nil)
}
