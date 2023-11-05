package deleteslicenode

func Delete(slice []int, i int) []int {
	if len(slice) <= 0 {
		return []int{}
	}
	slice = append(slice[:i], slice[i+1:]...)
	return slice
}
