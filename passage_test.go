package passage_test

import (
	"crypto/rand"
	"fmt"
	"os"
	"strings"
	"testing"

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
