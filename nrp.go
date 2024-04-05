package main

import "fmt"

func nrp(s []int) []int {
	length := len(s)
	resultLength := make([]int, length)
	last := make([]int, length)

	// init
	for i := range resultLength {
		resultLength[i] = 1
		last[i] = -1
	}

	for i := 1; i < length; i++ {
		for j := 0; j < i; j++ {
			if s[i] > s[j] && resultLength[i] < resultLength[j]+1 {
				resultLength[i] = resultLength[j] + 1
				last[i] = j
			}
		}
	}

	i := maxIndex(resultLength)
	var res []int
	for i != -1 {
		res = append(res, s[i])
		i = last[i]
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}

func maxIndex(s []int) int {
	maxIdx := 0
	maxValue := s[0]
	for i, v := range s {
		if v > maxValue {
			maxValue = v
			maxIdx = i
		}
	}
	return maxIdx
}


func main(){
	data := []int{100, 3, 14, 15, 92, 65, 35, 89, 79, 32, 38, 46, 26, 43}
	
	res := nrp(data)
	
	fmt.Println(res)
}