package testing

const (
	Negative = 1
	Zero = 2
	Positive = 3
)

func DetermineType(n int) int {
	if(n < 0) {
		return Negative
	} else if (n == 0) {
		return Zero
	} else {
		return Positive
	}
}