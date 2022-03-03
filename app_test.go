package passage_test

import (
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
	}

	magicLink, err := psg.CreateMagicLink(createMagicLinkBody)
	require.Nil(t, err)
	assert.Equal(t, createMagicLinkBody.Email, magicLink.Identifier)
}
