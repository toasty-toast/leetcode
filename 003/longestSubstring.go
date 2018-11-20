func lengthOfLongestSubstring(s string) int {
    m := make(map[byte]int)
    longest := 0
    for start, end := 0, 0; end < len(s); {
        if _, contains := m[s[end]]; contains {
            delete (m, s[start])
            start++
        } else {
            m[s[end]] = end
            end++
            longest = max(longest, end - start)
        }
    }
    return longest
}

func max(first int, second int) int {
    if first > second {
        return first
    } else {
        return second
    }
}