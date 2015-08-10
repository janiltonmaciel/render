package render

import (
	"net/http"

	"gitlab.globoi.com/bastian/falkor/errors"
)

func WriteError(w http.ResponseWriter, err error) error {
	code := http.StatusInternalServerError
	message := http.StatusText(code)

	if err, ok := err.(*errors.HTTP); ok {
		return WriteJSON(w, err, err.Code)
	}

	return WriteJSON(w, errors.NewHttpError(code, message), code)
}
