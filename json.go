package render

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, v interface{}, code int, maxAge ...int) error {
	result, err := json.Marshal(v)
	if err != nil {
		return err
	}

	head := NewHead(ApplicationJSONCharsetUTF8, code, maxAge...)
	head.Write(w)
	w.Write(result)

	return nil
}

func WriteJSONString(w http.ResponseWriter, v string, code int, maxAge ...int) (err error) {
	head := NewHead(ApplicationJSONCharsetUTF8, code, maxAge...)
	head.Write(w)
	_, err = fmt.Fprintf(w, "%s", v)
	return err
}

func WriteJSONP(w http.ResponseWriter, callback string, v interface{}, code ...int) (err error) {
	w.Header().Set(ContentType, ApplicationJavaScriptCharsetUTF8)
	w.WriteHeader(getStatusCode(code...))
	w.Write([]byte(callback + "("))
	if err = json.NewEncoder(w).Encode(v); err == nil {
		w.Write([]byte(");"))
	}
	return err
}
