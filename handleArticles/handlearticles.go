package articles

func HandleArticles(s []string) []string {
	vowels := []string{"a", "e", "i", "o", "u", "h", "A", "E", "I", "O", "U", "H"}
	for i, word := range s {
		for _, vow := range vowels {
			if word == "a" && string(s[i+1][0]) == vow {
				s[i] = "an"
			}
			if word == "A" && string(s[i+1][0]) == vow {
				s[i] = "An"
			}
		}
	}
	return s
}
