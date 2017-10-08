package math

import (
	"errors"
)

func BigIntSubtract(sum, subtractor []int) ([]int, int, error) {
	if sum == nil || subtractor == nil {
		return nil, 0, errors.New("sum or subtractor is nil")
	}
	sumCalc := sum
	subtractorCalc := subtractor
	sign := 0
	if findBigger(sum, subtractor) == 1 {
		sign = 1
		sumCalc = subtractor
		subtractorCalc = sum
	}
	result := make([]int, max(len(sum), len(subtractor)))
	for i := len(sumCalc) - 1; i >= 0; i-- {
		if sumCalc[i] >= subtractorCalc[i] {
			result[i] = sumCalc[i] - subtractorCalc[i]
		} else {
			// borrow from high digit
			result[i] = 10 + sumCalc[i] - subtractorCalc[i]
			for j := i - 1; j >= 0; j-- {
				if sumCalc[j] >= 1 {
					sumCalc[j]--
					break
				} else {
					sumCalc[j] = 9
				}
			}
		}
	}
	return result, sign, nil
}

func findEffectDigits(a []int) int {
	elen := 0
	for i := 0; i < len(a); i++ {
		if a[i] != 0 {
			elen = len(a) - i
			break
		}
	}
	return elen
}

func findBigger(a, b []int) int {
	alen := findEffectDigits(a)
	blen := findEffectDigits(b)
	if alen > blen {
		return 0
	} else if alen < blen {
		return 1
	} else {
		// have same digits
		for i := 0; i < alen; i++ {
			indexDelta := alen - i
			if a[len(a)-indexDelta] > b[len(b)-indexDelta] {
				return 0
			} else if a[len(a)-indexDelta] < b[len(b)-indexDelta] {
				return 1
			}
		}
		return 0
	}

}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
