package util

import (
	"encoding/json"
	"github.com/tonquangtu/http_server/internal/dto"
	"net/http"
)

func RenderJson(w http.ResponseWriter, res dto.ApiResponse) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	for key, value := range res.Headers {
		w.Header().Set(key, value)
	}

	w.WriteHeader(res.StatusCode)
	if (res.Data == nil && res.Message == "" && res.Errors == nil) || res.StatusCode == http.StatusNoContent {
		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Ok(data interface{}) dto.ApiResponse {
	return dto.ApiResponse{StatusCode: http.StatusOK, Data: data}
}

func BadRequest(err error) dto.ApiResponse {
	return dto.ApiResponse{
		StatusCode: http.StatusBadRequest,
		Errors: map[string][]string{
			"Message": {err.Error()},
		},
	}
}

func UnprocessableEntity(err error) dto.ApiResponse {
	return dto.ApiResponse{
		StatusCode: http.StatusUnprocessableEntity,
		Errors: map[string][]string{
			"Message": {err.Error()},
		},
	}
}
