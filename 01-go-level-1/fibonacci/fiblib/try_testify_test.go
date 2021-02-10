package fiblib

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSmth(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)
	assert.Equal(uint64(2), FibonacciGoldenRatio(3), "they should be equal")
	require.NotEqual(uint64(3), FibonacciGoldenRatio(3), "they should not be equal")

}
