package main

import (
	"fmt"
	"sort"
)

func main() {
	// m := [][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}
	// m := [][]int{{1, 2}}
	m := [][]int{{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, {19, 18, 17, 16, 15, 14, 13, 12, 11, 10}, {20, 21, 22, 23, 24, 25, 26, 27, 28, 29}, {39, 38, 37, 36, 35, 34, 33, 32, 31, 30}, {40, 41, 42, 43, 44, 45, 46, 47, 48, 49}, {59, 58, 57, 56, 55, 54, 53, 52, 51, 50}, {60, 61, 62, 63, 64, 65, 66, 67, 68, 69}, {79, 78, 77, 76, 75, 74, 73, 72, 71, 70}, {80, 81, 82, 83, 84, 85, 86, 87, 88, 89}, {99, 98, 97, 96, 95, 94, 93, 92, 91, 90}, {100, 101, 102, 103, 104, 105, 106, 107, 108, 109}, {119, 118, 117, 116, 115, 114, 113, 112, 111, 110}, {120, 121, 122, 123, 124, 125, 126, 127, 128, 129}, {139, 138, 137, 136, 135, 134, 133, 132, 131, 130}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}

	fmt.Println(longestIncreasingPath(m)) // 3
}

var sortedArray [][2]int

var matrixToArrayIndex map[int]int // 矩阵的下标到排序后数字的下标的对应关系

var maxLength = 0

var matrixColumnCount = 0
var matrixRowCount = 0

func longestIncreasingPath(matrix [][]int) int {
	sortedArray = make([][2]int, 0)
	matrixToArrayIndex = make(map[int]int)
	maxLength = 0
	matrixColumnCount = 0
	matrixRowCount = 0

	sortMatrix(matrix)
	matrixColumnCount = len(matrix[0])
	matrixRowCount = len(matrix)
	if len(matrix) == 1 && len(matrix[0]) == 1 {
		return 1
	}
	// maxLengthOnSortedArray(false, 3)

	for i, _ := range sortedArray {
		if i+maxLength >= len(sortedArray) {
			break
		}
		maxLengthOnSortedArray(false, i)
	}
	return maxLength
}

// SortedArray 从 fromIndex 到队尾的最长值
func maxLengthOnSortedArray(nested bool, fromArrayIndex int) int {
	current := sortedArray[fromArrayIndex]
	if !nested {

	}
	currentValue := current[0]
	mi := current[1] // matrixIndex
	left := mi - 1
	right := mi + 1
	upper := mi - matrixColumnCount
	downer := mi + matrixColumnCount
	validMaIndexes := make([]int, 0, 4)
	if mi%matrixColumnCount == 0 {
		left = -1
	}
	if (mi+1)%matrixColumnCount == 0 {
		right = -1
	}
	if mi >= 0 && mi < matrixColumnCount {
		upper = -1
	}
	if mi >= (matrixRowCount-1)*matrixColumnCount {
		downer = -1
	}
	aroundTemp := []int{left, right, upper, downer}
	for _, v := range aroundTemp {
		if v == -1 {
			continue
		}
		if arrayIndex, ok := matrixToArrayIndex[v]; ok && arrayIndex > fromArrayIndex {
			validMaIndexes = append(validMaIndexes, v)
		}
	}
	fmt.Println(mi, ".........validMaIndexes....", validMaIndexes)
	maxLengthArray := make([]int, 0, 4)
	for _, around := range validMaIndexes {
		// fmt.Println(mi, ".........around....", around)
		arrayIndex, ok := matrixToArrayIndex[around]
		// fmt.Println(arrayIndex, "---------ok-------", ok)
		if ok {
			aroundValue := sortedArray[arrayIndex][0]
			// MIndex := sortedArray[arrayIndex][1]
			// fmt.Println(around, "---------around-------", sortedArray[arrayIndex])
			if aroundValue > currentValue {
				// fmt.Println(fromArrayIndex, ".....maxLengthOnSortedArray........", MIndex)
				maxLengthArray = append(maxLengthArray, maxLengthOnSortedArray(false, arrayIndex))
			} else {
				maxLengthArray = append(maxLengthArray, 0)
			}
		} else {
			maxLengthArray = append(maxLengthArray, 0)
		}
	}
	if len(maxLengthArray) > 0 {
		sort.Slice(maxLengthArray, func(i, j int) bool { return maxLengthArray[i] < maxLengthArray[j] })
		if maxLengthArray[len(maxLengthArray)-1]+1 > maxLength {
			maxLength = maxLengthArray[len(maxLengthArray)-1] + 1
		}
		return maxLengthArray[len(maxLengthArray)-1] + 1
	}
	return 1
}

func sortMatrix(matrix [][]int) {
	columnCount := len(matrix[0])
	for i, _ := range matrix {
		for j, _ := range matrix[i] {
			serialNumber := i*columnCount + j
			sortedArray = append(sortedArray, [2]int{matrix[i][j], serialNumber})
		}
	}
	sort.SliceStable(sortedArray, func(i, j int) bool { return sortedArray[i][0] < sortedArray[j][0] })
	matrixToArrayIndex = make(map[int]int, len(sortedArray))
	for i, v := range sortedArray {
		matrixToArrayIndex[v[1]] = i
	}
}
