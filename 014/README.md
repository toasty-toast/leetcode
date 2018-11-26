# 14 - Longest Common Prefix

## Problem

Difficulty: ![#00ff00](https://placehold.it/15/00ff00/000000?text=+) Easy

Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:
```
Input: ["flower","flow","flight"]
Output: "fl"
```
Example 2:
```
Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
```
Note:

All given inputs are in lowercase letters a-z.


## Solution

| Source File | Runtime | Runtime Complexity | Percentile |
| ----------- | ------- | ------------------ | ---------- |
| longestCommonPrefix.go | 0 ms | O(n) | 100% |

The solution is basically two steps, both of which are O(n).
1. Iterate over the list of strings and get the length of the shortest one. We will only have to check that many characters from all the strings since that is the longest the common prefix could possibly be.
1. Iterate over the list of strings, checking one character at a time that the strings are all still the same. As soon as we encounter a character that is different between two or more words, we have reached the end of the common prefix and can return all the characters up to the current one.