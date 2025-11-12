package main

func IsCapital(s rune)bool {
	if s >= 'A' && s <= 'Z' {
		return true
	}
	return false
}

func CamelToSnakeCase(s string) string{
	if len(s) == 0 {
		return ""
	}

	for i, _ := range s {
		if IsCapital(s[i]) == IsCapital(s[i+1]){
			return s
		}
		if IsCapital(s[i]){
			return s
		}
	}

	for i := range s[1] {
		if IsCapital(s[i]) {
			s[i] = "_"
		}
	}

}