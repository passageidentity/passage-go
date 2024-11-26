package passage

import (
	"fmt"
	"strings"
)

type PassageError struct {
	Message    string
	StatusCode int
	ErrorCode  string
}

func (e PassageError) Error() string {
	var ps strings.Builder
	ps.WriteString("PassageError - ")

	if e.Message != "" {
		fmt.Fprintf(&ps, "message: %s, ", e.Message)
	}
	if e.StatusCode != 0 {
		fmt.Fprintf(&ps, "status_code: %v, ", e.StatusCode)
	}
	if e.ErrorCode != "" {
		fmt.Fprintf(&ps, "error_code: %s, ", e.ErrorCode)
	}

	return strings.TrimSuffix(ps.String(), ", ")
}
