package helperFunctions

func GetCharacter(lines []string, c rune) []string {
	index := int(c) - 32
	start := (index * 9) + 1
	return lines[start : start+8]
}
