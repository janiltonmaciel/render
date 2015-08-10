package render

import (
	"encoding/xml"
	"net/http"
)

func WriteXML(w http.ResponseWriter, v interface{}, code ...int) error {
	w.Header().Set(ContentType, ApplicationXMLCharsetUTF8)
	w.WriteHeader(getStatusCode(code...))
	w.Write([]byte(xml.Header))

	return xml.NewEncoder(w).Encode(v)
}
