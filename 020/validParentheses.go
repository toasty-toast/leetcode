func isValid(s string) bool {
    var stack []rune
    stackSize := 0
    for _, char := range s {
        switch char {
        case '(':
            fallthrough
        case '{':
            fallthrough
        case '[':
            stack = append(stack, char)
            stackSize++
            break
        case ')':
            if stackSize == 0 || stack[stackSize - 1] != '(' {
                return false
            }
            stack = stack[:(stackSize - 1)]
            stackSize--
            break
        case '}':
            if stackSize == 0 || stack[stackSize - 1] != '{' {
                return false
            }
            stack = stack[:(stackSize - 1)]
            stackSize--
            break
        case ']':
            if stackSize == 0 || stack[stackSize - 1] != '[' {
                return false
            }
            stack = stack[:(stackSize - 1)]
            stackSize--
            break
        }
    }
    return stackSize == 0
}
