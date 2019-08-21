package components

import (
	"errors"
	"math/rand"
)

// GenSingleRandom generates a single random number between the limits min and max
func GenSingleRandom(min int, max int) (int, error) {
	if min >= max {
		return 0, errors.New("Min value needs to be lower than max value")
	}
	return (rand.Intn(max-min) + min), nil
}
