package response

import (
	"encoding/json"
	"healtEvaluationApp/internal/healthEvaluationApp/common"
	"healtEvaluationApp/internal/utility"
	"net/http"
)

type ErrorResponse struct {
	StatusCode   int    `json:"statusCode"`
	ErrorMessage string `json:"errorMessage"`
}

func Success(w http.ResponseWriter, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		utility.LogPrometheusErrorCustom(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set(common.ContentTypeKey, common.ContentTypeValue)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func Error(w http.ResponseWriter, payload interface{}, statusCode int) {
	response, err := json.Marshal(payload)
	if err != nil {
		utility.LogPrometheusErrorCustom(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set(common.ContentTypeKey, common.ContentTypeValue)
	w.WriteHeader(statusCode)
	w.Write(response)
}

// BuildErrorMessage Error message builder
func BuildErrorMessage(status int, msg string) ErrorResponse {
	errResponse := ErrorResponse{}
	errResponse.StatusCode = status
	errResponse.ErrorMessage = msg

	return errResponse
}
