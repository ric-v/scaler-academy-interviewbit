# Q8. Pick from both sides!
## Problem Description
You are given an integer array A of size N.

You have to perform B operations. In one operation, you can remove either the leftmost or the rightmost element of the array A.

Find and return the maximum possible sum of the elements that were removed after B operations.

NOTE: Suppose B = 4, and array A contains 10 elements, then

You can remove the first four elements or can remove the last four elements, or can remove 1 from front and 3 from the back, etc. You need to return the maximum possible sum of elements you can remove.


## Problem Constraints
1 <= N <= 105

1 <= B <= N

-103 <= A[i] <= 103



## Input Format
First argument is an integer array A.

Second argument is an integer B.



## Output Format
Return an integer denoting the maximum possible sum of elements you removed.



## Example Input
Input 1:

 A = [5, -2, 3 , 1, 2]
 B = 3

Input 2:

 A = [ 2, 3, -1, 4, 2, 1 ]
 B = 4


## Example Output
Output 1:

 8

Output 2:

 9


## Example Explanation
Explanation 1:

 Remove element 5 from front and element (1, 2) from back so we get 5 + 1 + 2 = 8

Explanation 2:

 Remove the first element and the last 3 elements. So we get 2 + 4 + 2 + 1 = 9