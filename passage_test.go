package passage_test

import (
	"crypto/rand"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/joho/godotenv"
	"github.com/passageidentity/passage-go"
)

var (
	PassageAppID     string
	PassageApiKey    string
	PassageUserID    string
	PassageAuthToken string
	RandomEmail      = generateRandomEmail(14)
	CreatedUser      passage.User
)

func generateRandomEmail(prefixLength int) string {
	n := prefixLength
	randomChars := make([]byte, n)
	if _, err := rand.Read(randomChars); err != nil {
		panic(err)
	}
	email := fmt.Sprintf("%X@email.com", randomChars)
	return strings.ToLower(email)
}

func TestMain(m *testing.M) {
	_ = godotenv.Load(".env")

	PassageAppID = os.Getenv("PASSAGE_APP_ID")
	PassageApiKey = os.Getenv("PASSAGE_API_KEY")
	PassageUserID = os.Getenv("PASSAGE_USER_ID")
	PassageAuthToken = os.Getenv("PASSAGE_AUTH_TOKEN")

	exitVal := m.Run()
	os.Exit(exitVal)
}

func userNotFoundAsserts(t *testing.T, err error, message string) {
	splitError := strings.Split(err.Error(), ", ")
		assert.Len(t, splitError, 5)
		assert.Equal(t, message, splitError[0])
		assert.Equal(t, "status_code: 404", splitError[1])
		assert.Equal(t, "status_text: 404 Not Found", splitError[2])
		assert.Equal(t, "error_code: user_not_found", splitError[3])
		assert.Equal(t, "error: User not found", splitError[4])
}

func unauthorizedAsserts(t *testing.T, err error) {
	splitError := strings.Split(err.Error(), ", ")
		assert.Len(t, splitError, 3)
		assert.Equal(t, "Passage Error - message: failed to get Passage User" , splitError[0])
		assert.Equal(t, "status_code: 401", splitError[1])
		assert.Equal(t, "status_text: 401 Unauthorized", splitError[2])
}
