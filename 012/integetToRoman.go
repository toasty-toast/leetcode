import "strings"

var numerals = []struct {
    value int
    symbol string
}{
    {1000, "M"},
    {900, "CM"},
    {500, "D"},
    {400, "CD"},
    {100, "C"},
    {90, "XC"},
    {50, "L"},
    {40, "XL"},
    {10, "X"},
    {9, "IX"},
    {5, "V"},
    {4, "IV"},
    {1, "I"},
}

func intToRoman(num int) string {
    var builder strings.Builder
    for _, numeral := range numerals {
        for num >= numeral.value  {
            builder.WriteString(numeral.symbol)
            num -= numeral.value
        }
    }
    return builder.String()
}