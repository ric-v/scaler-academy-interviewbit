class Solution:
    # @param A : list of integers
    # @return a list of list of integers
    def solve(self, A):
        result = []

        # i is the starting index of the subarray
        for i in range(0, len(A)):

            # j is the ending index of the subarray
            for j in range(i, len(A)):

                intermediate = []
                # k is the index of the element in the subarray A[i..j]
                for k in range(i, j + 1):
                    intermediate.append(A[k])

                # append the subarray to the result
                result.append(intermediate)

        return result


# write test cases here for Solution class
def test_solution():
    s = Solution()
    assert s.solve([1, 2, 3]) == [[1], [1, 2], [1, 2, 3], [2], [2, 3], [3]]
    assert s.solve([1, 2, 3, 4]) == [[1], [1, 2], [1, 2, 3], [
        1, 2, 3, 4], [2], [2, 3], [2, 3, 4], [3], [3, 4], [4]]
    assert s.solve([1, 2, 3, 4, 5]) == [[1], [1, 2], [1, 2, 3], [1, 2, 3, 4], [1, 2, 3, 4, 5], [
        2], [2, 3], [2, 3, 4], [2, 3, 4, 5], [3], [3, 4], [3, 4, 5], [4], [4, 5], [5]]
