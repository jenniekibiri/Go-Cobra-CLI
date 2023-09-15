package validation

func IsEmpty(s string) bool {
    return len(s) == 0 || (len(s) > 0 && len(s) == len(s[:1]) && s[:1] == " ")
}

func IsEmptyInt(i int) bool {
	return i == 0
}

func IsEmptyFloat(f float64) bool {

	return f == 0.0
}