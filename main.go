package main

import (
	"flag"
	"fmt"
	"kanbaru/config/path"
	"kanbaru/kanbaru"
	"kanbaru/kanbaru/cli"
	"kanbaru/path/index"
	"net/http"
	"os"
	"strings"

	"github.com/fatih/color"
)

var white = color.New(color.FgWhite).SprintFunc()
var hiblue = color.New(color.FgHiBlue).SprintFunc()
var hiyellow = color.New(color.FgHiYellow).SprintFunc()

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "cli" {
			cli.Main()
			return
		}
	}

	port := flag.Int("port", 1337, "Port for the server")
	flag.Parse()

	Path()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Printf("%s%s%s%s%s", hiblue("["), white(fmt.Sprintf("%v", r.Method)), hiblue("]"), white(" - "), hiyellow(fmt.Sprintf("%v", r.URL.Path)))
		fmt.Println()
		for i := 0; i < len(path.Paths); i++ {
			rr := kanbaru.HttpReq{
				Request:  r,
				CheckUrl: path.Paths[i]["path"].(string),
			}
			if strings.EqualFold("/", r.URL.Path) {
				index.Main(w, rr)
				return
			}
			res := checkMatch(path.Paths[i], r.URL.Path)
			if res != nil {
				if res["method"] == "ALL" {
					res["target"].(func(w http.ResponseWriter, r kanbaru.HttpReq))(w, rr)
					return
				}
				if res["method"] == r.Method {
					res["target"].(func(w http.ResponseWriter, r kanbaru.HttpReq))(w, rr)
					return
				}
			}
		}
		w.WriteHeader(404)
		fmt.Fprintf(w, "Not found")
	})

	fmt.Printf("%s%s%s", hiblue("Listening on port "), hiyellow(fmt.Sprintf("%v", *port)), hiblue("..."))
	fmt.Println()
	http.ListenAndServe(":"+fmt.Sprint(*port), nil)
}

func checkMatch(check map[string]interface{}, url string) map[string]interface{} {
	url = strings.TrimSuffix(url, "/")
	checkUrlSplitted := strings.Split(check["path"].(string)[1:], "/")
	currentUrlSplitted := strings.Split(url[1:], "/")
	if len(checkUrlSplitted) == len(currentUrlSplitted) {
		checkUrl := ""
		currentUrl := ""
		for ii := 0; ii < len(checkUrlSplitted); ii++ {
			if strings.HasPrefix(checkUrlSplitted[ii], ":") {
				checkUrl += "/" + checkUrlSplitted[ii][1:]
				currentUrl += "/" + checkUrlSplitted[ii][1:]
				continue
			}
			checkUrl += "/" + checkUrlSplitted[ii]
			currentUrl += "/" + currentUrlSplitted[ii]
		}

		if strings.EqualFold(checkUrl, currentUrl) {
			return check
		}
	}
	return nil
}
