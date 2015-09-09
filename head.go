package render

import (
	"fmt"
	"net/http"
)

type Head struct {
	StatusCode  int
	MaxAge      int
	contentType string
}

func (h Head) Write(w http.ResponseWriter) {
	w.Header().Set(ContentType, h.contentType)
	if h.MaxAge > 0 {
		w.Header().Set(CacheControl, fmt.Sprintf("max-age=%d", h.MaxAge))
	} else {
		w.Header().Set(CacheControl, "no-cache")
	}
	w.WriteHeader(h.StatusCode)
}

func NewHead(contentType string, statusCode int, maxAge ...int) Head {
	max := getMaxAge(maxAge...)
	return Head{
		MaxAge:      max,
		StatusCode:  statusCode,
		contentType: contentType,
	}
}
