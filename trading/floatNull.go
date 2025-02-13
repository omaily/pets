package main

import (
	"encoding/json"
	"math"
)

type FloatNull float64

func (f FloatNull) MarshalJSON() ([]byte, error) {
	if math.IsNaN(float64(f)) {
		return []byte(`"NaN"`), nil // NaN → null
	}
	return json.Marshal(float64(f))
}

func (f *FloatNull) UnmarshalJSON(data []byte) error {
	if string(data) == `"NaN"` {
		*f = FloatNull(math.NaN())
		return nil
	}

	// var temt float64
	// if err := json.Unmarshal(data, &temt); err != nil {
	// 	return err
	// }
	// *f = FloatNull(temt)
	return json.Unmarshal(data, (*float64)(f))
}
