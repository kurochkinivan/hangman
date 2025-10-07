package wordslist

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadWordsList_HappyPath(t *testing.T) {
	yamlContents := `
animals:
  easy:
    - word: cat
      hint: domestic pet
  medium:
    - word: lion
  hard:
`

	r := strings.NewReader(yamlContents)
	wl, err := loadWordsList(r)
	require.NoError(t, err)

	animals := wl.Animals

	assert.Len(t, animals.Easy, 1)
	assert.Len(t, animals.Medium, 1)
	assert.Empty(t, animals.Hard)

	assert.Equal(t, "cat", animals.Easy[0].Value)
	assert.Equal(t, "domestic pet", animals.Easy[0].Hint)

	assert.Equal(t, "lion", animals.Medium[0].Value)
	assert.Equal(t, "", animals.Medium[0].Hint)
}
