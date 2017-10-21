package math

import (
	"errors"
)

// big integer subtract with int arrays
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
	aux := make([]int, len(sumCalc))
	copy(aux, sumCalc)
	sumCalc = aux
	result := make([]int, len(sumCalc))
	for i1, i2 := len(sumCalc)-1, len(subtractorCalc)-1; i1 >= 0; i1, i2 = i1-1, i2-1 {
		if i2 < 0 {
			result[i1] = sumCalc[i1]
		} else if sumCalc[i1] >= subtractorCalc[i2] {
			result[i1] = sumCalc[i1] - subtractorCalc[i2]
		} else {
			// borrow from high digit
			result[i1] = 10 + sumCalc[i1] - subtractorCalc[i2]
			for j := i1 - 1; j >= 0; j-- {
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
