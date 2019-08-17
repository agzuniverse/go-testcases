package components

import (
	"errors"
	"math/rand"
	"time"
)

// GenSingleRandom generates a single random number between the limits min and max
func GenSingleRandom(min int, max int) (int, error) {
	rand.Seed(time.Now().UnixNano())
	if min >= max {
		return 0, errors.New("Min value needs to be lower than max value")
	}
	if min < 0 || max <= 0 {
		return 0, errors.New("Both min and max values needs to be greater than 0")
	}
	return (rand.Intn(max-min) + min), nil
}
