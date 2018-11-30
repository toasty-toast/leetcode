# 14 - Longest Common Prefix

## Problem

Difficulty: ![#ff9900](https://placehold.it/15/ff9900/000000?text=+) Medium

Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000

For example, two is written as II in Roman numeral, just two one's added together. Twelve is written as, XII, which is simply X + II. The number twenty seven is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

    I can be placed before V (5) and X (10) to make 4 and 9. 
    X can be placed before L (50) and C (100) to make 40 and 90. 
    C can be placed before D (500) and M (1000) to make 400 and 900.

Given an integer, convert it to a roman numeral. Input is guaranteed to be within the range from 1 to 3999.

Example 1:
```
Input: 3
Output: "III"
```
Example 2:
```
Input: 4
Output: "IV"
```
Example 3:
```
Input: 9
Output: "IX"
```
Example 4:
```
Input: 58
Output: "LVIII"
Explanation: L = 50, V = 5, III = 3.
```
Example 5:
```
Input: 1994
Output: "MCMXCIV"
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
```

## Solution

| Source File | Runtime | Percentile |
| ----------- | ------- | ---------- |
| longestCommonPrefix.go | 24 ms | 100% |

To start, we statically define an array of structs that each contain a symbol and its corresponding value. To make less work during the actual processing, we can include combinations of symbols in this array, like including an entry for `400` and `CD`. To convert the number to roman numerals, we iterate over the structs from highest value to lowest and do the following for each.
1. Grab the next struct.
1. Check if the input value is greater than the value in the struct.
1. If it is, subtract the value in the struct from the input value, append the symbol in the struct to the roman numeral string, and go back to step #2.
1. If it is not, move on to the next struct.

Once we exhaust all the structs, the input value should have been reduced to zero, and we should have built up the full roman numeral string for the input value.