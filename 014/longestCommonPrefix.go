import (
    "math"
    "strings"
)

func longestCommonPrefix(strs []string) string {
    length := shortestStringLength(strs)
    var builder strings.Builder
    for i := 0; i < length; i++ {
        for j := 1; j < len(strs); j++ {
            if strs[j][i] != strs[0][i] {
                return builder.String()
            }
        }
        builder.WriteByte(strs[0][i])
    }
    return builder.String()
}

func shortestStringLength(strs []string) int {
    if len(strs) == 0 {
        return 0
    }
    
    min := math.MaxInt32
    for _, s := range strs {
        length := len(s)
        if length < min {
            min = length
        }
    }
    return min
}
