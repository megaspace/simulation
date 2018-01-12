package stars

import (
	"github.com/megaspace/simulation/engine/core"
	"github.com/megaspace/simulation/engine/utils"
)

// Service is the service layer for all things about stars
type Service struct {
	stars map[StarID]*Star
}

// NewService returns a new service to manage all the stars.
// Only one of these is needed per game.
func NewService() *Service {
	starService := new(Service)
	starService.stars = make(map[StarID]*Star)
	return starService
}

// GenerateGalaxy creates the specified number of stars in a round galaxy
func (s *Service) GenerateGalaxy(numberOfStars int, galaxyRadius float64) {
	utils.ShuffleStringSlice(names)
	for i := 0; i < numberOfStars; i++ {
		coordinates := core.CreateVectorWithinRadius(galaxyRadius)
		classification := GetRandomClassification()

		star := &Star{
			ID:             StarID(i),
			Name:           names[i],
			Coordinates:    coordinates,
			Classification: classification.name,
			Temperature:    classification.temperature.RandomValue(),
			Mass:           classification.mass.RandomValue(),
			Radius:         classification.radius.RandomValue(),
			Color:          classification.color,
		}

		s.stars[star.ID] = star
	}
}
