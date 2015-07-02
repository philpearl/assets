package lib

var ExtsToMatch = []string{".js", ".css"}

func isInterestingExt(ext string) bool {
	for _, match := range ExtsToMatch {
		if ext == match {
			return true
		}
	}
	return false
}
