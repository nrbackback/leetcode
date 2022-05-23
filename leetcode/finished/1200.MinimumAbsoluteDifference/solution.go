package solution

import (
	"fmt"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	fmt.Println("...arr..", arr)
	minVal := arr[1] - arr[0]
	out := make([][]int, 0)
	for j := 0; j < len(arr)-1; j++ {
		if arr[j+1]-arr[j] > minVal {
			continue
		}
		if arr[j+1]-arr[j] < minVal {
			minVal = arr[j+1] - arr[j]
			out = make([][]int, 0)
		}
		out = append(out, []int{arr[j], arr[j+1]})
	}
	return out
	/*
		out := make([][]int, 0)
		minVal := 0
		for j := 0; j < len(arr); j++ {
			for i := len(arr) - 1; i > j; i-- {
				if arr[i] < arr[i-1] {
					temp := arr[i-1]
					arr[i-1] = arr[i]
					arr[i] = temp
					fmt.Println(i, "----iii----", arr)
				}
			}
			fmt.Println(j, "--------", arr)
			if j == 0 {
				minVal = arr[j+1] - arr[j]
				fmt.Println("---minVal----", minVal)
				// out = append(out, []int{arr[j], arr[j+1]})
				continue
			}
			// if j == 1 {
			// 	// out = make([][]int, 0)
			// }
			fmt.Println(arr[j], arr[j-1], ".......", arr[j]-arr[j-1])
			if arr[j]-arr[j-1] > minVal {
				continue
			}
			if arr[j]-arr[j-1] < minVal {
				minVal = arr[j] - arr[j-1]
				out = make([][]int, 0)
			}
			out = append(out, []int{arr[j-1], arr[j]})
		}

		return out
	*/
}
