package testProject

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("[1,2,3,4,5]", testSumFunc([]int{1, 2, 3, 4, 5}, 15))
	t.Run("[1,2,3,4,-5]", testSumFunc([]int{1, 2, 3, 4, -5}, 5))
}


func testSumFunc(numbers []int,expected int)func(t*testing.T){
	return func(t *testing.T) {
		actual := Sum(numbers)
		if actual != expected {
			t.Errorf("Expected the sum of %v to be %d but instead got %d!", numbers, expected, actual)
		}
	}
}

func ExampleSum() {
	numbers := []int{5, 5, 5}
	fmt.Println(Sum(numbers))
	// Output:
	// 15
}