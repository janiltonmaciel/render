package render

import (
	"fmt"
	"net/http"
)

func WriteString(w http.ResponseWriter, value string, code ...int) error {
	return WriteStringf(w, "%s", getStatusCode(code...), value)
}

func WriteStringf(w http.ResponseWriter, format string, code int, a ...interface{}) (err error) {
	w.Header().Set(ContentType, TextPlain)
	w.WriteHeader(code)
	_, err = fmt.Fprintf(w, format, a...)
	return err
}
