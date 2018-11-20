func lengthOfLongestSubstring(s string) int {
    m := make([]int, 256)
    longest := 0
    for start, end := 0, 0; end < len(s); {
        if m[s[end]] == 0 {
            m[s[end]]++
            end++
            longest = max(longest, end - start)
        } else {
            m[s[start]]--
            start++
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