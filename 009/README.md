# 9 - Palindrome Number

## Problem

Difficulty: ![#00ff00](https://placehold.it/15/00ff00/000000?text=+) Easy

Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.

Example 1:
```
Input: 121
Output: true
```
Example 2:
```
Input: -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
```
Example 3:
```
Input: 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
```

## Solution

| Source File | Runtime | Percentile |
| ----------- | ------- | ---------- |
| palindromeNumber.go | 56 ms | 100% |

This solution is straightforward. If you reverse the number and the result is the same as the original, you have a palindrome. If the result is different, you do not. A quick check at the top of the function lets us return early if the number is less than zero, since negative numbers are never palindromes.