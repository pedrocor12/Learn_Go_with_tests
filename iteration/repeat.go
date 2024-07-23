package iteration

func Repeat(character string, toRepeat int) string {
	var repeated string
	for i := 0; i < toRepeat; i++ {
		repeated += character
	}
	return repeated
}
