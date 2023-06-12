# Q4. Maximum Submatrix Sum
## Problem Description
Given a row-wise and column-wise sorted matrix A of size N * M.
Return the maximum non-empty submatrix sum of this matrix.

## Problem Constraints
1 <= N, M <= 1000

-109 <= A[i][j] <= 109

## Input Format
The first argument is a 2D integer array A.

## Output Format
Return a single integer that is the maximum non-empty submatrix sum of this matrix.

## Example Input
    Input 1:-
    A = -5 -4 -3
        -1  2  3
         2  2  4

    Input 2:-
    A = 1 2 3
        4 5 6
        7 8 9


## Example Output
Output 1:-
12

Output 2:-
45

## Example Explanation
Expanation 1:-
    
    The submatrix with max sum is 
    -1 2 3
     2 2 4
     Sum is 12.

Explanation 2:-
   
    The largest submatrix with max sum is 
    1 2 3
    4 5 6
    7 8 9
    The sum is 45.