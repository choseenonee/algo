package other

func Reverse(arr []int) []int {
	for idx, _ := range arr[:len(arr)/2] {
		arr[len(arr)-1-idx], arr[idx] = arr[idx], arr[len(arr)-1-idx]
	}
	return arr
}
