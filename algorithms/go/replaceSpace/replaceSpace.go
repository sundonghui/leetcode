package replacespace

func replaceSpace(s string) string {
	count := 0
	for _, v := range s {
		if v == ' ' {
			count++
		}
	}

	r := make([]byte, len(s)+count*2)
	index := 0
	for _, v := range s {
		if v == ' ' {
			r[index] = '%'
			r[index+1] = '2'
			r[index+2] = '0'
			index += 3
		} else {
			r[index] = byte(v)
			index++
		}
	}
	return string(r)
}
