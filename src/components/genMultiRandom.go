package components

import (
	"errors"
)

// GenMultiRandom generates n random numbers between min and max.
func GenMultiRandom(n int, min int, max int) ([]int, error) {
	if n <= 0 {
		return nil, errors.New("n must be at least 1")
	}
	result := make([]int, 0)
	for i := 0; i < n; i++ {
		v, err := GenSingleRandom(min, max)
		if err != nil {
			return nil, err
		}
		result = append(result, v)
	}
	return result, nil
}
