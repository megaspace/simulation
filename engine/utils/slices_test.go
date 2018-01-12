package utils_test

import (
	"math/rand"
	"testing"

	"github.com/megaspace/simulation/engine/utils"
	"github.com/stretchr/testify/assert"
)

func TestShuffleStringSlice(t *testing.T) {
	rand.Seed(1)
	s := []string{"a", "b", "c", "d", "e", "f", "g"}

	utils.ShuffleStringSlice(s)

	assert.NotEqual(t, s, []string{"a", "b", "c", "d", "e", "f", "g"})
}
