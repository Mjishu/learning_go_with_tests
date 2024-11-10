package iterator

func Repeat(item string, counts int) string {
	var repeated string
	for i := 0; i < counts; i++ {
		repeated += item
	}
	return repeated
}
