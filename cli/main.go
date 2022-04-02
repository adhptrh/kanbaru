package cli

import (
	"os"
	"strings"

	"github.com/fatih/color"
)

func Main() {
	if os.Args[2] == "add" {
		if os.Args[3] == "basic-path" {
			color.HiBlue("Creating path " + os.Args[4] + "...")
			os.Mkdir("path/"+os.Args[4], os.ModeAppend)
			file, _ := os.Create("path/" + os.Args[4] + "/a.go")
			template, _ := os.ReadFile("cli/template/basic-path")
			templateStr := string(template)
			templateStr = strings.ReplaceAll(templateStr, "__package_name__", os.Args[4])
			file.Write([]byte(templateStr))
			file.Close()
			color.HiGreen("Successfully created path " + os.Args[4] + "!")
		}
	}
}
