package main

import (
	"kanbaru/kanbaru/path"
	"kanbaru/path/sample"
)

func Path() {
	path.Get("/sample", sample.Main)
	path.Get("/sample/:inidata/data", sample.Dynamic)
	path.All("/allmethod", sample.Main)
}
