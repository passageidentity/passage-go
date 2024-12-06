package passage

import (
	"encoding/json"
	"fmt"
	"strings"
)

type PassageError struct {
	Message    string
	ErrorCode  string
	StatusCode int
}

func (e PassageError) Error() string {
	var sb strings.Builder
	sb.WriteString("PassageError - ")

	if e.Message != "" {
		sb.WriteString(fmt.Sprintf("message: %s, ", e.Message))
	}

	if e.ErrorCode != "" {
		sb.WriteString(fmt.Sprintf("errorCode: %v, ", e.ErrorCode))
	}

	if e.StatusCode != 0 {
		sb.WriteString(fmt.Sprintf("statusCode: %v, ", e.StatusCode))
	}

	return strings.TrimSuffix(sb.String(), ", ")
}

func errorFromResponse(body []byte, statusCode int) error {
	var errorBody struct {
		Code  string `json:"code"`
		Error string `json:"error"`
	}
	if err := json.Unmarshal(body, &errorBody); err != nil {
		return err
	}

	return PassageError{
		Message:    errorBody.Error,
		ErrorCode:  errorBody.Code,
		StatusCode: statusCode,
	}
}
