package cli

import (
	"os"

	"github.com/fatih/color"
)

func Main() {
	if os.Args[2] == "add" {
		if os.Args[3] == "basic-path" {
			color.HiBlue("Creating path " + os.Args[4] + "...")
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
			color.HiGreen("Successfully created path " + os.Args[4] + "!")
		}
	}
}
