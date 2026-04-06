package main

func main() {
	print(ToCapitalCase("Hello! Great to see you! How-are-you-doing-2day?"))
}
func ToCapitalCase(s string) string {
    result := ""
    prev := ' '

    for _, v := range s {
        isAlphaNum := isAlpha(prev) || isDigit(prev)

        switch {
        case isLower(v) && !isAlphaNum:
            result += string(v - 32)
        case isUpper(v) && isAlphaNum:
            result += string(v + 32)
        default:
            result += string(v)
        }
        prev = v
    }
    return result
}

func isLower(r rune) bool { return r >= 'a' && r <= 'z' }
func isUpper(r rune) bool { return r >= 'A' && r <= 'Z' }
func isDigit(r rune) bool { return r >= '0' && r <= '9' }
func isAlpha(r rune) bool { return isLower(r) || isUpper(r) }