package jsonp

import "embed"

//go:embed *.html *.js
var static embed.FS

func GetStaticFiles() embed.FS {
	return static
}
