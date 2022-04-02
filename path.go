package main

import (
	"kanbaru/config/path"
	"kanbaru/path/sample"
)

func Path() {
	path.Get("/sample", sample.Main)
	path.Get("/sample/:inidata/data", sample.Dynamic)
	path.All("/allmethod", sample.Main)
}
