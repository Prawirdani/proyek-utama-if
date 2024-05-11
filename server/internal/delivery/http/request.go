package http

import (
	"encoding/json"
	"net/http"

	"github.com/prawirdani/golang-restapi/pkg/utils"
)

func BindAndValidate[T any](r *http.Request) (data T, err error) {
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&data); err != nil {
		return data, err
	}

	if err := utils.Validate.Struct(data); err != nil {
		return data, err
	}

	return data, nil
}
