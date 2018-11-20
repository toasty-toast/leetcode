# 3 - Longest Substring Without Repeating Characters

## Problem

Difficulty: ![#ff9900](https://placehold.it/15/ff9900/000000?text=+) Medium

Given a string, find the length of the longest substring without repeating characters.

Example 1:
```
Input: "abcabcbb"
Output: 3 
Explanation: The answer is "abc", with the length of 3. 
```
Example 2:
```
Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
```
Example 3:
```
Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3. 
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
```

## Initial Solution

| Source File | Runtime | Percentile |
| ----------- | ------- | ---------- |
| longestSubstring.go | 16 ms | 49.29% |

My solution was to walk through the string with two indexes into the string, `start` and `end`, incrementing `end` until it pointed to a character that was already in the substring. At that point, I moved `start` forward until the substring was once again comprised of unique characters. I used a `map` to track which characters were in the current substring.

## Optimized Solution

| Source File | Runtime | Percentile |
| ----------- | ------- | ---------- |
| longestSubstringOptimized.go | 4 ms | 100% |

I don't take credit for this optimization. I just looked at the top submissions to get an idea of what I missed.

The optimal solution is to do something similar to what I had done, but to just use an `array` instead of a `map`. The array is  indexed by the character in the string, so if the value at a given index is non-zero, we know that character already exists somewhere in the substring. I will note that while the `array` is faster than the `map`, it only works for ASCII strings since the values of non-ASCII characterd would fall outside the bounds of the array.