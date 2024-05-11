package httputil

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func ParamInt(r *http.Request, key string) (int, error) {
	paramValue := chi.URLParam(r, key)
	val, err := strconv.Atoi(paramValue)
	if err != nil {
		return 0, err
	}
	return val, nil
}
