package passage_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/passageidentity/passage-go"
	"github.com/stretchr/testify/assert"
)

func TestPassageErrorWithAllFields(t *testing.T) {
	err := passage.PassageError{
		Message:    "some message",
		ErrorCode:  "some_error_code",
		StatusCode: http.StatusBadRequest,
	}
	assert.Equal(t, fmt.Sprintf("PassageError - message: %s, errorCode: %s, statusCode: %d", err.Message, err.ErrorCode, err.StatusCode), err.Error())
}

func TestPassageErrorWithOnlyMessage(t *testing.T) {
	err := passage.PassageError{
		Message: "some message",
	}
	assert.Equal(t, fmt.Sprintf("PassageError - message: %s", err.Message), err.Error())
}
