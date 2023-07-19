# Q3. Tens Digit Sorting
## Problem Description
Given an array A of N integers. Sort the array in increasing order of the value at the tens place digit of every number.

If a number has no tens digit, we can assume value to be 0.
If 2 numbers have same tens digit, in that case number with max value will come first
Solution should be based on comparator.

## Problem Constraints
1 <= N <= 105

1 <= A[i] <= 109

## Input Format
First argument A is an array of integers.

## Output Format
Return the array after sorting

## Example Input
Input 1:
A = [15, 11, 7, 19]

Input 2:
A = [2, 24, 22, 19]

## Example Output
Output 1:
[7, 19, 15, 11]

Output 2:
[2, 19, 24, 22]

## Example Explanation
For Input 1:
The sorted order is [7, 19, 15, 11]. The tens digit of 7 is 0, 
and that of 19, 15 and 11 is 1.

For Input 2:
The sorted order is [2, 19, 24, 22]. The tens digit of 2 is 0, 
that of 19 is 1 and that of 22 and 24 is 2.
