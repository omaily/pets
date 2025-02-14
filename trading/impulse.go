package main

import "math"

// скользящее среднее
func ta_sma(values []float64) float64 {
	summ := 0.0
	for _, v := range values {
		summ += v
	}
	return summ / float64(len(values))
}

// ta.rma() = Rolling Moving Average (сглаженное среднее)
// RMA = ((RMA[t-1] * (n-1)) + Xt) / n
func ta_rma(values []float64, window int) []float64 {
	period := float64(window)
	var result []float64
	var prev float64 = values[0]

	result = append(result, prev)
	for i := 1; i < len(values); i++ {
		prev = prev*(period-1) + values[i]
		prev /= period
		result = append(result, prev)
	}
	return result
}

// ta.atr - Average True Range - (средний истинный диапазон)
// который измеряет средний истинный диапазон цены за определённый период.
// ATR учитывает не только дневные максимумы и минимумы, но и возможные ценовые разрывы (гэпы).
// ATR = ((ATR[t-1] * (n-1)) + tr) / n
//
// ta.tr (истинный диапазон) - выбираем максимальное значение разрыва цены.
// tr := max(high - low, abs(high - close[t-1]), abs(low - close[t-1])).
func ta_atr(graph []Candle, mult, window int) []FloatNull {
	var result []FloatNull

	var trLenght []float64
	var prevClose float64 = -1
	for _, s := range graph {
		trLenght = append(trLenght, ta_tr(s.High, s.Low, prevClose))
		prevClose = s.Close
	}

	if len(graph) != len(trLenght) {
		return nil //error
	}

	var prevAtr float64
	for i := 0; i < window; i++ {
		if i < window-1 {
			result = append(result, FloatNull(math.NaN()))
			continue
		}
		prevAtr = ta_sma(trLenght[:i])
	}
	result = append(result, FloatNull(float64(mult)*prevAtr))

	period := float64(window)
	for i := window; i < len(graph); i++ {
		prevAtr := prevAtr*(period-1) + trLenght[i]
		prevAtr /= period
		result = append(result, FloatNull(float64(mult)*prevAtr))
	}

	return result
}

func ta_tr(high, low, close float64) float64 {
	if close < 0 {
		return high - low
	}
	return max(math.Abs(high-low), math.Abs(high-close), math.Abs(low-close))
}

func computeImpulse(barChart []Candle) []int {
	result := make([]int, len(barChart))

	shortWindow := 10 / 2
	mediumWindow := 30 / 2
	shortMult := 1
	mediumMult := 3

	graph := make([]float64, len(barChart))
	for i := 0; i < len(barChart); i++ {
		graph[i] = barChart[i].Close
	}

	rmaShort := ta_rma(graph, shortWindow)
	rmaMedium := ta_rma(graph, mediumWindow)

	atrShort := ta_atr(barChart, shortMult, shortWindow)
	atrMedium := ta_atr(barChart, mediumMult, mediumWindow)

	for i, j := 0, 0; i < len(barChart); i, j = i+1, j+1 {
		var rsh float64
		if i < shortWindow/2 {
			rsh = barChart[i].Close
		} else {
			rsh = rmaShort[i-shortWindow/2]
		}

		var msh float64
		if i < mediumWindow/2 {
			msh = barChart[i].Close
		} else {
			msh = rmaMedium[i-mediumWindow/2]
		}

		sct := FloatNull(rsh) + atrShort[i]
		scb := FloatNull(rsh) - atrShort[i]
		mct := FloatNull(msh) + atrMedium[i]
		mcb := FloatNull(msh) - atrMedium[i]
		scmm := (sct + scb) / 2

		omed := (scmm - mcb) / (mct - mcb)
		oshort := (FloatNull(barChart[i].Close) - mcb) / (mct - mcb)

		if oshort > 0.5 {
			if oshort > 1.0 {
				barChart[i].Impulse = 2 //lime
				result[i] = 2
			} else if oshort > omed {
				barChart[i].Impulse = 1 //green
				result[i] = 1
			} else {
				barChart[i].Impulse = 0 //white
				result[i] = 0
			}
		} else if oshort < 0 {
			barChart[i].Impulse = -2 //red
			result[i] = -2
		} else if oshort < omed {
			barChart[i].Impulse = -1 //borrow
			result[i] = -1
		} else {
			barChart[i].Impulse = 0 //white
			result[i] = 0
		}
	}
	return result
}
