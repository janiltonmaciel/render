package render

import "net/http"

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
	CacheControl   = "Cache-Control"
)

func WriteNoContent(w http.ResponseWriter) error {
	w.WriteHeader(http.StatusNoContent)
	return nil
}

func getStatusCode(code ...int) int {
	if len(code) > 0 {
		return code[0]
	}
	return http.StatusOK
}

func getMaxAge(maxAge ...int) int {
	if len(maxAge) > 0 {
		return maxAge[0]
	}
	return 0
}
