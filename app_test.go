package passage_test

import (
	"sync"
	"testing"

	"github.com/passageidentity/passage-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateMagicLink(t *testing.T) {
	psg, err := passage.New(PassageAppID, &passage.Config{
		APIKey: PassageApiKey, // An API_KEY environment variable is required for testing.
	})
	require.Nil(t, err)

	createMagicLinkBody := passage.CreateMagicLinkBody{
		Email:   "chris@passage.id",
		Channel: passage.EmailChannel,
		TTL:     12,
		Type:    passage.LoginType,
	}

	magicLink, err := psg.CreateMagicLink(createMagicLinkBody)
	require.Nil(t, err)
	assert.Equal(t, createMagicLinkBody.Email, magicLink.Identifier)
	assert.Equal(t, createMagicLinkBody.TTL, magicLink.TTL)
}

func TestGetApp(t *testing.T) {
	psg, err := passage.New(PassageAppID, &passage.Config{
		APIKey: PassageApiKey, // An API_KEY environment variable is required for testing.
	})
	require.Nil(t, err)

	appInfo, err := psg.GetApp()
	assert.Nil(t, err)
	assert.Equal(t, PassageAppID, appInfo.ID)

}

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
