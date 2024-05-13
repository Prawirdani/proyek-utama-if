package http

import (
	"net/http"

	"github.com/prawirdani/golang-restapi/pkg/httputil"
)

type resOption func(*httputil.Response)

func data(v any) resOption {
	return func(r *httputil.Response) {
		r.Data = v
	}
}
func message(msg string) resOption {
	return func(r *httputil.Response) {
		r.Message = &msg
	}
}

func status(status int) resOption {
	return func(r *httputil.Response) {
		r.Status = status
	}
}

func response(w http.ResponseWriter, opts ...func(*httputil.Response)) error {
	res := &httputil.Response{
		Status: 200, // Default
	}

	for _, opt := range opts {
		opt(res)
	}

	return httputil.WriteJSON(w, res.Status, res)
}

// ToResponseList converts data to response list
// T is the type of data, U is the type of response
func ToResponseList[T, U any](data []T, toResponse func(T) U) []U {
	var result []U
	for i := 0; i < len(data); i++ {
		result = append(result, toResponse(data[i]))
	}
	return result
}
