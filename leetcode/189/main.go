package main

func rotate(nums []int, k int) {
	if k >= len(nums) {
		k = k % len(nums)
	}
	temp := []int{}
	for i := len(nums) - k; i < len(nums); i++ {
		temp = append(temp, nums[i])
	}
	for i := 0; i < len(nums)-k; i++ {
		temp = append(temp, nums[i])
	}
	copy(nums, temp)
}

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	rotate([]int{1, 2, 3}, 4)
}
