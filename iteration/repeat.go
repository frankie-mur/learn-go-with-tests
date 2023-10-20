package iteration

const repeatCount = 3

// Repeat Can use varadic param so imulate optional
func Repeat(s string, count ...int) string {
	var repCount int
	if len(count) == 0 {
		repCount = repeatCount
	} else {
		repCount = count[0]
	}
	var result string
	for i := 0; i < repCount; i++ {
		result += s
	}
	return result
}
