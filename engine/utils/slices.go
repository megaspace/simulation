package utils

import (
	"math/rand"
	"sort"
)

type Shufflable interface {
	Len() int
	Swap(i, j int)
}

func ShuffleStringSlice(a []string) {
	shuffle(sort.StringSlice(a))
}

func shuffle(data Shufflable) {
	n := data.Len()
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data.Swap(i, j)
	}
}
