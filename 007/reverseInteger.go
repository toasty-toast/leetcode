import "math"

func reverse(x int) int {
    value := int64(x)
    var reverse int64 = 0
    for ; value != 0; value /= 10 {
        reverse = (reverse * 10) + (value % 10)
    }
    
    if reverse > math.MaxInt32 || reverse < math.MinInt32 {
        return 0
    }
    
    return int(reverse)
}