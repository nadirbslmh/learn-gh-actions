package main

func main() {

}

func sum(numbers []int) int {
	var result int = 0

	for _, number := range numbers {
		result += number
	}

	return result
}
