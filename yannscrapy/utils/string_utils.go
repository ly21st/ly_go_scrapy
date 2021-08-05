package utils

// 去掉重复字符
func UniqueString(str string, a byte) string {
	byteArr := []byte(str)
	newByteArr := make([]byte, len(byteArr))
	var i int
	for index, c := range byteArr {
		if index == 0 {
			newByteArr[i] = c
			i++
			continue
		}

		if c == byteArr[index-1] && c == a {
			continue
		}
		newByteArr[i] = c
		i++
	}
	return string(newByteArr[0:i])
}
