func twoSum(nums []int, target int) []int {
    valueMap := make(map[int]int);
    for i := 0; i < len(nums); i++ {
        index, success := valueMap[target - nums[i]];
        if success {
            return []int{index, i};
        }
        valueMap[nums[i]] = i;
    }
    return nil;
}