package leetcode

func circularArrayLoop(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		// slow/fast pointer
		slow, fast, val := i, getNextIndex(nums, i), 0
		for nums[fast]*nums[i] > 0 && nums[getNextIndex(nums, fast)]*nums[i] > 0 {
			if slow == fast {
				// check for loop with only one element
				if slow == getNextIndex(nums, slow) {
					break
				}
				return true
			}
			slow = getNextIndex(nums, slow)
			fast = getNextIndex(nums, getNextIndex(nums, fast))
		}
		// loop not found, set all element along the way to 0
		slow, val = i, nums[i]
		for nums[slow]*val > 0 {
			next := getNextIndex(nums, slow)
			nums[slow] = 0
			slow = next
		}
	}
	return false
}

func getNextIndex(nums []int, index int) int {
	return ((nums[index]+index)%len(nums) + len(nums)) % len(nums)
}
