package utils

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func Lcm(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = (result * numbers[i]) / Gcd(result, numbers[i])
	}
	return result
}

func Abs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}
