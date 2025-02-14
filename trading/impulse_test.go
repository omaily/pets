package main

import (
	"fmt"
	"testing"
)

func TestComputeImpulse(t *testing.T) {
	testDataImpuls := make([]string, len(testData))
	for i := range testData {
		testDataImpuls[i] = testData[i].ImpulseStr
	}

	candles, err := readCSV("data.csv")
	if err != nil {
		t.Errorf("failed to read file")
	}

	impulse := computeImpulse(candles)

	for i := 0; i < len(testDataImpuls); i++ {
		reliable := testDataImpuls[i]
		temp := impulse[i]

		if reliable != fmt.Sprintf(`"default": %v`, temp) {
			t.Errorf("Row %d: Expected impuls = %v, Got: %v", i, reliable, temp)
		}
	}
}
