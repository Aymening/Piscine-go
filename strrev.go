package piscine

func StrRev(s string) string {
	RevStr := ""
	for _, char := range s {
		RevStr = string(char) + RevStr
	}
	return RevStr
}
