package main


import "fmt"

func main() {

	s := []int{100,34,3,5,56,676,87,4,1}
	fmt.Println(s)
	sort(s, 0, len(s) -1)
	fmt.Println(s)
}

func sort(s []int, begin int, end int) {
	left := begin
	right := end
	pivot := s[(left+right)/2]
	
	for left <=right {

		for s[right] > pivot {
			right--
		}
		for s[left] < pivot {
			left++
		}

		if left <= right {
			s[left], s[right] = s[right], s[left]
			left++
			right--
		}
	}

	if begin < right {
		sort(s, begin, right)
	}
	if left < end {
		sort(s, left, end)
	}
}




