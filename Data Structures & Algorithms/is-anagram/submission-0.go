func isAnagram(s string, t string) bool {
	mapS := make(map[rune]int)
	mapT := make(map[rune]int)

	if len(s) != len(t){
		return false
	}

	for i, ch := range s {
		mapS[ch]++
		mapT[rune(t[i])]++
	}

	for char, count := range mapS {
        if count != mapT[char] {
            return false
        }
    }

	return true
}
