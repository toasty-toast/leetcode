# 14 - Longest Common Prefix

## Problem

Difficulty: ![#00ff00](https://placehold.it/15/00ff00/000000?text=+) Easy

Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

    Open brackets must be closed by the same type of brackets.
    Open brackets must be closed in the correct order.

Note that an empty string is also considered valid.

Example 1:
```
Input: "()"
Output: true
```
Example 2:
```
Input: "()[]{}"
Output: true
```
Example 3:
```
Input: "(]"
Output: false
```
Example 4:
```
Input: "([)]"
Output: false
```
Example 5:
```
Input: "{[]}"
Output: true
```

## Solution

| Source File | Runtime | Runtime Complexity | Percentile |
| ----------- | ------- | ------------------ | ---------- |
| validParentheses.go | 0 ms | O(n) | 100% |

To solve this problem we use a `stack`, which I approximated using an array because Go does not have a builtin `stack` type, and walk through the string one character at a time doing the following for each character.
1. Check if the character is an opening character such as `'('`, `'{'`, or `'['`, or a closing character such as `')'`, `'}'`, or `']'`.
	1. If is an opening character, push it onto the top of the stack.
	1. If it is a closing character, check whether it's "pair" is on the top of the stack. For example, if we reach the character `')'`, we expect `'('` to be on the top of the stack.
		1. If the expected character is on top of the stack, pop it off.
		1. If an unexpected character is on top of the stack, return `false`, because the order of the characters in the string is not valid.
1. After processing all the characters in the string, if the stack is empty, the string is valid because all opening characters that we put on the stack had a corresponding closing character that allowed us to pop them off the stack.