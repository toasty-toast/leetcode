# 2 - Add Two Numbers

## Problem

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:
```
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
```

## Solution

| Source File | Runtime | Percentile |
| ----------- | ------- | ---------- |
| addTwoNumbers.go | 20 ms | 100% |

The first solution that might come to mind is to extract the values in the linked lists into two integer values, add them, and put the result back into a linked list. This approach has a couple downsides. First, you have to iterate over each linked list once and then iterate over the resultant value to put it back into a linked list. Second (and more importantly), the problem description doesn't give an upper limit to the values, meaning the linked lists might have values too large to fit into a 32-bit or even a 64-bit integer.

The fact that the digits are stored in reverse order gives us the ability to solve this just like we would add two numbers using a pencil and paper. Starting with the lowest digits (the first ones in the list), we add the two digits. If the sum is greater than or equal to 10, carry a one and subtract 10 from the sum. That value is the first digit of the solution. Then we just continue on adding each pair of digits (and the carry!) until we've made it through both linked lists.