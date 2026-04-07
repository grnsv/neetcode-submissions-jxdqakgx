func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    res := [][]int{}
    for i := 0; i < len(nums)-2; i++ {
        if nums[i] > 0 {
            break
        }
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        l := i+1
        r := len(nums)-1
        for l < r {
            sum := nums[i] + nums[l] + nums[r]
            if sum < 0 {
                l++
                continue
            }
            if sum > 0 {
                r--
                continue
            }
            if sum == 0 {
                res = append(res, []int{nums[i], nums[l], nums[r]})
                l++
                for l < r && nums[l] == nums[l-1] {
                    l++
                }
                r--
                for l < r && nums[r] == nums[r+1] {
                    r--
                }
            }
        }
    }

    return res
}
