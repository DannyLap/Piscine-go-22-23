package piscine

func SplitWhiteSpaces(s string) []string {
	var string_tmp string
	var tab_string []string

	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " {
			if len(string_tmp) >= 1 {
				tab_string = append(tab_string, string_tmp)
			}
			string_tmp = ""
		} else if i == len(s)-1 {
			string_tmp = string_tmp + string(s[i])
			tab_string = append(tab_string, string_tmp)
		} else {
			string_tmp = string_tmp + string(s[i])
		}
	}
	return tab_string
}
