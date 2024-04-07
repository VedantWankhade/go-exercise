package random

func SliceOfRandomStringsWithLength(length int, size int) []string {
	slice := make([]string, size)
	for i := 0; i < size; i++ {
		slice[i] = RandomStringWithLength(length)
	}
	return slice
}
