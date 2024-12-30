package passage_test

import (
	"sync"
	"testing"

	"github.com/passageidentity/passage-go"
	"github.com/stretchr/testify/require"
)

// should be run with the -race flag, i.e. `go test -race -run TestAppJWKSCacheWriteConcurrency`
func TestAppJWKSCacheWriteConcurrency(t *testing.T) {
	goRoutineCount := 2

	var wg sync.WaitGroup
	wg.Add(goRoutineCount)

	for i := 0; i < goRoutineCount; i++ {
		go func() {
			defer wg.Done()

			_, err := passage.New(PassageAppID, &passage.Config{
				APIKey: PassageApiKey, // An API_KEY environment variable is required for testing.
			})
			require.Nil(t, err)
		}()
	}

	wg.Wait()
}
