package kanbaru

import (
	"net/http"
	"strings"
)

type HttpReq struct {
	*http.Request
	CheckUrl string
}

func (r HttpReq) GetVar(key string) (string, bool) {
	url := strings.TrimSuffix(r.URL.Path, "/")
	checkUrlSplitted := strings.Split(r.CheckUrl[1:], "/")
	currentUrlSplitted := strings.Split(url[1:], "/")
	if len(checkUrlSplitted) == len(currentUrlSplitted) {
		for i := 0; i < len(checkUrlSplitted); i++ {
			if strings.HasPrefix(checkUrlSplitted[i], ":"+key) {
				return currentUrlSplitted[i], true
			}
		}
	}
	return "", false
}
