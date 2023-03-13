package piscine

func ConcatParams(args []string) string {
	var new_string string
	for i := 0; i < len(args); i++ {
		new_string = new_string + args[i]
		if i != len(args)-1 {
			new_string = new_string + "\n"
		}
	}
	return new_string
}
