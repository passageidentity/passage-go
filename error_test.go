package passage_test

import (
	"net/http"
	"testing"

	"github.com/passageidentity/passage-go"
	"github.com/stretchr/testify/assert"
)

func TestErrorWithAllFields(t *testing.T) {
	Error := passage.Error{
		Message:    "some message",
		StatusCode: http.StatusBadRequest,
		StatusText: http.StatusText(http.StatusBadRequest),
		ErrorText:  "some error",
	}
	errorString := Error.Error()
	assert.Equal(t, "Passage Error - message: some message, status_code: 400, status_text: Bad Request, error: some error", errorString)

}

func TestErrorWithOnlyMessage(t *testing.T) {
	Error := passage.Error{
		Message: "some message",
	}
	errorString := Error.Error()
	assert.Equal(t, "Passage Error - message: some message", errorString)

}
