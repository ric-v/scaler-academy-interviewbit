from math import sqrt


class Solution:

    def solve(self, n):
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
