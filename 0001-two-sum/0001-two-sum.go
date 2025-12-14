func twoSum(nums []int, target int) []int {
    queue := make(map[int]int, len(nums))
    var complement int
    for i, num := range nums {
        complement = target - num
        if index, ok := queue[complement]; ok {
            return []int{index,i}
        }
        queue[num] = i
    }
    return []int{}
}