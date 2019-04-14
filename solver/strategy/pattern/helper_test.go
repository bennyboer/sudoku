package pattern

import (
	"testing"
)

func Test_binomialCoefficient(t *testing.T) {
	result := binomialCoefficient(4, 2)
	if result != 6 {
		t.Errorf("Expected binomial coefficient from n: %d and k: %d to be %d", 4, 2, 6)
	}

	result = binomialCoefficient(4, 3)
	if result != 4 {
		t.Errorf("Expected binomial coefficient from n: %d and k: %d to be %d", 4, 3, 4)
	}
}
