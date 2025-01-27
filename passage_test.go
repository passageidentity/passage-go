package passage_test

import (
	"os"
	"sync"
	"testing"

	"github.com/joho/godotenv"
	"github.com/passageidentity/passage-go/v2"
	"github.com/stretchr/testify/require"
)

// should be run with the -race flag, i.e. `go test -race -run TestAppJWKSCacheWriteConcurrency`
func TestAppJWKSCacheWriteConcurrency(t *testing.T) {
	_ = godotenv.Load(".env")

	appID := os.Getenv("PASSAGE_APP_ID")
	apiKey := os.Getenv("PASSAGE_API_KEY")

	goRoutineCount := 2

	var wg sync.WaitGroup
	wg.Add(goRoutineCount)

	for i := 0; i < goRoutineCount; i++ {
		go func() {
			defer wg.Done()

			_, err := passage.New(appID, apiKey)
			require.Nil(t, err)
		}()
	}

	wg.Wait()
}
