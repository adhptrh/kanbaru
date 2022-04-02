package cli

import (
	"fmt"
	"os"
)

func Main() {
	fmt.Println("ok")
	if os.Args[2] == "add" {
		if os.Args[3] == "basic-path" {

			os.Mkdir("path/"+os.Args[4], os.ModeAppend)
			file, _ := os.Create("path/" + os.Args[4] + "/a.go")
			file.Write([]byte(`package ` + os.Args[4] + `

import (
"kanbaru/kanbaru"
"net/http"
)
			
func Main(w http.ResponseWriter, r kanbaru.HttpReq) {
w.Write([]byte("Basic response"))
}`))
			file.Close()
		}
	}
}
