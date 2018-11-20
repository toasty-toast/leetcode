# 3 - Reverse Integer

## Problem

Difficulty: ![#00ff00](https://placehold.it/15/00ff00/000000?text=+) Easy

Given a 32-bit signed integer, reverse digits of an integer.

Example 1:
```
Input: 123
Output: 321
```
Example 2:
```
Input: -123
Output: -321
```
Example 3:
```
Input: 120
Output: 21
```
Note:
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−2^31,  2^31 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.


## Solution

| Source File | Runtime | Percentile |
| ----------- | ------- | ---------- |
| reverseInteger.go | 4 ms | 100% |

The method of reversing an integer value is as follows.
1. Grab the last digit of the value using the mod operator. For example, `123 % 10 = 3`.
1. Multiply the working value by 10 and add the digit. Starting with `0`, we would have `0 * 10 + 3 = 3`. For the next digit, `3 * 10 + 2 = 32`.
1. Repeat until you've gotten through all the digits.

What's trick about this one is that you must return `0` if the value would overflow a 32-bit integer. Luckily, Go has a 64-bit integer type, `int64`, so we can just convert the original value to a 64-bit value and do all the work of reversing the integer without worrying about overflow. At the end, we can check if the result is outside the range of a 32-bit integer and return `0` if we need to. Otherwise we can safely convert the value back to an `int` and return it.