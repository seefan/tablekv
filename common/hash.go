package common

//hash编码
import (
	"strconv"
)

//取hash值
func HashCode(in string) int32 {
	// Initialize output
	var hash int32
	// Empty string has a hashcode of 0
	if len(in) == 0 {
		return hash
	}
	// Convert string into slice of bytes
	b := []byte(in)
	// Build hash
	for i := range b {
		char := b[i]
		hash = ((hash << 5) - hash) + int32(char)
	}
	return hash
}

//取hash值，换成串
func HashString(in string) string {
	code := HashCode(in)
	return strconv.Itoa(int(code))
}

func HashPath(path string) string {
	var ss []byte
	bp := []byte(path)
	for _, s := range bp {
		if s >= 48 && s <= 57 || s >= 65 && s <= 90 || s >= 97 && s <= 122 {
			ss = append(ss, s)
		} else {
			ts := strconv.AppendInt(nil, int64(s), 10)
			ss = append(ss, []byte(ts)...)
		}
	}
	return string(ss)
}
