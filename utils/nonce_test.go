package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateNonce(t *testing.T) {
	s1, err := GenerateNonce()
	require.NoError(t, err)
	t.Log(s1)

	s2, err := GenerateNonce()
	require.NoError(t, err)
	t.Log(s2)

	assert.NotEqual(t, s1, s2)
}
