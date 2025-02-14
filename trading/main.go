package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
)

type Candle struct {
	Time       int64   `json:"time"`
	Open       float64 `json:"open"`
	High       float64 `json:"high"`
	Low        float64 `json:"low"`
	Close      float64 `json:"close"`
	T3         float64 `json:"t3"`
	Impulse    int     `json:"impulse"`
	ImpulseStr string  `json:"default"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // Замените на нужный домен
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == "OPTIONS" {
		return
	}

	candles, err := readCSV("data.csv")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = computeT3(candles, make(map[string]float64))
	_ = computeImpulse(candles)

	// Отправляем данные в виде JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(candles)
}

func main() {
	http.HandleFunc("/candles", handler)
	log.Println("Server running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}

// ComputeT3 Функция для вычисления T3 (версия с map[string]float64)
func computeT3(prices []Candle, params map[string]float64) error {
	length := 10
	if val, ok := params["length"]; ok {
		length = int(val)
	}

	// Параметры для расчета T3
	vFactor := float64(1.0)
	closePrices := make([]float64, len(prices))
	for i, p := range prices {
		closePrices[i] = p.Close
	}

	if len(closePrices) < length {
		return fmt.Errorf("closePrices < length")
	}

	xe1 := calculateEMA2(closePrices, length)
	xe2 := calculateEMA2(xe1, length)
	xe3 := calculateEMA2(xe2, length)
	xe4 := calculateEMA2(xe3, length)
	xe5 := calculateEMA2(xe4, length)
	xe6 := calculateEMA2(xe5, length)
	var (
		dThree = float64(3)
		c1     = 0 - math.Pow(vFactor, 3)
		c2     = (dThree * math.Pow(vFactor, 2)) + (dThree * math.Pow(vFactor, 3))
		c3     = 0 - (6 * vFactor * vFactor) - (dThree * vFactor) - (dThree * math.Pow(vFactor, 3))
		c4     = 1 + (dThree * (vFactor)) + math.Pow(vFactor, 3) + (dThree * math.Pow(vFactor, 3))
	)

	for i := 0; i < len(prices); i++ {
		prices[i].T3 = c1*xe6[i] + c2*xe5[i] + c3*xe4[i] + c4*xe3[i]
	}

	return nil
}

func calculateEMA2(prices []float64, period int) []float64 {
	ema := make([]float64, len(prices))

	alpha := 2.0 / float64(period+1)
	alpha2 := 1 - alpha
	ema[0] = prices[0]
	for i := 1; i < len(prices); i++ {
		ema[i] = prices[i]*alpha + ema[i-1]*alpha2
	}

	return ema
}
