package leetcode

func numJewelsInStones(jewels string, stones string) int {
    c := 0
    for i :=  range jewels {
        for j := range stones{
            if jewels[i] == stones[j]{
                c++
            }
        }
    }
    return c
}