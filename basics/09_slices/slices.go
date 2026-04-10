package main

import "fmt"

// slice
// most used construct in go
// +useful methods
func main() {
	// uninitialized slice is nil
	// var numbers []int
	// // fmt.Println(numbers == nil)  // true
	// fmt.Println(len(numbers))

	/*
	   	// for not nil
	   	var nums = make([]int, 2, 5)

	   	// fmt.Println(cap(nums))
	   	// 	fmt.Println(nums == nil) // false
	   	// 	fmt.Println(nums)
	   	nums = append(nums, 1)
	   	nums = append(nums, 2)

	   	fmt.Println(nums)
	   	fmt.Println(cap(nums))

	    str := []string{}

	    str = append(str, "hello")
	    str = append(str, "world")
	    fmt.Println(str == nil)
	    fmt.Println(len(str))
	    fmt.Println(cap(str))
	*/

	// Copy Function

	// var nums = make([]int, 2, 5)

	// nums = append(nums, 2)
	// var nums2 = make([]int, len(nums))
	// copy(nums2, nums)
	// fmt.Println(nums, nums2)

	// var nums = []int{1, 2, 3, 4, 5}
	// // excludes last last index
	// fmt.Println(nums[1:3]) // slice operator from 1 index
	// fmt.Println(nums[:2]) // slice operator from 0
	// fmt.Println(nums[2:]) // slice operator from 2 to last index

	// slice package
	// var nums1 = []int {1, 2}
	// var nums2 = []int {1, 2, 3, 4}

	// fmt.Println(slices.Equal(nums1, nums2)) // checks equality

	// TwoD slices

	var nums = [][]int{{1, 2}, {3, 4}}

	fmt.Println(nums) // [[1 2] [3 4]]

}
