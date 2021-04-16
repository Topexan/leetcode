package Leetcode

type SubrectangleQueries struct {
    data [][]int
}


func Constructor(rectangle [][]int) SubrectangleQueries {
    var res SubrectangleQueries
    res.data = rectangle
    return res
}


func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int)  {
    for i := row1; i <= row2; i++ {
        for j := col1; j <= col2; j++ {
            this.data[i][j] = newValue
        }
    }
}


func (this *SubrectangleQueries) GetValue(row int, col int) int {
    return this.data[row][col]
}