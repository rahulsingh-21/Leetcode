package main

func removeStars(s string) string {

	var str []byte

	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			str = str[:len(str)-1]
		} else {
			str = append(str, s[i])
		}
	}
	return string(str)
}
