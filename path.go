package main

import (
	"kanbaru/path/sample"
)

var Paths []map[string]interface{} = []map[string]interface{}{
	{
		"path":   "/sample",
		"target": sample.Main,
	},
	{
		"path":   "/sample/:inidata/data",
		"target": sample.Dynamic,
	},
}
