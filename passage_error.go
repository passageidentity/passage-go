package passage

import (
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
