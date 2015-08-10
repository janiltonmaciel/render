package render

import (
	"fmt"
	"net/http"
)

func WriteHTML(w http.ResponseWriter, html string, code ...int) error {
	return WriteHTMLf(w, "%s", getStatusCode(code...), html)
}

func WriteHTMLf(w http.ResponseWriter, format string, code int, a ...interface{}) (err error) {
	w.Header().Set(ContentType, TextHTMLCharsetUTF8)
	w.WriteHeader(code)
	_, err = fmt.Fprintf(w, format, a...)
	return err
}
