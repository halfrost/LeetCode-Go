package leetcode

func removeDuplicates80(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	last, finder, startFinder := 0, 0, -1
	for last < len(nums)-1 {
		startFinder = -1
		for nums[finder] == nums[last] {
			if startFinder == -1 || startFinder > finder {
				startFinder = finder
			}
			if finder == len(nums)-1 {
				break
			}
			finder++
		}
		if finder-startFinder >= 2 && nums[finder-1] == nums[last] && nums[finder] != nums[last] {
			nums[last+1] = nums[finder-1]
			nums[last+2] = nums[finder]
			last += 2
		} else {
			nums[last+1] = nums[finder]
			last++
		}
		if finder == len(nums)-1 {
			if nums[finder] != nums[last-1] {
				nums[last] = nums[finder]
			}
			return last + 1
		}
	}
	return last + 1
}
