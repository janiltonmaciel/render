package render

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

const (

	//-------------
	// Charset
	//-------------
	CharsetUTF8 = "charset=utf-8"

	//-------------
	// Media types
	//-------------
	ApplicationJSON                  = "application/json"
	ApplicationJSONCharsetUTF8       = ApplicationJSON + "; " + CharsetUTF8
	ApplicationJavaScript            = "application/javascript"
	ApplicationJavaScriptCharsetUTF8 = ApplicationJavaScript + "; " + CharsetUTF8
	ApplicationXML                   = "application/xml"
	ApplicationXMLCharsetUTF8        = ApplicationXML + "; " + CharsetUTF8
	ApplicationForm                  = "application/x-www-form-urlencoded"
	TextHTML                         = "text/html"
	TextHTMLCharsetUTF8              = TextHTML + "; " + CharsetUTF8
	TextPlain                        = "text/plain"
	TextPlainCharsetUTF8             = TextPlain + "; " + CharsetUTF8

	//---------
	// Headers
	//---------
	Accept         = "Accept"
	AcceptEncoding = "Accept-Encoding"
	Authorization  = "Authorization"
	ContentType    = "Content-Type"
)

func WriteHTML(w http.ResponseWriter, code int, html string) error {
	return WriteHTMLf(w, code, "%s", html)
}

func WriteHTMLf(w http.ResponseWriter, code int, format string, a ...interface{}) (err error) {
	w.Header().Set(ContentType, TextHTMLCharsetUTF8)
	w.WriteHeader(code)
	_, err = fmt.Fprintf(w, format, a...)
	return err
}

func WriteString(w http.ResponseWriter, code int, value string) error {
	return WriteStringf(w, code, "%s", value)
}

func WriteStringf(w http.ResponseWriter, code int, format string, a ...interface{}) (err error) {
	w.Header().Set(ContentType, TextPlain)
	w.WriteHeader(code)
	_, err = fmt.Fprintf(w, format, a...)
	return err
}

func WriteJSON(w http.ResponseWriter, code int, v interface{}) error {
	result, err := json.Marshal(v)
	if err != nil {
		return err
	}

	w.Header().Set(ContentType, ApplicationJSONCharsetUTF8)
	w.WriteHeader(code)
	w.Write(result)

	return nil
}

func WriteJSONP(w http.ResponseWriter, code int, callback string, v interface{}) (err error) {
	w.Header().Set(ContentType, ApplicationJavaScriptCharsetUTF8)
	w.WriteHeader(code)
	w.Write([]byte(callback + "("))
	if err = json.NewEncoder(w).Encode(v); err == nil {
		w.Write([]byte(");"))
	}
	return err
}

func WriteXML(w http.ResponseWriter, code int, v interface{}) error {
	w.Header().Set(ContentType, ApplicationXMLCharsetUTF8)
	w.WriteHeader(code)
	w.Write([]byte(xml.Header))

	return xml.NewEncoder(w).Encode(v)
}
