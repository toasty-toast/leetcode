func twoSum(nums []int, target int) []int {
    valueMap := make(map[int]int)
    for i, num := range nums {
        index1, success := valueMap[target - num]
        if success {
            return []int{index1, i}
        } else {
            valueMap[num] = i
        }
    }
    return nil;
}