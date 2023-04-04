# Q4. Matrix Subtraction
## Problem Description

You are given two integer matrices A and B having same size(Both having same number of rows (N) and columns (M). You have to subtract matrix B from A and return the resultant matrix. (i.e. return the matrix A - B).

If A and B are two matrices of the same order (same dimensions). Then A - B is a matrix of the same order as A and B and its elements are obtained by doing an element wise subtraction of A from B.


## Problem Constraints

1 <= N, M <= 103

-109 <= A[i][j], B[i][j] <= 109



## Input Format

The first argument is the 2D integer array A
The second argument is the 2D integer array B


## Output Format

Return a 2D matrix denoting A - B.



## Example Input

Input 1:

    A =  [[1, 2, 3], 
        [4, 5, 6], 
        [7, 8, 9]]

    B =  [[9, 8, 7], 
        [6, 5, 4], 
        [3, 2, 1]]

Input 2:

    A = [[1, 1]]
    
    B = [[2, 3]] 


## Example Output

Output 1:

    [[-8, -6, -4],
    [-2, 0, 2],
    [4, 6, 8]]
    Output 2:

    [[-1, -2]]


## Example Explanation

Explanation 1:

    [[1 2 3]
    [4 5 6]
    [7 8 9]] - [[9 8 7]
                [6 5 4]
                [3 2 1]] = [[-8 -6 -4]
                            [-2 0 2]
                            [4 6 8]]

Explanation 2:

    [[1, 1]] - [[2, 3]] = [[-1, -2]]

