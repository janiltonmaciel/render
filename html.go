package render

import (
	"fmt"
	"net/http"
)

func WriteHTML(w http.ResponseWriter, html string, code int, maxAge ...int) error {
	return WriteHTMLf(w, "%s", code, getMaxAge(maxAge...), html)
}

func WriteHTMLf(w http.ResponseWriter, format string, code int, maxAge int, a ...interface{}) (err error) {
	head := NewHead(TextHTMLCharsetUTF8, code, maxAge)
	head.Write(w)
	_, err = fmt.Fprintf(w, format, a...)
	return err
}
