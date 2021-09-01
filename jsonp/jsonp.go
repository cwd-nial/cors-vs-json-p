package jsonp

import "embed"

//go:embed *.html *.ico *.js
var static embed.FS

func GetStaticFiles() embed.FS {
	return static
}
