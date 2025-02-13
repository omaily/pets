package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

func readCSV(filename string) ([]Candle, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var candles []Candle
	for i, row := range records {
		if i == 0 {
			continue
		}
		time, _ := strconv.ParseInt(row[0], 10, 64)
		open, _ := strconv.ParseFloat(row[1], 64)
		high, _ := strconv.ParseFloat(row[2], 64)
		low, _ := strconv.ParseFloat(row[3], 64)
		close, _ := strconv.ParseFloat(row[4], 64)
		candles = append(candles, Candle{Time: time, Open: open, High: high, Low: low, Close: close})
	}

	return candles, nil
}
