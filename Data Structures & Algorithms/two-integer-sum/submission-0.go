func twoSum(nums []int, target int) []int {
    for i := 0; i<len(nums)-1; i++ {
		for j := 1; j<len(nums); j++ {
			if(i != j && (target == (nums[i]+nums[j]))){
				return []int{i, j}
			}
		}
	} 
	var empty = []int{}
	return empty
}