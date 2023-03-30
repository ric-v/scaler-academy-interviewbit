from math import sqrt


class Solution:
    # @param A : long
    # @return an integer
    def solve(self, A):
        if A == 1:
            return 0
        if self.count_prime(A) == 2:
            return 1
        return 0

    def count_prime(self, n):
        count = 0
        for i in range(1, int(sqrt(n))):
            if n % i == 0:
                if n / i == i:
                    count += 1
                else:
                    count += 2
        return count


# main
if __name__ == "__main__":
    n = int(input())
    print(Solution().solve(n))
