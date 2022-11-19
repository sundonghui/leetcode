package spiralorder

// spiralOrder
// 1 2 3
// 4 5 6
// 7 8 9
// 1,2,3,6,9,8,7,4,5
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
        return []int{}
    }
	
	list := []int{}
	col, row := len(matrix[0]), len(matrix)
	top,bottom,left,right :=0,row-1,0,col-1
	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {  //top行的从左到右遍历
			list = append(list, matrix[top][i])
		}
		top++
		for i:=top;i<=bottom;i++{   //right列的从上到下遍历
            list =append(list, matrix[i][right])
        }
		right--
		if left>right || top>bottom{  //这里一定要做一个判断，因为如果只剩下一个元素就会出错
            break
        }
		for i:=right;i>=left;i--{      //bottom行的从右往左遍历
            list=append(list,matrix[bottom][i])
        }
        bottom--                         //bottom往上移动
        for i:=bottom;i>=top;i--{        //left列的从下到上遍历
            list=append(list,matrix[i][left])
        }
        left++    

	}
	return list
}	