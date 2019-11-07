package diffsquares

func SquareOfSum(input int) int {
	ret := 0
	for i := 1; i <= input; i++ {
		ret += i
	}
	return ret*ret
}

func SumOfSquares(input int) int {
	ret := 0
	for i := 1; i <= input; i++ {
		ret += i*i
	}
	return ret
}

func Difference(input int) int {
	return SquareOfSum(input) - SumOfSquares(input)
}