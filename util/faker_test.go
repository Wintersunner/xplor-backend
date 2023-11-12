package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRandomString(t *testing.T) {
	randomString := RandomString(6)
	require.NotEmpty(t, randomString)
	require.Equal(t, len(randomString), 6)
}

func TestRandomInt(t *testing.T) {
	randomInt := RandomInt(10, 50)
	require.NotEmpty(t, randomInt)
	require.GreaterOrEqual(t, randomInt, int64(10))
	require.LessOrEqual(t, randomInt, int64(50))
}

func TestRandomMessage(t *testing.T) {
	randomMessage := RandomMessage()
	require.NotEmpty(t, randomMessage)
	require.GreaterOrEqual(t, len(randomMessage), 10)
	require.LessOrEqual(t, len(randomMessage), 20)
}
