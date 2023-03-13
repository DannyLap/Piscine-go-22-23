package piscine

func ft_max_substring(s string) int {
	var string_max string
	for i := 0; i < len(s); i++ {
		chaine_tmp := ""

		for j := i; j < len(s); j++ {
			oui_non := true

			for k := 0; k < len(chaine_tmp); k++ {
				if s[j] == chaine_tmp[k] {
					oui_non = false
				}
			}

			if oui_non {
				chaine_tmp = chaine_tmp + string(s[j])
				if len(s) == len(chaine_tmp) {
					string_max = chaine_tmp
					chaine_tmp = ""
				} else if chaine_tmp == s[len(s)-len(chaine_tmp):] && len(chaine_tmp) >= len(string_max) {
					string_max = chaine_tmp
					chaine_tmp = ""
				}
			} else if len(chaine_tmp) >= len(string_max) {
				string_max = chaine_tmp
				chaine_tmp = ""
			}
		}
	}
	return len(string_max)
}
