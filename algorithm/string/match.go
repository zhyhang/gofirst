package string

// find the shortest sub-string of s, the sub-string contains the letters
func SearchShortestSubString(s, letters string) string {
	if s == "" || letters == "" {
		return ""
	}
	rs := []rune(s)
	shortestSubInfo := []int{0, len(rs) + 2}
	letterMap := makeLetterMap(letters)
	subBegin := -1
	for i := 0; i < len(rs); i++ {
		if letterMap[rs[i]] {
			delete(letterMap, rs[i])
			if subBegin == -1 {
				if len(letterMap) == 0 {
					return letters //the letters only contains one letter
				}
				subBegin = i
			} else if len(letterMap) == 0 {
				if i+1-subBegin < shortestSubInfo[1]-shortestSubInfo[0] {
					shortestSubInfo[0] = subBegin
					shortestSubInfo[1] = i + 1
				}
				letterMap = makeLetterMap(letters)
				i = subBegin
				subBegin = -1
			}
		}
	}
	if shortestSubInfo[1] != len(rs)+2 {
		return string(rs[shortestSubInfo[0]:shortestSubInfo[1]])
	}
	return ""
}

func makeLetterMap(letters string) map[rune]bool {
	lmap := make(map[rune]bool)
	for _, l := range letters {
		lmap[l] = true
	}
	return lmap
}

func IndexOfAll(txt, sub string) (subIndexes []int) {
	subIndexes = nil
	if txt == "" && sub == "" {
		subIndexes = []int{0}
		return
	} else if txt == "" || sub == "" {
		return
	}
	txtSlice := []rune(txt)
	subSlice := []rune(sub)
	for i := 0; i < len(txtSlice); {
		if len(txtSlice)-i < len(subSlice) {
			break
		}
		for j := 0; j < len(subSlice); {
			if txtSlice[i] == subSlice[j] {
				if j == len(subSlice)-1 {
					subIndexes = append(subIndexes, i-j)
					i = i - j + 1 //lookup next matched
					break
				}
				i++
				j++

			} else {
				i = i - j + 1
				break
			}
		}
	}
	return
}

func IndexOfAllSimple(txt, sub string) (subIndexes []int) {
	if txt == "" && sub == "" {
		return []int{0}
	}
	if txt == "" || sub == "" {
		return nil
	}
	subIndexes = nil
	txtUnicodes := []rune(txt)
	subUnicodes := []rune(sub)
	for i, j := 0, 0; i < len(txtUnicodes) && j < len(subUnicodes); {
		if txtUnicodes[i] == subUnicodes[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
		if j == len(subUnicodes) {
			subIndexes = append(subIndexes, i-j)
			i = i - j + 1
			j = 0
		}
	}
	return
}
