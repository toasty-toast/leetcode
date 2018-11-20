# 1 - Two Sum

## Problem

Difficulty: ![#00ff00](https://placehold.it/15/00ff00/000000?text=+) Easy

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:
```
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
```

## Initial Solution

| Source File | Runtime | Percentile |
| ----------- | ------- | ---------- |
| twoSum.go | 8 ms | 52.37% |

My initial solution was to build up a map that mapped a value from the array to its index. If at any point, the difference between the target value and the current value in the array existed in the map, the first index of the solution could be extracted from the map and the second index was just the current loop index.

## Optimized Solution

| Source File | Runtime | Percentile |
| ----------- | ------- | ---------- |
| twoSumOptimized.go | 4 ms | 100% |

I don't take credit for this optimization. I just looked at the top submissions to get an idea of what I missed.

Overall the approach with the map didn't look much different from my initial attempt. I did a little reading and it seems like `len()` has a bit of overhead for non-local slices, so perhaps using `range` instead of `len()` bought me those few extra milliseconds.