package passage

import (
	"fmt"
	"strings"
)

type Error struct {
	Message    string
	StatusCode int
	ErrorCode  string
	// Deprecated
	StatusText string
	// Deprecated
	ErrorText string
}

type HTTPError struct {
	ErrorText string `json:"error,omitempty"`
}

func (e Error) Error() string {
	var ps strings.Builder
	ps.WriteString("Passage Error - ")

	if e.Message != "" {
		fmt.Fprintf(&ps, "message: %s, ", e.Message)
	}
	if e.StatusCode != 0 {
		fmt.Fprintf(&ps, "status_code: %v, ", e.StatusCode)
	}
	if e.StatusText != "" {
		fmt.Fprintf(&ps, "status_text: %s, ", e.StatusText)
	}
	if e.ErrorCode != "" {
		fmt.Fprintf(&ps, "error_code: %s, ", e.ErrorCode)
	}
	if e.ErrorText != "" {
		fmt.Fprintf(&ps, "error: %s, ", e.ErrorText)
	}

	return strings.TrimSuffix(ps.String(), ", ")
}
