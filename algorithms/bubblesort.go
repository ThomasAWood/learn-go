package algorithms

func BubbleSort(list []int) []int {
	for range len(list) {
		for i := 0; i < len(list)-1; i++ {
			if list[i] > list[i+1] {
				temp := list[i]
				list[i] = list[i+1]
				list[i+1] = temp
			}
		}
	}
	return list
}
