package util

//Bytereverse copies the byte values to a newly created array in reverse order
func Bytereverse(s []byte) []byte {
	b := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		b[i] = s[len(s)-1-i]
	}
	return b
}
