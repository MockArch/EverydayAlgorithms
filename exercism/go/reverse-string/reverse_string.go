package reverse

func Reverse(r string) string {
	run_string := []rune(r)
	for i, j := 0, len(run_string)-1; i < j; i, j = i+1, j-1 {
		run_string[i], run_string[j] = run_string[j], run_string[i]
	}
	return string(run_string)
}
