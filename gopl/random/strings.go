package random

import "math/rand"

var charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// RandomString generates a random string with a random length.
//
// Parameter:
//
//	maxlength - max length of string
//
// Return type:
//
//	string
func RandomString(maxlength int) string {
	return RandomStringWithLength(rand.Intn(maxlength))
}

// RandomStringWithLength generates a random string with the specified length.
//
// Parameter:
//
//	length - the length of the random string to generate.
//
// Return type:
//
//	string
func RandomStringWithLength(length int) string {
	s := make([]byte, length)
	for i := 0; i < length; i++ {
		s[i] = charset[rand.Intn(length)]
	}
	return string(s)
}
