package Leetcode

func groupThePeople(groupSizes []int) [][]int {
    var res [][]int
    m := make(map[int][]int)
    for i, v := range groupSizes {
        m[v] = append(m[v], i)
        if v == len(m[v]) {
            res = append(res, m[v])
            m[v] = []int{}
        }
    }
    return res
}