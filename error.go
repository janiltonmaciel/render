package render

import (
	"net/http"

	"github.com/janiltonmaciel/errors"
)

func WriteError(w http.ResponseWriter, err error) error {
	code := http.StatusInternalServerError
	message := http.StatusText(code)

	if err, ok := err.(*errors.HTTP); ok {
		return WriteJSON(w, err, err.Code)
	}

	return WriteJSON(w, errors.NewHttpError(code, message), code)
}
