package array

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(TwoSum(nums, target))
}
