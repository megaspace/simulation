package stars

import (
	"math/rand"

	"github.com/megaspace/simulation/engine/core"
)

var (
	// Classifications is a collection of main sequence star types
	Classifications = map[string]Classification{
		"O": {
			name:        "O",
			color:       core.NewColor(0.604, 0.686, 1.0),
			temperature: core.NewSpan(33000, 40000),
			mass:        core.NewSpan(16, 17),
			radius:      core.NewSpan(6.6, 8),
		},
		"B": {
			name:        "B",
			color:       core.NewColor(0.792, 0.843, 1.0),
			temperature: core.NewSpan(10000, 33000),
			mass:        core.NewSpan(2.1, 16),
			radius:      core.NewSpan(1.8, 6.6),
		},
		"A": {
			name:        "A",
			color:       core.NewColor(0.973, 0.969, 1.0),
			temperature: core.NewSpan(7500, 10000),
			mass:        core.NewSpan(1.4, 2.1),
			radius:      core.NewSpan(1.4, 1.8),
		},
		"F": {
			name:        "F",
			color:       core.NewColor(0.988, 1.0, 0.827),
			temperature: core.NewSpan(6000, 7500),
			mass:        core.NewSpan(1.04, 1.4),
			radius:      core.NewSpan(1.15, 1.4),
		},
		"G": {
			name:        "G",
			color:       core.NewColor(1.0, 0.949, 0.631),
			temperature: core.NewSpan(5200, 6000),
			mass:        core.NewSpan(0.8, 1.04),
			radius:      core.NewSpan(0.96, 1.15),
		},
		"K": {
			name:        "K",
			color:       core.NewColor(1.0, 0.639, 0.318),
			temperature: core.NewSpan(3700, 5200),
			mass:        core.NewSpan(0.45, 0.8),
			radius:      core.NewSpan(0.7, 0.96),
		},
		"M": {
			name:        "M",
			color:       core.NewColor(1.0, 0.38, 0.318),
			temperature: core.NewSpan(2000, 3700),
			mass:        core.NewSpan(0.3, 0.45),
			radius:      core.NewSpan(0.5, 0.7),
		},
	}
)

// Classification is a description of the range of stars that exist in a certain main sequence classification
type Classification struct {
	name        string
	color       core.Color
	temperature core.Span
	mass        core.Span
	radius      core.Span
}

// GetRandomClassification returns a classification based on commonality of main sequence stars
func GetRandomClassification() Classification {
	selection := rand.Float64() * 100

	if selection < 0.00003 {
		return Classifications["O"]
	}

	if selection < 0.13 {
		return Classifications["B"]
	}

	if selection < 0.73 {
		return Classifications["A"]
	}

	if selection < 3.73 {
		return Classifications["F"]
	}

	if selection < 11.33 {
		return Classifications["G"]
	}

	if selection < 23.43 {
		return Classifications["K"]
	}

	return Classifications["M"]
}
