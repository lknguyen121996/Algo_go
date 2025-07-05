/*
Bạn là một tên trộm chuyên nghiệp đang lên kế hoạch cướp các ngôi nhà trên một con phố.
Mỗi nhà có một số tiền nhất định, nhưng bạn không thể cướp hai nhà liền kề
(hệ thống báo động sẽ kích hoạt).

Cho một mảng nums đại diện cho số tiền trong mỗi nhà,
hãy trả về số tiền tối đa bạn có thể cướp mà không kích hoạt báo động.

Ví dụ:
Input: nums = [2,7,9,3,1]
Output: 12 (cướp nhà 0, 2, 4: 2 + 9 + 1 = 12)
*/
package main

// var cache = map[int]int{}

func main() {
	// Ví dụ sử dụng
	nums := []int{2, 7, 9, 3, 1}
	result := rob(nums)
	println(result) // Output: 12
}

func rob(nums []int) int {
	buoc := len(nums) - 2
	cach := len(nums) - 3
	if buoc < 0 {
		return 0
	}
	if cach < 0 {
		return 0
	}
	// Viết code của bạn ở đây
	if len(nums) == 1 {
		return nums[0]
	}
	return rob(nums[:cach]) + rob(nums[:buoc])
}

func max(nums []int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}
