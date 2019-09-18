package array

func TwoSum(nums []int, target int) []int {
    m := make(map[int]int)
    var fir int
    var snd int
    for index, val := range nums {
        if i := m[val]; i != 0 {
            fir = i - 1
            snd = index
            break
        }
        sub := target - val
        m[sub] = index + 1
    }
    res := make([]int, 2)
    res[0] = fir
    res[1] = snd
    return res
}

