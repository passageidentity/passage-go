package passage_test

import (
	"sync"
	"testing"

	"github.com/passageidentity/passage-go/v2"
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

			// a network call is made upon initialization to retrieve the JWKs from a real source
			_, err := passage.New("passage", "some-api-key")
			require.Nil(t, err)
		}()
	}

	wg.Wait()
}
