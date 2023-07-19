# Q3. Main Diagonal Sum
## Problem Description
You are given a N X N integer matrix. You have to find the sum of all the main diagonal elements of A.

Main diagonal of a matrix A is a collection of elements A[i, j] such that i = j.


## Problem Constraints
1 <= N <= 103

-1000 <= A[i][j] <= 1000



## Input Format
There are 1 lines in the input. First 2 integers R, C are the number of rows and columns. Then R * C integers follow corresponding to the rowwise numbers in the 2D array A.



## Output Format
Return an integer denoting the sum of main diagonal elements.



## Example Input
Input 1:

    [1 -2 -3]
    [-4 5 -6]
    [-7 -8 9]

Input 2:

    [3 2]
    [2 3]

Example Output
Output 1:

 15 

Output 2:

 6 


Example Explanation
Explanation 1:

    The size of matrix is 3.
    So, considering the indexing from 0.
    Main diagonal elements will be A[0][0], A[1][1] and A[2][2]
    A[1][1] + A[2][2] + A[3][3] = 1 + 5 + 9 = 15

Explanation 2:

    The size of matrix is 2.
    So, considering the indexing from 0.
    Main diagonal elements will be A[0][0] and A[1][1].
    A[1][1] + A[2][2] = 3 + 3 = 6